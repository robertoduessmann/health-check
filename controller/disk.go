package controller

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/robertoduessmann/health-check/model"
)

type Devices struct {
	BlockDevices []DeviceDescriptor `json:"blockdevices"`
}

type DeviceDescriptor struct {
	Mountpoint string             `json:"mountpoint"`
	Type       string             `json:"type"`
	Name       string             `json:"name"`
	Children   []DeviceDescriptor `json:"children"`
}

func HardDiskCheck() (hardDisks []model.HardDisk) {
	lsblkCommand := ExecCommand("lsblk --json")
	lsblkResult := lsblkCommand.String()

	devices, err := parseLsblk(lsblkResult)
	if err != nil {
		log.Println("Failed to parse lsblk output")
		return
	}

	diskNames := getDiskNames(devices)

	for _, diskName := range diskNames {
		dfCommand := ExecCommand("df -B1 /dev/" + diskName)
		dfResult := dfCommand.String()
		diskInfo, err := parseDf(dfResult)
		if err != nil {
			log.Printf("df parse error for /dev/%s: %q", diskName, err)
			continue
		}
		diskInfo.Name = diskName
		hardDisks = append(hardDisks, diskInfo)
	}

	return
}

func getDiskNames(devices Devices) (diskNames []string) {
	for _, device := range devices.BlockDevices {
		if device.Type == "disk" {
			if device.Mountpoint != "" && device.Mountpoint != "[SWAP]" {
				diskNames = append(diskNames, device.Name)
			}

			for _, part := range device.Children {
				if part.Mountpoint != "" && part.Mountpoint != "[SWAP]" {
					diskNames = append(diskNames, part.Name)
				}
			}
		}
	}

	return
}

func parseLsblk(cmdOutput string) (devices Devices, err error) {
	err = json.Unmarshal([]byte(cmdOutput), &devices)
	return
}

func parseDf(cmdOutput string) (disk model.HardDisk, err error) {
	lines := strings.Split(cmdOutput, "\n")
	if len(lines) != 3 { // 3 is empty newline at the end
		return disk, errors.New("Invalid df output")
	}

	fields := strings.Fields(lines[1])
	disk.Total = fields[1]
	disk.Used = fields[2]
	disk.Free = fields[3]

	return
}
