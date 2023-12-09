package sentinel

import (
	"log"
	"time"
)

type Sentinel interface {
	UpdateTime()
}

type sentinelImpl struct {
	lastUpdateTime time.Time
	checkPeriod    time.Duration
	alertChan      chan<- CheckStatus
	updateChan     <-chan struct{}
	state          sentinelState
}

func NewSentinel(checkPeriodSeconds int, alertChan chan CheckStatus, updateChan chan struct{}) Sentinel {
	s := &sentinelImpl{
		lastUpdateTime: time.Now(),
		checkPeriod:    time.Duration(checkPeriodSeconds) * time.Second,
		alertChan:      alertChan,
		updateChan:     updateChan,
		state:          stateRunning,
	}
	go s.UpdateTime()
	go s.CheckTime()
	return s
}

func (s *sentinelImpl) UpdateTime() {
	for range s.updateChan {
		log.Println("Updating time")
		s.lastUpdateTime = time.Now()
		if s.state == stateWaiting {
			s.state = stateRunning
			s.alertChan <- CheckStatusOk
		}
	}
}

func (s *sentinelImpl) CheckTime() {
	for {
		deadLine := s.lastUpdateTime.Add(s.checkPeriod)
		if s.state == stateRunning && time.Now().After(deadLine) {
			log.Println("Ping expired")
			s.state = stateWaiting
			s.alertChan <- CheckStatusBad
		}
		time.Sleep(1 * time.Second)
	}
}

type sentinelState = int

const (
	stateRunning sentinelState = 0
	stateWaiting sentinelState = 1
)

type CheckStatus int

const (
	CheckStatusOk  CheckStatus = 0
	CheckStatusBad CheckStatus = 1
)
