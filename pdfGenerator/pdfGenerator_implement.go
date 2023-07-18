package pdfGenerator

import (
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)


type wkhtmlpdf struct {

	rootPath string
}

func NewWkToPdf(rootPath string) PDFGeneratorInterface {

	return (&wkhtmlpdf{rootPath: rootPath})

}

func (w *wkhtmlpdf) Create(htmlFile string) (string, error) {

		f, err := os.Open(htmlFile)
		if err != nil {
			return "", nil
		}	
		
		pdfg, err := wkhtmltopdf.NewPDFGenerator()	

		pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

		if err := pdfg.Create(); err != nil {
			return "", err
		}


		fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

		if err := pdfg.WriteFile(fileName); err != nil {
			return "", err
		}


			return fileName, nil

}