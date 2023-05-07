package tool

import "reflect"

func StructConvert(source interface{}, target interface{}) {
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)
	targetType := reflect.TypeOf(target).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < sourceType.NumField(); i++ {
		fieldName := sourceType.Field(i).Name
		fieldValue := sourceValue.FieldByName(fieldName)

		if targetFieldType, ok := targetType.FieldByName(fieldName); ok {
			targetFieldValue := targetValue.FieldByName(fieldName)

			if targetFieldValue.CanSet() && fieldValue.Type() == targetFieldType.Type {
				targetFieldValue.Set(fieldValue)
			}
		}
	}
}
