package main

import (
	"regexp"
	"fmt"
)


// cheng shi jiexiqi   jiexi chulai   hen duo de yonghu   qu qian 10 ge


const context = "880218end        12345yyy   444 555 666tyuio   123456"
const regex = "\\d{6}"

func main() {


	reg := regexp.MustCompile(regex)
	data := reg.Find([]byte(context))

	fmt.Println(string(data))  //880218

	matches := reg.FindAllSubmatch([]byte(context) , -1)
	fmt.Println(len(matches))  //880218


	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~



	//const context2 = `<div class="m-btn purple" data-v-8b1eac0c>workplace:nanjing</div>`
	//const regex2 = `<div class="m-btn purple"*>workplace:(.*?)</div>`

	const context2 = `<div class="m-btn purple" data-v-8b1eac0>workplace:nanjing</div>`
	//const regex2 = `<div class="m-btn purple"*?>workplace:(.*?)</div>`
	const regex2 = `>workplace:(.*?)</div>`

	reg2 := regexp.MustCompile(regex2)
	data2 := reg2.Find([]byte(context2))

	matches2 := reg2.FindAllSubmatch([]byte(context2) , -1)


	fmt.Println(string(matches2[0][1]))//

	fmt.Println(string(data2))
	fmt.Println(len(matches2))



}

//
//func ParseCity(contents []byte) engine.ParseResult {
//
//
//
//
//
//	re := regexp.MustCompile(regex)
//	matches := re.FindAllSubmatch(contents , -1)
//
//	result := engine.ParseResult{}
//	for _, m := range matches {
//		// zheh ge user jiashang weile   qu fen
//		result.Items = append(result.Items,"User "+ string(m[2]))
//		result.Requests = append(result.Requests, engine.Request{
//			Url : string(m[1]),
//			ParserFunc : engine.NilParser,
//		})
//	}
//
//	return result
//
//}
