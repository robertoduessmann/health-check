package controller

import (
	"encoding/json"
	"log"

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
		// Get "df /dev/diskName" output
		log.Println(diskName)
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
