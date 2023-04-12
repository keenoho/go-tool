package tool

func MakePageSize(page int, size int) (offset int, limit int) {
	limit = size
	offset = (page - 1) * size
	return offset, limit
}
