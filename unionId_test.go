package tool_test

import (
	"fmt"
	"testing"

	"github.com/keenoho/go-tool"
)

func TestUnionId16(t *testing.T) {
	idMap := map[int64]int{}
	num := 0
	// 大于每秒1e5时会有重复id
	for num < 1e5 {
		id := tool.UnionId16()
		_, isExist := idMap[id]
		if isExist {
			idMap[id] += 1
			fmt.Println("repeatId 16:", id)
		} else {
			idMap[id] = 1
		}
		num += 1
	}

	t.Log("ok")
}

func TestUnionId32(t *testing.T) {
	idMap := map[float64]int{}
	num := 0
	// 大于每秒1e5时会有重复id
	for num < 1 {
		id := tool.UnionId32()
		_, isExist := idMap[id]
		if isExist {
			idMap[id] += 1
			fmt.Println("repeatId 32:", id)
		} else {
			idMap[id] = 1
		}
		num += 1
		fmt.Println(id)
	}

	t.Log("ok")
}
