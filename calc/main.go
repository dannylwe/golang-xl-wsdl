package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	gen "github.com/danny/wsdl/calc/myservice"
	"github.com/hooklift/gowsdl/soap"
)

type AddResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AddResponse"`

	AddResult int32 `json:"AddResult,omitempty"`
}

func main() {
	client := soap.NewClient("http://www.dneonline.com/calculator.asmx?WSDL")
	service := gen.NewCalculatorSoap(client)
	reply, err := service.Add(&gen.Add{IntA: 2, IntB: 5})

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(reply)
	// &{{http://tempuri.org/ AddResponse} 7}

	responseStruct := AddResponse{}

	out, _ := json.Marshal(reply)
	// fmt.Println(string(out))
	// {"XMLName":{"Space":"http://tempuri.org/","Local":"AddResponse"},"AddResult":7}

	if err := json.Unmarshal(out, &responseStruct); err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", responseStruct.AddResult)
}
