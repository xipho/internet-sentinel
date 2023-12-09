package main

import (
	"internet-pinger/internal/notifier"
	"internet-pinger/internal/pinger"
	"internet-pinger/internal/sentinel"
	"log"
	"os"
	"strconv"
)

func main() {
	token := os.Getenv("PINGER_TOKEN")
	if token == "" {
		log.Fatal("No token provided!")
	}
	chatId, err := strconv.ParseInt(os.Getenv("PINGER_CHAT"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	addr := os.Getenv("PINGER_ADDR")
	if addr == "" {
		addr = "localhost:16385"
	}
	checkPeriodSecondsStr := os.Getenv("PINGER_CHECK_PERIOD")
	var checkPeriodSeconds int
	if checkPeriodSecondsStr == "" {
		checkPeriodSeconds = 70
	} else {
		checkPeriodSeconds, err = strconv.Atoi(checkPeriodSecondsStr)
		if err != nil {
			checkPeriodSeconds = 70
		}
	}
	n := notifier.NewNotifier(token, chatId)

	alertChan := make(chan sentinel.CheckStatus)
	pingChan := make(chan struct{})

	go func() {
		for status := range alertChan {
			var message string
			switch status {
			case sentinel.CheckStatusOk:
				message = "Интернет на Живописной снова появился!"
			case sentinel.CheckStatusBad:
				message = "Интернет на Живописной пропал!"
			}
			err := n.Notify(message)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	sentinel.NewSentinel(checkPeriodSeconds, alertChan, pingChan)
	log.Fatal(pinger.Start(addr, pingChan))
}
