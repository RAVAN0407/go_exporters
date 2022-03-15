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

func random_counter() {

	a := float64(rand.Intn(200))
	fmt.Println(a)
	counter1.Add(a)
	counter2.Add(a)
	time.Sleep(2 * time.Second)
	counter1.Inc()

}

var (
	counter1 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "counter1",
		Help: "random number 1",
	})

	counter2 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "counter2",
		Help: "random number2",
	})
)

func main() {
	random_counter()
	fmt.Println("server started at port 2112")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}
