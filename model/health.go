package model

type Health struct {
	Memory    Memory     `json:"memory"`
	HardDisks []HardDisk `json:"hard_disks"`
	Processor Processor  `json:"processor"`
}
