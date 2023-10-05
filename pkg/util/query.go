package util

import (
	"fmt"

	"github.com/blockloop/scan/v2"
	"github.com/riyan-eng/golang-boilerplate-one/infrastructure"
)

type queryStruct struct{}

func NewQuery() *queryStruct {
	return &queryStruct{}
}

func (q *queryStruct) GetIDByUUID(table, uuid string) (id int) {
	if uuid == "" {
		return
	}
	query := fmt.Sprintf(`
		select t.id from %v t where t.uuid::text = '%v'
	`, table, uuid)
	sqlrows, err := infrastructure.SqlDB.Query(query)
	if err != nil {
		PanicIfNeeded(err)
	}
	err = scan.Row(&id, sqlrows)
	if err != nil {
		PanicIfNeeded(BadRequest{
			Message: fmt.Sprintf("%v tidak ditemukan.", table),
		})
	}
	return
}
