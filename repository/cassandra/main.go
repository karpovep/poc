package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"poc/app"
	"poc/config"
	"poc/model"
	"poc/protos/cloud"
	"poc/protos/nodes"
	"poc/repository/cassandra/queries"
	"poc/repository/impls"
)

type (
	ICassandraRepository interface {
		impls.IRepositoryImpl
		CreateTable(params *queries.CreateTableQueryParams) error
	}

	CassandraRepository struct {
		config  *config.CloudConfig
		cluster *gocql.ClusterConfig
		session *gocql.Session
		queries *queries.Queries
	}
)

const isoTable = "iso"
const activeIsoTable = "active_iso"

func NewCassandraRepository(appContext app.IAppContext) ICassandraRepository {
	cfg := appContext.Get("config").(*config.CloudConfig)
	return &CassandraRepository{
		config:  cfg,
		queries: queries.NewQueries(cfg.Cassandra.TemplatesRoot),
	}
}

func (r *CassandraRepository) CreateTable(params *queries.CreateTableQueryParams) error {
	return r.session.Query(r.queries.CreateTable(params)).Exec()
}

func (r *CassandraRepository) DeleteActiveIso(iso *nodes.ISO) error {
	deleteActiveIsoQuery, err := r.queries.DeleteIfExists(&queries.DeleteQueryParams{
		Table:       activeIsoTable,
		WhereClause: "node_id = ? AND type = ? AND id = ?",
	})
	if err != nil {
		return err
	}
	return r.session.Query(deleteActiveIsoQuery, r.config.NodeId, iso.CloudObj.Entity.TypeUrl, iso.CloudObj.Id).Exec()
}

func (r *CassandraRepository) SaveIso(iso *nodes.ISO) error {
	partitionKey, err := r.extractPartitionKey(iso.CloudObj.Id)
	if err != nil {
		return err
	}
	serializedMetadata, err := proto.Marshal(iso.Metadata)
	if err != nil {
		return err
	}
	insertIsoQuery, err := r.queries.Insert(&queries.InsertQueryParams{
		Table:  isoTable,
		Fields: []string{"partition_key", "type", "id", "cloud_object", "metadata", "is_final"},
	})
	if err != nil {
		return err
	}
	// upsert iso object
	if err = r.session.Query(insertIsoQuery, partitionKey, iso.CloudObj.Entity.TypeUrl, iso.CloudObj.Id, iso.CloudObj.Entity.Value, serializedMetadata, iso.CloudObj.IsFinal).Exec(); err != nil {
		return err
	}

	if iso.CloudObj.IsFinal {
		// if object is final - delete active_iso record
		if err = r.DeleteActiveIso(iso); err != nil {
			return err
		}
	} else {
		// otherwise - upsert active_iso record
		insertActiveIsoQuery, err := r.queries.Insert(&queries.InsertQueryParams{
			Table:  activeIsoTable,
			Fields: []string{"node_id", "type", "id", "cloud_object", "metadata"},
		})
		if err != nil {
			return err
		}
		if err = r.session.Query(insertActiveIsoQuery, r.config.NodeId, iso.CloudObj.Entity.TypeUrl, iso.CloudObj.Id, iso.CloudObj.Entity.Value, serializedMetadata).Exec(); err != nil {
			return err
		}
	}

	return nil
}

func (r *CassandraRepository) FindIsoByTypeAndId(objType string, id string) (*nodes.ISO, error) {
	partitionKey, err := r.extractPartitionKey(id)
	if err != nil {
		return nil, err
	}

	query, err := r.queries.Select(&queries.SelectQueryParams{
		Table:       isoTable,
		Fields:      []string{"cloud_object", "metadata", "is_final"},
		WhereClause: "partition_key = ? AND type = ? AND id = ?",
	})
	if err != nil {
		return nil, err
	}

	var cloudObject []byte
	var metadata []byte
	var isFinal bool
	if err := r.session.Query(query, partitionKey, objType, id).Consistency(gocql.One).Scan(&cloudObject, &metadata, &isFinal); err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("session.Query error")
	}

	var isoMeta nodes.IsoMeta
	if err := proto.Unmarshal(metadata, &isoMeta); err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("Could not unmarshal metadata from db")
	}

	return model.NewIsoFromCloudObjectAndMeta(&cloud.CloudObject{
		Id:      id,
		Entity:  &anypb.Any{TypeUrl: objType, Value: cloudObject},
		IsFinal: isFinal,
	}, &isoMeta), nil
}

func (r *CassandraRepository) extractPartitionKey(id string) (string, error) {
	timeuuid, err := gocql.ParseUUID(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d-%02d", timeuuid.Time().Year(), timeuuid.Time().Month()), nil
}

func (r *CassandraRepository) ListActiveIso(nodeId string, limit int, page []byte) ([]*nodes.ISO, []byte, error) {
	query, err := r.queries.Select(&queries.SelectQueryParams{
		Table:       activeIsoTable,
		Fields:      []string{"type", "id", "metadata", "cloud_object"},
		WhereClause: "node_id = ?",
	})
	if err != nil {
		return nil, nil, err
	}

	itr := r.session.Query(query, nodeId).PageSize(limit).PageState(page).Iter()
	defer itr.Close()

	// Set next page state.
	nextPage := itr.PageState()

	activeIsoList := make([]*nodes.ISO, 0, itr.NumRows())
	scanner := itr.Scanner()
	for scanner.Next() {
		var metadata, cloudObject []byte
		var objType, id string
		if err := scanner.Scan(&objType, &id, &metadata, &cloudObject); err != nil {
			return nil, nil, err
		}

		var isoMeta nodes.IsoMeta
		if err := proto.Unmarshal(metadata, &isoMeta); err != nil {
			return nil, nil, err
		}

		activeIsoList = append(activeIsoList, model.NewIsoFromCloudObjectAndMeta(&cloud.CloudObject{
			Id:     id,
			Entity: &anypb.Any{TypeUrl: objType, Value: cloudObject},
		}, &isoMeta))
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return activeIsoList, nextPage, nil
}

func (r *CassandraRepository) ResetActiveIsoNodeId(iso *nodes.ISO) error {
	batch := r.session.NewBatch(1) // 1 - UnloggedBatch
	serializedMetadata, err := proto.Marshal(iso.Metadata)
	if err != nil {
		return err
	}
	insertActiveIsoQuery, err := r.queries.Insert(&queries.InsertQueryParams{
		Table:  activeIsoTable,
		Fields: []string{"node_id", "type", "id", "cloud_object", "metadata"},
	})
	if err != nil {
		return err
	}
	batch.Query(insertActiveIsoQuery, r.config.NodeId, iso.CloudObj.Entity.TypeUrl, iso.CloudObj.Id, iso.CloudObj.Entity.Value, serializedMetadata)

	deleteActiveIsoQuery, err := r.queries.Delete(&queries.DeleteQueryParams{
		Table:       activeIsoTable,
		WhereClause: "node_id = ? AND type = ? AND id = ?",
	})
	if err != nil {
		return err
	}
	batch.Query(deleteActiveIsoQuery, iso.SenderNodeId, iso.CloudObj.Entity.TypeUrl, iso.CloudObj.Id)

	return r.session.ExecuteBatch(batch)
}

func (r *CassandraRepository) Start() {
	r.cluster = gocql.NewCluster(r.config.Cassandra.Hosts...)
	r.cluster.Keyspace = r.config.Cassandra.Keyspace
	r.cluster.Consistency = gocql.Quorum

	if r.cluster.Keyspace == "" {
		log.Fatal("cassandra keyspace config is missing")
		return
	}

	session, err := r.cluster.CreateSession()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("cannot create connection session to cassandra cluster")
		return
	}
	r.session = session

	// ensure tables are exists
	err = r.CreateTable(&queries.CreateTableQueryParams{
		Keyspace:   r.cluster.Keyspace,
		Table:      isoTable,
		PrimaryKey: "(partition_key), type, id",
		Fields: []struct {
			Name string
			Type string
		}{
			{"partition_key", "VARCHAR"},
			{"type", "VARCHAR"},
			{"id", "TIMEUUID"},
			{"metadata", "BLOB"},
			{"cloud_object", "BLOB"},
			{"is_final", "BOOLEAN"},
		},
	})
	if err != nil {
		log.WithFields(log.Fields{"table": isoTable, "error": err}).Fatal("cannot create table")
	}

	err = r.CreateTable(&queries.CreateTableQueryParams{
		Keyspace:   r.cluster.Keyspace,
		Table:      activeIsoTable,
		PrimaryKey: "(node_id), type, id",
		Fields: []struct {
			Name string
			Type string
		}{
			{"node_id", "VARCHAR"},
			{"type", "VARCHAR"},
			{"id", "TIMEUUID"},
			{"metadata", "BLOB"},
			{"cloud_object", "BLOB"},
		},
	})
	if err != nil {
		log.WithFields(log.Fields{"table": activeIsoTable, "error": err}).Fatal("cannot create table")
	}
}

func (r *CassandraRepository) Stop() {
	r.session.Close()
}
