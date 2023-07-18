package pdfGenerator

type PDFGeneratorInterface interface{
	Create(pdfFile string) (string, error)
}