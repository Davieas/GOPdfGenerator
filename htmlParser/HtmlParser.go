package htmlParser

type htmlParserInterface interface{

	Create(templateName string, data interface{}) (string, error)
}