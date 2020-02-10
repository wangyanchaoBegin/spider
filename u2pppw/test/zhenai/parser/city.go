package parser

import (
	"u2pppw/test/engine"
	"regexp"
)

// cheng shi jiexiqi   jiexi chulai   hen duo de yonghu   qu qian 10 ge

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {


	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents , -1)

	result := engine.ParseResult{}
	for _, m := range matches {

		name := string(m[2])

		// zheh ge user jiashang weile   qu fen
		result.Items = append(result.Items,"User "+ name)
		result.Requests = append(result.Requests, engine.Request{
			Url : string(m[1]),
			//ParserFunc : engine.NilParser,
			// !!! han shu shi  biancheng !!! liantonglingwaiyigelei  !!!
			ParserFunc : func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes ,name)
			} ,



		})
	}

	return result

}
