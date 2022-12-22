package main

import (
	"log"
	"net/http"

	"github.com/signintech/gopdf"
)

func pdf_hello(res http.ResponseWriter, req *http.Request) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "./ttf/Anonymous_Pro.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetTextColor(156, 197, 140)

	pdf.SetXY(10, 10)
	pdf.Cell(nil, "Привет из gopdf")

	pdf.SetLineWidth(2)
	pdf.SetLineType("dashed")
	pdf.Line(10, 30, 585, 30)

	pdf.SetXY(10, 40)
	pdf.Cell(nil, "Hello from gopdf")

	res.Header().Set("Content-type", "application/pdf")
	pdf.Write(res)
}

func main() {

	http.HandleFunc("/", pdf_hello)
	//http.ListenAndServe("localhost:8080", nil) //locally
	http.ListenAndServe(":8080", nil) //SAP BTP CloudFounry

}
