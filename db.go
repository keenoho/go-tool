package tool

import (
	"fmt"
	"reflect"
	"strings"
)

func DbMakePageSize(page int, size int) (offset int, limit int) {
	limit = size
	offset = (page - 1) * size
	return offset, limit
}

func DbStructToSelectSql(stc any) string {
	stcType := reflect.TypeOf(stc)
	if stcType.Kind() != reflect.Struct {
		return "*"
	}
	var sql string
	for i := 0; i < stcType.NumField(); i++ {
		field := stcType.Field(i)
		gormTag, isExist := field.Tag.Lookup("gorm")
		if !isExist {
			continue
		}
		splitArr := strings.Split(gormTag, ";")
		var column string
		var table string
		var originColumn string
		for _, v := range splitArr {
			hasColumn := strings.Contains(v, "column:")
			hasTable := strings.Contains(v, "table:")
			hasOriginColumn := strings.Contains(v, "originColumn:")
			if hasColumn {
				column = strings.ReplaceAll(v, "column:", "")
			}
			if hasTable {
				table = strings.ReplaceAll(v, "table:", "")
			}
			if hasOriginColumn {
				originColumn = strings.ReplaceAll(v, "originColumn:", "")
			}
		}
		if len(column) > 0 && len(table) > 0 && len(originColumn) > 0 {
			if len(sql) < 1 {
				sql += fmt.Sprintf("`%s`.`%s` AS `%s`", table, originColumn, column)
			} else {
				sql += fmt.Sprintf(" ,`%s`.`%s` AS `%s`", table, originColumn, column)
			}
		}
	}
	return sql
}
