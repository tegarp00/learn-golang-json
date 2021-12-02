package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	LogJson("Eko")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Tegar", "pratama", "p00"})
}
