package model

type Health struct {
	Memory    Memory    `json:"memory"`
	HardDisk  HardDisk  `json:"hard_disk"`
	Processor Processor `json:"processor"`
}
