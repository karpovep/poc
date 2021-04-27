package repository

import (
	"bytes"
	"fmt"
	"github.com/gocql/gocql"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"poc/app"
	"poc/model"
	"poc/protos/cloud"
	"poc/protos/nodes"
	"text/template"
)

type (
	ICassandraRepository interface {
		IRepository
		CreateTable(params *CreateTableQueryParams) error
		SaveInternalServerObject(obj *nodes.ISO) error
		FindByTypeAndId(objType string, id string) (*nodes.ISO, error)
	}

	CassandraRepository struct {
		*Repository
		cluster *gocql.ClusterConfig
		session *gocql.Session
	}

	CreateTableQueryParams struct {
		Keyspace   string
		Table      string
		PrimaryKey string
		Fields     []struct {
			Name string
			Type string
		}
	}
)

const createTableQueryTemplate = "" +
	"create table if not exists {{ .Keyspace }}.{{ .Table }} " +
	"(" +
	"{{ range .Fields }}" +
	"{{ .Name }} {{ .Type }}," +
	"{{ end }}" +
	"PRIMARY KEY ({{ .PrimaryKey }})" +
	");"

const isoTable = "internal_server_object"

func NewCassandraRepository(appContext app.IAppContext) ICassandraRepository {
	return &CassandraRepository{
		Repository: NewRepository(appContext),
	}
}

func (r *CassandraRepository) CreateTable(params *CreateTableQueryParams) error {
	tmlpt, err := template.New("createTableQueryTemplate").Parse(createTableQueryTemplate)
	if err != nil {
		return err
	}
	createTableQuery := &bytes.Buffer{}
	err = tmlpt.Execute(createTableQuery, params)
	if err != nil {
		return err
	}
	return r.session.Query(createTableQuery.String()).Exec()
}

func (r *CassandraRepository) SaveInternalServerObject(iso *nodes.ISO) error {
	session, err := r.cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	partitionKey, err := r.extractPartitionKey(iso.CloudObj.Id)
	if err != nil {
		return err
	}
	if err := session.Query(`INSERT INTO internal_server_object (partition_key, type, id, cloud_object) VALUES (?, ?, ?, ?)`,
		partitionKey, iso.CloudObj.Entity.TypeUrl, iso.CloudObj.Id, iso.CloudObj.Entity.Value).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *CassandraRepository) FindByTypeAndId(objType string, id string) (*nodes.ISO, error) {
	partitionKey, err := r.extractPartitionKey(id)
	if err != nil {
		return nil, err
	}

	var cloudObject []byte
	if err := r.session.Query(`SELECT cloud_object FROM internal_server_object WHERE partition_key = ? AND type = ? AND id = ?`,
		partitionKey, objType, id).Consistency(gocql.One).Scan(&cloudObject); err != nil {
		log.Fatal(err)
	}

	return model.NewIsoFromCloudObject(&cloud.CloudObject{
		Id:     id,
		Entity: &anypb.Any{TypeUrl: objType, Value: cloudObject},
	}), nil
}

func (r *CassandraRepository) extractPartitionKey(id string) (string, error) {
	timeuuid, err := gocql.ParseUUID(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d-%02d", timeuuid.Time().Year(), timeuuid.Time().Month()), nil
}

func (r *CassandraRepository) Start() {
	r.cluster = gocql.NewCluster(r.config.Cassandra.Hosts...)
	r.cluster.Keyspace = r.config.Cassandra.Keyspace
	r.cluster.Consistency = gocql.Quorum

	if r.cluster.Keyspace == "" {
		log.Fatalln("cassandra keyspace config is missing")
		return
	}

	session, err := r.cluster.CreateSession()
	if err != nil {
		log.Fatalln("cannot create connection session to cassandra cluster", err)
		return
	}
	r.session = session

	// ensure table is created
	err = r.CreateTable(&CreateTableQueryParams{
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
			{"cloud_object", "BLOB"},
		},
	})
	if err != nil {
		log.Fatalln("cannot create table:", isoTable, err)
	}
	go r.setupIncomingHandler()
	r.EventBus.Subscribe(r.inboundChannelName, r.inboundRepoChan)
}

func (r *CassandraRepository) Stop() {
	r.EventBus.Unsubscribe(r.inboundChannelName, r.inboundRepoChan)
	r.session.Close()
}

func (r *CassandraRepository) setupIncomingHandler() {
	for evnt := range r.inboundRepoChan {
		internalServerObject := evnt.Data.(*nodes.ISO)
		err := r.SaveInternalServerObject(internalServerObject)
		//todo handle errors
		if err != nil {
			log.Fatalln("SaveInternalServerObject error", err)
		}
	}
}
