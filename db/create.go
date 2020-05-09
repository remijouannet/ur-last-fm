package db

import (
	"database/sql"
	"fmt"
)

type CreateTable struct {
	Name        string
	IfNotExists bool
	Columns     []sql.ColumnType
}

func createStmt(table *CreateTable) (stmt string) {

	stmt = fmt.Sprintf("CREATE TABLE %s", table.Name)

	if table.IfNotExists {
		stmt = stmt + " IF NOT EXISTS"
	}

	return
}
