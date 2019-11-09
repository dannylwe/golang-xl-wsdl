package main

import (
	"fmt"

	gen "github.com/danny/wsdl/calc/myservice"
	"github.com/hooklift/gowsdl/soap"
)

func main(){
	client := soap.NewClient("http://www.dneonline.com/calculator.asmx?WSDL")
	service := gen.NewCalculatorSoap(client)
	reply, err := service.Add(&gen.Add{ IntA: 2, IntB:5 })

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
	// &{{http://tempuri.org/ AddResponse} 7}
}