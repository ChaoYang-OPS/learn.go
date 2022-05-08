package apis

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	persionalInformation := PersonalInformation{
		name:   "TF",
		Sex:    "男",
		Tall:   1.68,
		Weight: 62,
		Age:    21,
	}
	fmt.Printf("%+v\n", persionalInformation)
	data, err := json.Marshal(persionalInformation)
	if err != nil {
		return
	}
	fmt.Println("Marshal的结果是原始", data)
	fmt.Println("Marshal的结果是string", string(data))
}