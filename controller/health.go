package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robertoduessmann/health-check/model"
)

func HealtCheck(w http.ResponseWriter, r *http.Request) {

	var health model.Health
	health.Memory = MemoryCheck()
	health.HardDisks = HardDiskCheck()

	fmt.Fprintf(w, string(toJSON(health)))
}

func toJSON(health model.Health) []byte {
	respose, err := json.Marshal(health)
	if err != nil {
		fmt.Println(err)
	}
	return respose
}
