package db

import "fmt"

func toLimitSql(pageNo, pageSize uint) string {
	return fmt.Sprintf(" LIMIT %d, %d", (pageNo-1)*pageSize, pageSize)
}
