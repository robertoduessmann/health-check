package controller

import (
	"log"
	"strconv"

	"github.com/robertoduessmann/health-check/model"
	"github.com/shirou/gopsutil/cpu"
)

func ProcessorCheck() *model.Processor {
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		log.Println("Failed to retrieve CPU usage information")
		return nil
	}
	total := strconv.FormatFloat(cpuUsage[0], 'g', 3, 64)
	proc := model.Processor{Total: total}
	return &proc
}
