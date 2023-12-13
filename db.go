package tool

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func DbMakePageSize(page int, size int) (offset int, limit int) {
	limit = size
	offset = (page - 1) * size
	return offset, limit
}

func DbStructToSqlSelect(stc any) string {
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

func DbInsertEntityValue(source any, field string, value any) error {
	tp := reflect.TypeOf(source)
	if !(tp.Kind() == reflect.Ptr && tp.Elem().Kind() == reflect.Struct) {
		return errors.New("the source must be a gorm entity Ptr Struct")
	}
	if len(field) < 1 {
		return errors.New("the format of field is error")
	}

	elem := tp.Elem()
	elemField, hasElemField := elem.FieldByName(field)
	if !hasElemField {
		return fmt.Errorf("the field %s is not exist in this entity", field)
	}
	val := reflect.ValueOf(source)
	if elemField.Type.Kind() != reflect.ValueOf(value).Kind() {
		return fmt.Errorf("the value kind is not '%v'", elemField.Type.Kind())
	}
	val.Elem().FieldByName(field).Set(reflect.ValueOf(value))

	return nil
}
