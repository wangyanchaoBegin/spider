package engine

import (
	"u2pppw/test/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {

	var requests []Request
	for _,r := range seeds {
		requests = append(requests , r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		//// chuli r       request remove 0 ,1- end
		//
		//log.Printf("Fetching %s",r.Url)
		//body , err := fetcher.Fetch(r.Url)
		//
		//if err != nil {
		//	log.Printf("Fetcher : error " + " fetching url %s : %v",
		//		r.Url,err)
		//	continue
		//}
		//
		//// xin de request fanhuilai ,zai shuzulimian , zhi hou zai jinru zhe ge xun huan !!!
		//parseResult := r.ParserFunc(body)


		// zhege shi sm yufa  ???       parseResult,err := worker(r)  zhi qian haode   xianzai baocuo
		//parseResult,err := e.worker(r)

		parseResult,err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests,parseResult.Requests...)


		for _,item := range parseResult.Items {
			log.Printf("simple.go ----- Run method ---- Got item %v" , item)
		}

	}

}


func worker(r Request) (ParseResult,error ){

	// chuli r       request remove 0 ,1- end

	log.Printf("simple.go Fetching %s",r.Url)
	body , err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("simple.go    Fetcher : error " + " fetching url %s : %v",
			r.Url,err)
		//continue
		return ParseResult{},err
	}

	// xin de request fanhuilai ,zai shuzulimian , zhi hou zai jinru zhe ge xun huan !!!
	//parseResult := r.ParserFunc(body)

	return r.ParserFunc(body),nil
}