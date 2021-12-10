package main

import (
	"encoding/json"
	"fmt"
)

// GSCV ...
type GSCV struct {
	Detail Detail `json:"contactDetail,omitempty"`
}

// Detail ...
type Detail struct {
	PhoneNumber []*Phone `json:"phoneNumber"`
}

// Phone ...
// Phone Struct in Contact for GSCV
type Phone struct {
	Number          string `json:"number,omitempty"`
	PhoneNumberType string `json:"phoneNumberType,omitempty"`
	IsPrimary       bool   `json:"isPrimary,omitempty"`
	CountryCode     string `json:"countryCode,omitempty"`
}

func NewGscv() *GSCV {
	return &GSCV{}
}

func main() {
	fmt.Println("test")

	g := NewGscv()

	//g.Detail.PhoneNumber = []*Phone{}
	g.Detail.PhoneNumber = []*Phone{}

	b, _ := json.Marshal(g)

	fmt.Printf("b: %+v", string(b))
}
