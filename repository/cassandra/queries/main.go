package queries

import (
	"bytes"
	"strings"
	"text/template"
)

type (
	CreateTableQueryParams struct {
		Keyspace   string
		Table      string
		PrimaryKey string
		Fields     []struct {
			Name string
			Type string
		}
	}

	InsertQueryParams struct {
		Table  string
		Fields []string
	}

	SelectQueryParams struct {
		Table       string
		Fields      []string
		WhereClause string
	}

	DeleteQueryParams struct {
		Table       string
		WhereClause string
	}

	Queries struct {
		templates *template.Template
	}
)

const (
	CREATE_TABLE_QUERY_TEMPLATE_NAME     string = "create_table.gotmpl"
	INSERT_QUERY_TEMPLATE_NAME           string = "insert.gotmpl"
	SELECT_QUERY_TEMPLATE_NAME           string = "select.gotmpl"
	DELETE_QUERY_TEMPLATE_NAME           string = "delete.gotmpl"
	DELETE_IF_EXISTS_QUERY_TEMPLATE_NAME string = "delete_if_exists.gotmpl"
)

func NewQueries(templatesRoot string) *Queries {
	return &Queries{
		templates: template.Must(template.New(INSERT_QUERY_TEMPLATE_NAME).Funcs(template.FuncMap{
			"StringsJoin":   strings.Join,
			"StringsRepeat": strings.Repeat,
			"Sum": func(i, j int) int {
				return i + j
			},
		}).ParseFiles(
			templatesRoot+"/"+CREATE_TABLE_QUERY_TEMPLATE_NAME,
			templatesRoot+"/"+INSERT_QUERY_TEMPLATE_NAME,
			templatesRoot+"/"+SELECT_QUERY_TEMPLATE_NAME,
			templatesRoot+"/"+DELETE_QUERY_TEMPLATE_NAME,
			templatesRoot+"/"+DELETE_IF_EXISTS_QUERY_TEMPLATE_NAME,
		)),
	}
}

func (q *Queries) query(templateName string, params interface{}) (string, error) {
	queryBuf := &bytes.Buffer{}
	err := q.templates.ExecuteTemplate(queryBuf, templateName, params)
	if err != nil {
		return "", err
	}
	return queryBuf.String(), nil

}

func (q *Queries) CreateTable(params *CreateTableQueryParams) (string, error) {
	return q.query(CREATE_TABLE_QUERY_TEMPLATE_NAME, params)
}

func (q *Queries) Insert(params *InsertQueryParams) (string, error) {
	return q.query(INSERT_QUERY_TEMPLATE_NAME, params)
}

func (q *Queries) Select(params *SelectQueryParams) (string, error) {
	return q.query(SELECT_QUERY_TEMPLATE_NAME, params)
}

func (q *Queries) Delete(params *DeleteQueryParams) (string, error) {
	return q.query(DELETE_QUERY_TEMPLATE_NAME, params)
}

func (q *Queries) DeleteIfExists(params *DeleteQueryParams) (string, error) {
	return q.query(DELETE_IF_EXISTS_QUERY_TEMPLATE_NAME, params)
}
