package htmlParser

import (
	"os"
	"text/template"
	"github.com/google/uuid"
)

type htmlStruct struct{

	rootPath string
}

func New(rootPath string) htmlParserInterface {

	return &htmlStruct{rootPath: rootPath}
}


func (a *htmlStruct) Create(templateName string, data interface{}) (string, error){
	
	templateGen, err := template.ParseFiles(templateName)
	if err != nil {
		return "", nil
	}

	fileName := a.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	err = templateGen.Execute(fileWriter, data)
	if err != nil {
		return "", err
	}



	return fileName, nil
}