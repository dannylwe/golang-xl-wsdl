# wsdl
http://www.dneonline.com/calculator.asmx?WSDL

#### Get wsdl
curl http://www.dneonline.com/calculator.asmx?WSDL > calculator.wsdl

#### see wsdl
cat calculator.wsdl

#### WSDL2Go code generation as well as its SOAP proxy
go get github.com/hooklift/gowsdl/...

#### generate methods and structures to implement SOAP client and server based on wsdl
gowsdl -i calculator.wsdl Reading File ./calculator.wsdl
`generates in folder "myservice". `
