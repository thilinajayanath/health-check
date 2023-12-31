package handler

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/thilinajayanath/health-check/internal/notify"
)

var isAlerted atomic.Bool
var alertTime time.Time
var lastPingTime atomic.Pointer[time.Time]

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tNow := time.Now()
		lastPingTime.Store(&tNow)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, fmt.Sprintf("Expecting a HTTP POST request. Received a HTTP %s request.", r.Method), http.StatusMethodNotAllowed)
	}
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		isAlerted.Store(false)
		tNow := time.Now()
		lastPingTime.Store(&tNow)
		w.WriteHeader(http.StatusOK)
		log.Println("Resetting the alert")
	} else {
		http.Error(w, fmt.Sprintf("Expecting a HTTP POST request. Received a HTTP %s request.", r.Method), http.StatusMethodNotAllowed)
	}
}

func alert(interval int, realert float64, threshold int, topicArn *string) {
	tk := time.NewTicker(time.Duration(1) * time.Minute)
	defer tk.Stop()

	for range tk.C {
		if time.Since(*lastPingTime.Load()).Minutes() > float64(time.Duration(interval)*time.Duration(threshold)) {
			if !isAlerted.Load() {
				notify.Notify(fmt.Sprintf("Alert! an issue with the health check. Time since the last ping: %v", time.Since(*lastPingTime.Load()).Minutes()), topicArn)

				alertTime = time.Now()
				isAlerted.Store(true)
			} else if isAlerted.Load() && (time.Since(alertTime).Minutes() > realert) {
				notify.Notify(fmt.Sprintf("Alert! an issue with the health check. Time since the last ping: %v", time.Since(*lastPingTime.Load()).Minutes()), topicArn)
				alertTime = time.Now()
			} else {
				log.Printf("Already alerted at %v. Elapsed time %v\n", alertTime, time.Since(alertTime).Minutes())
			}
		}
	}
}

func HandleRequests(interval, realert, threshold int, topicArn *string) {
	pingTime := time.Now()
	lastPingTime.Store(&pingTime)
	isAlerted.Store(false)
	http.HandleFunc("/", pingHandler)
	http.HandleFunc("/reset", resetHandler)

	go alert(interval, float64(realert), threshold, topicArn)

	log.Println("Starting http handler")
	err := http.ListenAndServe(":8080", nil)
	log.Printf("An error occured: %v\n", err)
	log.Println("Shutting down")
}
