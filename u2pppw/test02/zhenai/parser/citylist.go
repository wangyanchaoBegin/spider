package parser

import (
	"regexp"
	"u2pppw/test/engine"
)

const pattern = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//func printCityList(contents []byte){

func ParserCityList(contents []byte) engine.ParseResult {

	//pattern := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

	result := engine.ParseResult{}

	// ! 10 ge !!!
	limit := 10

	re := regexp.MustCompile(pattern)

	matches := re.FindAllSubmatch(contents , -1)

	for _, m := range matches {

        //result.Items = append(result.Items,m[2])
		result.Items = append(result.Items,"City " + string(m[2]))
        // ParserFunc : nil ,    hou mian de , bu jia   error
		result.Requests = append(result.Requests, engine.Request{
			Url : string(m[1]),
			//ParserFunc : nil ,
			//ParserFunc : engine.NilParser,              !!!!! xiamian gaicheng xin jian de cityParser
			// func ParseCity(contents []byte) engine.ParseResult
			// zhendui shang mian  de    Url yong xiamian de  ParserFunc : ParseCity, qu zuo
			ParserFunc : ParseCity,
			})

		//fmt.Printf("city : %s,URL : %s\n",m[2],m[1])
		//fmt.Println()
		limit --
		if limit == 0 {
			break
		}

	}

	//fmt.Println("matches size : %s",len(matches))
	return result
}
