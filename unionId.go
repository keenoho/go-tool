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
var serverIpWithoutPoint string = ""

func unionIdGetYearId(now time.Time) string {
	yearId := fmt.Sprintf("%d", (now.Year()))[2:]
	return yearId
}

func unionIdGetWorkerId(unit int) string {
	strLen := 0
	if unit == 16 {
		strLen = 3
	} else if unit == 32 {
		strLen = 6
	}
	str := serverIpWithoutPoint[len(serverIpWithoutPoint)-strLen:]
	return string(str)
}

func unionIdGetTimeId(now time.Time, unit int) string {
	timeId := fmt.Sprintf("%d", now.UnixMicro())
	strLen := 0
	if unit == 16 {
		strLen = 5
	} else if unit == 32 {
		strLen = 10
	}
	timeId = timeId[len(timeId)-strLen:]
	return timeId
}

func unionIdGetRandId(unit int) string {
	randMaxVal := 0
	randNumFmt := ""
	if unit == 16 {
		randMaxVal = 1e3
		randNumFmt = "%03d"
	} else if unit == 32 {
		randMaxVal = 1e10
		randNumFmt = "%010d"
	}
	randNum := rand.Intn(randMaxVal)
	if randNum >= randMaxVal {
		randNum -= 1
	}
	if randNum < 0 {
		randNum = 0
	}
	randNumStr := fmt.Sprintf(randNumFmt, randNum)
	return randNumStr
}

func unionIdGetNextId(unit int) string {
	nextNumStr := ""
	if unit == 16 {
		nextNum16 += 1
		if nextNum16 > 1e3-1 {
			nextNum16 = 0
		}
		nextNumStr = fmt.Sprintf("%03d", nextNum16)
	}
	if unit == 32 {
		nextNum32 += 1
		if nextNum32 > 1e4-1 {
			nextNum32 = 0
		}
		nextNumStr = fmt.Sprintf("%04d", nextNum32)
	}
	return nextNumStr
}

/**
 * 唯一id 16位 = 2 + 3 + 5 + 3 + 3
 * year(2)workerId(3)time(5)rand(3)nextId(3)
 * */
func UnionId16() int64 {
	now := time.Now()
	yearId := unionIdGetYearId(now)
	workerId := unionIdGetWorkerId(16)
	timeId := unionIdGetTimeId(now, 16)
	randId := unionIdGetRandId(16)
	nextId := unionIdGetNextId(16)
	finalId, _ := strconv.ParseInt(yearId+workerId+timeId+randId+nextId, 10, 64)
	return finalId
}

/**
 * 唯一id 16位 = 2 + 3 + 5 + 3 + 3
 * year(2)workerId(3)time(5)rand(3)nextId(3)
 * */
func UnionId16String() string {
	now := time.Now()
	yearId := unionIdGetYearId(now)
	workerId := unionIdGetWorkerId(16)
	timeId := unionIdGetTimeId(now, 16)
	randId := unionIdGetRandId(16)
	nextId := unionIdGetNextId(16)
	return yearId + workerId + timeId + randId + nextId
}

/**
 * 唯一id 32位 = 2 + 6 + 10 + 10 + 4
 * year(2)workerId(6)time(10)rand(10)nextId(4)
 * */
func UnionId32() float64 {
	now := time.Now()
	yearId := unionIdGetYearId(now)
	workerId := unionIdGetWorkerId(32)
	timeId := unionIdGetTimeId(now, 32)
	randId := unionIdGetRandId(32)
	nextId := unionIdGetNextId(32)
	finalId, _ := strconv.ParseFloat(yearId+workerId+timeId+randId+nextId, 64)
	return finalId
}

/**
 * 唯一id 32位 = 2 + 6 + 10 + 10 + 4
 * year(2)workerId(6)time(10)rand(10)nextId(4)
 * */
func UnionId32String() string {
	now := time.Now()
	yearId := unionIdGetYearId(now)
	workerId := unionIdGetWorkerId(32)
	timeId := unionIdGetTimeId(now, 32)
	randId := unionIdGetRandId(32)
	nextId := unionIdGetNextId(32)
	return yearId + workerId + timeId + randId + nextId
}

func init() {
	seg := strings.Split(ServerInternalIp(), ".")
	for i := 0; i < len(seg); i++ {
		val := seg[i]
		str := fmt.Sprintf("%03s", val)
		seg[i] = str
	}
	serverIpWithoutPoint = strings.Join(seg, "")
}
