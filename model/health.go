package model

type Health struct {
	Memory    Memory    `json:"memory"`
	Disks     []Disk    `json:"disks"`
	Processor Processor `json:"processor"`
}
