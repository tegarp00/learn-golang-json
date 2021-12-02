package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Tegar",
		MiddleName: "pratama",
		LastName:   "p00",
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Tegar","MiddleName":"pratama","LastName":"p00","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Tegar",
		Addresses: []Address{
			{
				Street:     "Jalan Buntu",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan Buntu Terus",
				Country:    "Indonesia",
				PostalCode: "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Tegar","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Buntu","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Buntu Terus","Country":"Indonesia","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Buntu","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Buntu Terus","Country":"Indonesia","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	addresses := []Address{
		{
			Street:     "Jalan Buntu",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Buntu Terus",
			Country:    "Indonesia",
			PostalCode: "8888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
