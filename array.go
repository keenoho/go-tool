package tool

type ArrayInterface interface{}

type Array struct {
	Value []any
}

func (arr *Array) Append(value ...any) {
	arr.Value = append(arr.Value, value...)
}

func (arr *Array) Prepend(value ...any) {
	arr.Value = append(value, arr.Value...)
}

func (arr *Array) Shift() any {
	if len(arr.Value) < 1 {
		return nil
	}
	firstOne := arr.Value[0]
	arr.Value = arr.Value[1:]
	return firstOne
}

func (arr *Array) Pop() any {
	if len(arr.Value) < 1 {
		return nil
	}
	lastOne := arr.Value[len(arr.Value)-1]
	arr.Value = arr.Value[:len(arr.Value)-1]
	return lastOne
}

func (arr *Array) Slice(start int, end int) Array {
	arr.Value = arr.Value[start:end]
	return *arr
}

func (arr *Array) Insert(index int, value any) {
	frtPart := arr.Value[0:index]
	aftPart := arr.Value[index:]
	newValue := make([]any, 0)
	newValue = append(newValue, frtPart...)
	newValue = append(newValue, value)
	newValue = append(newValue, aftPart...)
	arr.Value = newValue
}

func (arr *Array) Concat(targetArr Array) Array {
	arr.Value = append(arr.Value, targetArr.Value...)
	return *arr
}
