package scheduler

import "u2pppw/test/engine"

type QueueScheduler struct {

	requestChan chan engine.Request

	// ...???   chan engine.Request guan zai qian mian de chan li mian
	workerChain chan chan engine.Request

}

// ??? fang li mian...
func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChain <- w
}


func (s *QueueScheduler) Run() {

	// zhe liang ge de shi li hua,zai run zhi qian.
	s.requestChan = make(chan engine.Request)
	s.workerChain = make(chan chan engine.Request)


	go func(){

		var requestQ [] engine.Request
		var workerQ [] chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ , r)
			case w := <-s.workerChain:
				workerQ = append(workerQ , w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}


		}


	}()
}


// shou xian rang ta  jicheng  interface
// zhi zhen jie shou zhe   cai neng gai bian li mian  de nei rong

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}


func (s *QueueScheduler)  ConfigureMasterWorkerChan(chan engine.Request) {

}






















