package parser

type Request struct {
	Url string
	Parser func(string)[]Request
}
