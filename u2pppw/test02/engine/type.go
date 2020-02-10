package engine


type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult

}

type ParseResult struct {

	Requests []Request
	// what ?            all kind can
	Items []interface{}

}

func NilParser([] byte) ParseResult {
	return ParseResult{}
}