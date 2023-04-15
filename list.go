package tool

type ListInterface interface{}

type List struct {
	Value []any
}

func (li *List) Append(value ...any) {
	li.Value = append(li.Value, value...)
}

func (li *List) Prepend(value ...any) {
	li.Value = append(value, li.Value...)
}

func (li *List) Shift() any {
	if len(li.Value) < 1 {
		return nil
	}
	firstOne := li.Value[0]
	li.Value = li.Value[1:]
	return firstOne
}

func (li *List) Pop() any {
	if len(li.Value) < 1 {
		return nil
	}
	lastOne := li.Value[len(li.Value)-1]
	li.Value = li.Value[:len(li.Value)-1]
	return lastOne
}

func (li *List) Slice(start int, end int) List {
	li.Value = li.Value[start:end]
	return *li
}

func (li *List) Insert(index int, value any) {
	frtPart := li.Value[0:index]
	aftPart := li.Value[index:]
	newValue := make([]any, 0)
	newValue = append(newValue, frtPart...)
	newValue = append(newValue, value)
	newValue = append(newValue, aftPart...)
	li.Value = newValue
}

func (li *List) Concat(targetli List) List {
	li.Value = append(li.Value, targetli.Value...)
	return *li
}

func (li *List) IndexOf(index int) any {
	if len(li.Value) < 1 {
		return nil
	}
	return li.Value[index]
}
