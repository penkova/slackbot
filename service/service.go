package service

import (
	"bytes"
	"fmt"
	handl "github.com/apenkova/slackbot/service/handlers"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

// ListenAndServeHTTP wraps http.ListenAndServe
func ListenAndServeHTTP() error {
	c := ":8001"
	log.Printf("start ListenAndServeHTTP on %s", c)
	return http.ListenAndServe(c, GetHTTPHandler())
}

// GetHTTPHandler returns configured HTTP Handler
func GetHTTPHandler() http.Handler {
	defer func() {
		if r := recover(); r != nil {
			log.Println("GetHTTPHandler PANIC Recovered:", r)
		}
	}()
	r := chi.NewRouter()
	r.Get("/health", handl.Health)
	r.Get("/version", handl.Version)
	r.Post("/event", handl.PostEvent)

	return r
}

// Task describes task function
type Task func() error

// RunTasks starts tasks
func RunTasks() {
	//go tasks.Run(tasks.ESS3Backup, 1*time.Hour)
	go RunTask(Events, 6*time.Hour)
}
func Events() (err error) {
	url := "http://localhost:8001/event"
	str := fmt.Sprintf(`{"text":"Notification - %+v"}`, time.Now().Format(time.UnixDate))
	var jsonStr = []byte(str)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Send request", err)
	}
	defer resp.Body.Close()
	return err
}

// Run runs task in infinite loop with pause ignores task error
func RunTask(t Task, sleep time.Duration) {
	done := make(chan bool)
	runTask(t, done)
	for {
		<-done
		time.Sleep(sleep)
		runTask(t, done)
	}
}

func runTask(t Task, done chan bool) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic recovered: ", err)
				done <- false
			}
		}()
		err := t()
		if err != nil {
			log.Println("runTask: ", err)
			done <- false
		}
		done <- true
	}()
}
