package queries

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldInitTemplates(t *testing.T) {
	// When
	queries := NewQueries("templates/")

	// Then
	assert.Equal(t, 5, len(queries.templates.Templates()), "unexpected number of templates were initialised")
}

func Test_ShouldExecuteCreateTableTemplate(t *testing.T) {
	// Given
	params := &CreateTableQueryParams{
		Keyspace:   "test_keyspace",
		Table:      "test_table",
		PrimaryKey: "p_key",
		Fields: []struct {
			Name string
			Type string
		}{
			{"p_key", "VARCHAR"},
		},
	}
	expectedQuery := "" +
		"CREATE TABLE IF NOT EXISTS test_keyspace.test_table\n" +
		"(\n" +
		"    p_key VARCHAR,\n" +
		"    PRIMARY KEY (p_key)\n" +
		");\n"
	queries := NewQueries("templates/")

	// When
	actualQuery, err := queries.CreateTable(params)
	if err != nil {
		t.Fatal("CreateTable error", err)
	}

	// Then
	assert.Equal(t, expectedQuery, actualQuery, "unexpected query received")
}

func Test_ShouldExecuteInsertTemplate(t *testing.T) {
	// Given
	params := &InsertQueryParams{
		Table: "test_table",
		Fields: []string{
			"p_key",
			"test_field",
		},
	}
	expectedQuery := "INSERT INTO test_table (p_key, test_field) VALUES (?, ?);"
	queries := NewQueries("templates/")

	// When
	actualQuery, err := queries.Insert(params)
	if err != nil {
		t.Fatal("Insert error", err)
	}

	// Then
	assert.Equal(t, expectedQuery, actualQuery, "unexpected query received")
}

func Test_ShouldExecuteSelectTemplate(t *testing.T) {
	// Given
	params := &SelectQueryParams{
		Table: "test_table",
		Fields: []string{
			"p_key",
		},
		WhereClause: "p_key > 0",
	}
	expectedQuery := "SELECT p_key FROM test_table WHERE p_key > 0;"
	queries := NewQueries("templates/")

	// When
	actualQuery, err := queries.Select(params)
	if err != nil {
		t.Fatal("Select error", err)
	}

	// Then
	assert.Equal(t, expectedQuery, actualQuery, "unexpected query received")
}

func Test_ShouldExecuteDeleteTemplate(t *testing.T) {
	// Given
	params := &DeleteQueryParams{
		Table:       "test_table",
		WhereClause: "p_key > 0",
	}
	expectedQuery := "DELETE FROM test_table WHERE p_key > 0;"
	queries := NewQueries("templates/")

	// When
	actualQuery, err := queries.Delete(params)
	if err != nil {
		t.Fatal("Delete error", err)
	}

	// Then
	assert.Equal(t, expectedQuery, actualQuery, "unexpected query received")
}

func Test_ShouldExecuteDeleteIfExistsTemplate(t *testing.T) {
	// Given
	params := &DeleteQueryParams{
		Table:       "test_table",
		WhereClause: "p_key > 0",
	}
	expectedQuery := "DELETE FROM test_table WHERE p_key > 0 IF EXISTS;"
	queries := NewQueries("templates/")

	// When
	actualQuery, err := queries.DeleteIfExists(params)
	if err != nil {
		t.Fatal("Delete error", err)
	}

	// Then
	assert.Equal(t, expectedQuery, actualQuery, "unexpected query received")
}
