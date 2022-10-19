package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type Data struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

func randFloat(min, max float64) float64 {
	return rand.Float64() * (max - min)
}

func flood(c chan string) {
	for {
		randValue := randFloat(0, 99)
		data := Data{time.Now().UTC(), randValue}

		d, _ := json.Marshal(data)
		c <- string(d)

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	logInfo := log.New(os.Stdout, "INFO ", log.Ldate)
	// logErr := log.New(os.Stderr, "ERROR ", log.Ldate)

	c := make(chan string)

	go flood(c)

	for {
		v := <-c
		logInfo.Println(v)
	}
}
