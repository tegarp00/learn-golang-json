package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Tegar",
		MiddleName: "pratama",
		LastName:   "p00",
	}

	encoder.Encode(customer)
}
