package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)

	ConfigureMasterWorkerChan(chan Request)

}



//func (e ConcurrentEngine) Run(seeds ...Request){
func (e *ConcurrentEngine) Run(seeds ...Request){

	in := make(chan Request)
	out := make(chan ParseResult)

	//
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0 ; i< e.WorkCount ; i++ {
		createWork(in,out)
	}

	// zhu yi shun xv    zuofan zai submit .
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <- out
		//for _,item  := range result.Items {
		//	log.Printf("Got item : %v",item)
		//}

		for _,item  := range result.Items {

			//log.Printf("change change ---- Got item #%d: %v", itemCount ,item)
			log.Printf("#%d ", itemCount)
			log.Printf("Got item : %v" ,item)
			itemCount++
		}

		for _,request  := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

// zhe zhong san ge can shu ??? what ???
func createWork(in chan Request , out chan ParseResult) {


	// !!!
	go func() {
		for {

			request := <- in

			result,err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()


}