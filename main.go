package main

import (
	"fmt"
	"geradorPDF/htmlParser"
	"geradorPDF/pdfGenerator"
)

type Data struct {

	Name string

}



func main(){


	dataHtml := Data{

		Name: "Davi",
	}


h := htmlParser.New("tmp")
wk := pdfGenerator.NewWkToPdf("tmp")

htmlGenerated, err := h.Create("templates/example.html", dataHtml)

if err != nil {
	
	fmt.Println(err)
	return 
}

fmt.Println(htmlGenerated)


filePDFName, err := wk.Create(htmlGenerated)
if err != nil {
	
	fmt.Println(err)
	return 
}
fmt.Println("Pdf Gerado!", filePDFName)


}