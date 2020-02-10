package parser



import (
	"u2pppw/test/engine"
	"regexp"
	"u2pppw/test/model"
)

// cheng shi jiexiqi   jiexi chulai   hen duo de yonghu   qu qian 10 ge

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//const ageRe = `<td><span class="label">`


const workplaceRex = `>工作地:(.*?)</div>`

const incomeRex = `>月收入:(.*?)</div>`


//func ParseProfile(contents []byte) engine.ParseResult {

func ParseProfile(contents []byte , name string) engine.ParseResult {

	profile := model.Profile{}


	// !!!
	profile.Name = name

	//re := regexp.MustCompile(workplace)
	//match := re.FindSubmatch(contents)
	//
	//if match != nil {
	//	profile.Workplace = string(match[1])
	//}


	//re := regexp.MustCompile(workplace)
	//matcheWork := re.FindSubmatch(contents)
	//
	//if matcheWork != nil {
	//	profile.Workplace = string(matcheWork[1])
	//}

	reWork := regexp.MustCompile(workplaceRex)
	profile.Workplace = extractString(contents , reWork)

	reIncome := regexp.MustCompile(incomeRex)
	profile.Income = extractString(contents , reIncome)

	//result := engine.ParseResult{}
	//
	//
	//
	//
	//for _, m := range matches {
	//	// zheh ge user jiashang weile   qu fen
	//	result.Items = append(result.Items,"User "+ string(m[2]))
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url : string(m[1]),
	//		ParserFunc : engine.NilParser,
	//	})
	//}
	//
	//return result

	result := engine.ParseResult{
		Items : [] interface{}{profile},
	}

	return result
}



func extractString(contents []byte , re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >=2 {
		return string(match[1])
	} else {
		return ""
	}
}
