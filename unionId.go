package tool

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var nextNum16 int = 0
var nextNum32 int = 0
var serverIp string = ""

/**
 * 唯一id 16位 = 2 + 2 + 3 + 7 + 2
 * year(2)randId(2)serviceId(3)time(7)nextId(2)
 * */
func UnionId16() int64 {
	serviceId := serverIp[len(serverIp)-3:]
	nextNum16 += 1
	if nextNum16 >= 100 {
		nextNum16 = 0
	}
	now := time.Now()
	yearId := fmt.Sprintf("%d", (now.Year()))[2:]
	timeId := fmt.Sprintf("%d", now.UnixMicro()/1e3)
	timeId = timeId[len(timeId)-7:]
	randId := ""
	randNum := rand.Intn(99)
	if randNum < 10 {
		randId = fmt.Sprintf("0%d", randNum)
	} else {
		randId = fmt.Sprintf("%d", randNum)
	}
	nextId := ""
	if nextNum16 < 10 {
		nextId = fmt.Sprintf("0%d", nextNum16)
	} else {
		nextId = fmt.Sprintf("%d", nextNum16)
	}
	finalId, _ := strconv.ParseInt(yearId+randId+serviceId+timeId+nextId, 10, 64)
	return finalId
}

/**
 * 唯一id 32位 = 2 + 2 + 8 + 6 + 10 + 4
 * year(2)month(2)randId(8)serviceId(6)time(10)nextId(4)
 * */
func UnionId32() int64 {
	serviceId := serverIp[len(serverIp)-6:]
	nextNum32 += 1
	if nextNum32 >= 1000 {
		nextNum32 = 0
	}
	now := time.Now()
	yearId := fmt.Sprintf("%d", (now.Year()))[2:]
	monthId := fmt.Sprintf("%d", (now.Month()))[2:]
	timeId := fmt.Sprintf("%d", now.UnixMicro())
	timeId = timeId[len(timeId)-7:]
	randId := ""
	randNum := rand.Intn(1e9 - 1)
	if randNum < 10 {
		randId = fmt.Sprintf("0%d", randNum)
	} else {
		randId = fmt.Sprintf("%d", randNum)
	}
	nextId := ""
	if nextNum32 < 10 {
		nextId = fmt.Sprintf("0%d", nextNum32)
	} else {
		nextId = fmt.Sprintf("%d", nextNum32)
	}
	finalId, _ := strconv.ParseInt(yearId+monthId+randId+serviceId+timeId+nextId, 10, 64)
	return finalId
}

func init() {
	ipStr := strings.Split(ServerInternalIp(), ".")
	if len(ipStr) < 9 {
		ipStr = append(make([]string, 9-len(ipStr)), ipStr...)
	}
	serverIp = strings.Join(ipStr, "")

}
