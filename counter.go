package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func counter01() {

	a := float64(rand.Intn(200))
	random_Counter1_float64.Add(a)
	for {
		time.Sleep(2 * time.Second)
		random_Counter1_float64.Inc()
	}

}

func counter02() {
	b := float64(rand.Intn(200))
	random_Counter2_float64.Add(b)
	time.Sleep(2 * time.Second)
	random_Counter2_float64.Inc()
}

func gauge01() {
	c := float64(rand.Intn(200))
	random_gauge_float64.Set(c)
	time.Sleep(2 * time.Second)
	random_gauge_float64.Inc()
	random_gauge_float64.Dec()

}

var (
	random_Counter1_float64 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "random_Counter1_float64",
		Help: "random number 1",
	})
	random_Counter2_float64 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "random_Counter2_float64",
		Help: "random number 2",
	})
	random_gauge_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "random_gauge_float64",
		Help:      "random gauge ",
	})
)

func init() {
	prometheus.MustRegister(random_gauge_float64)
}

func main() {

	go counter01()
	go counter02()
	go gauge01()
	fmt.Println("server started at port 9000")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}

}

