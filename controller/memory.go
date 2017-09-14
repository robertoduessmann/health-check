package controller

import (
	"strconv"

	"github.com/robertoduessmann/health-check/model"
)

func MemoryCheck() (memory model.Memory) {
	memoryCommand := ExecCommand("free")
	memoryResult := memoryCommand.String()

	total, _ := strconv.ParseInt(memoryResult[92:99], 0, 64)
	used, _ := strconv.ParseInt(memoryResult[104:111], 0, 64)
	free := total - used

	memory.Total = strconv.FormatInt(total, 10) + " Kb"
	memory.Used = strconv.FormatInt(used, 10) + " Kb"
	memory.Free = strconv.FormatInt(free, 10) + " Kb"

	return
}
