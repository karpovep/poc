package repository

import (
	"bytes"
	"github.com/gocql/gocql"
	"poc/app"
	"poc/bus"
	"poc/config"
	"text/template"
)

type (
	ICassandraRepository interface {
		IRepository
		CreateKeyspace(params *CreateKeyspaceQueryParams) error
		CreateTable(params *CreateTableQueryParams) error
	}

	CassandraRepository struct {
		*Repository
		cluster *gocql.ClusterConfig
	}

	CreateKeyspaceQueryParams struct {
		Keyspace          string
		ReplicationClass  string
		ReplicationFactor int
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

const createKeyspaceQueryTemplate = "" +
	"create keyspace if not exists {{ .Keyspace }} " +
	"with replication = { 'class' : '{{ .ReplicationClass }}', 'replication_factor' : {{ .ReplicationFactor }} };"

const createTableQueryTemplate = "" +
	"create table if not exists {{ .Keyspace }}.{{ .Table }} " +
	"(" +
	"{{ range .Fields }}" +
	"{{ .Name }} {{ .Type }}," +
	"{{ end }}" +
	"PRIMARY KEY ({{ .PrimaryKey }})" +
	");"

func NewCassandraRepository(appContext app.IAppContext) ICassandraRepository {
	eventBus := appContext.Get("eventBus").(bus.IEventBus)
	cfg := appContext.Get("config").(*config.CloudConfig)
	return &CassandraRepository{
		Repository: &Repository{
			EventBus: eventBus,
			config:   cfg,
		},
	}
}

func (r *CassandraRepository) CreateKeyspace(params *CreateKeyspaceQueryParams) error {
	tmlpt, err := template.New("createKeyspaceQueryTemplate").Parse(createKeyspaceQueryTemplate)
	if err != nil {
		return err
	}
	createKeyspaceQuery := &bytes.Buffer{}
	err = tmlpt.Execute(createKeyspaceQuery, params)
	if err != nil {
		return err
	}
	return r.executeQuery(createKeyspaceQuery.String())
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
	return r.executeQuery(createTableQuery.String())
}

func (r *CassandraRepository) executeQuery(query string) error {
	session, err := r.cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()
	return session.Query(query).Exec()
}

func (r *CassandraRepository) Start() {
	r.cluster = gocql.NewCluster(r.config.Cassandra.Hosts...)
	r.cluster.Keyspace = r.config.Cassandra.Keyspace
	r.cluster.Consistency = gocql.Quorum
}

func (r *CassandraRepository) Stop() {
	panic("implement me")
}
