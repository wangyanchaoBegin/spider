package scheduler

import "u2pppw/test/engine"

type SimpleSheduler struct {
	workerChan chan engine.Request
}

//func (s SimpleSheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
//	//panic("implement me")
//	s.workerChan = c
//}
//
//
//func (s SimpleSheduler) Submit(r engine.Request) {
//
//	// sm yi si le ?
//	s.workerChan <- r
//
//}


func (s *SimpleSheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	//panic("implement me")
	s.workerChan = c
}


func (s *SimpleSheduler) Submit(r engine.Request) {

	// sm yi si le ?
	//s.workerChan <- r   !!! zheyang jiu haole  ??? duo xian cheng wen ti
	// yin wei xia mian de shi   li ke fan hui

	go func (){
		s.workerChan <- r
	}()
}