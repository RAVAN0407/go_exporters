package main

import (
	"fmt"
	"main/parser" 
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

<<<<<<< HEAD
var (
	memtotal_gauge01_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "memtotal_gauge01_float64",
		Help:      "random gauge ",
	})
	memfree_gauge02_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "memfree_gauge02_float64",
		Help:      "random gauge ",
	})
	memavailable_gauge03_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "memavailable_gauge03_float64",
		Help:      "random gauge ",
	})
	buffers_gauge04_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "buffers_gauge04_float64",
		Help:      "random gauge ",
	})

	cached_gauge05_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "cached_gauge05_float64",
		Help:      "random gauge ",
	})
	wall_clock_since_boot_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "wall_clock_since_boot_float64",
		Help:      "random gauge ",
	})
	combined_idle_cputime_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "combined_idle_cputime_float64",
		Help:      "random gauge ",
	})
	process_gauge01_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "process_gauge01_float64",
		Help:      "random gauge ",
	})
	process_gauge02_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "process_gauge02_float64",
		Help:      "random gauge ",
	})
	process_gauge03_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "process_gauge03_float64",
		Help:      "random gauge ",
	})
	process_gauge04_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "process_gauge04_float64",
		Help:      "random gauge ",
	})
	process_gauge05_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "process_gauge05_float64",
		Help:      "random gauge ",
	})
	process_gauge06_float64 = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "our_company",
		Subsystem: "blob_storage",
		Name:      "process_gauge06_float64",
		Help:      "random gauge ",
	})
)

func meminfo(){
	go func() {
			for {
					mi := parser.Get_mem_info()
					memfree_gauge02_float64.Set(mi["MemFree:"])
					memavailable_gauge03_float64.Set(mi["MemAvailable:"])
					memtotal_gauge01_float64.Set(mi["MemTotal:"])
					buffers_gauge04_float64.Set(mi["Buffers:"])
					cached_gauge05_float64.Set(mi["Cached:"])
					time.Sleep(2 * time.Second)

			}
	}()


}

func uptime(){
	go func() {
		for {
			ut :=parser.Get_uptime()
			wall_clock_since_boot_float64.Set(ut["walk_clock"])
			combined_idle_cputime_float64.Set(ut["combined_idletime"])
			time.Sleep(2 * time.Second)
		}
=======
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
>>>>>>> func3

	}()
}

<<<<<<< HEAD
func loadavg(){
	go func() {
		for {
				la :=parser.Get_loadavg()
				process_gauge01_float64.Set(la["process1"])
				process_gauge02_float64.Set(la["process2"])
				process_gauge03_float64.Set(la["process3"])
				process_gauge04_float64.Set(la["process4"])
				process_gauge05_float64.Set(la["process5"])
				process_gauge06_float64.Set(la["process6"])
				time.Sleep(2 * time.Second)
		}
	}()
}

func init() { 
	prometheus.MustRegister(memavailable_gauge03_float64)
	prometheus.MustRegister(memfree_gauge02_float64)
	prometheus.MustRegister(memtotal_gauge01_float64)
	prometheus.MustRegister(cached_gauge05_float64)
	prometheus.MustRegister(buffers_gauge04_float64)
	prometheus.MustRegister(wall_clock_since_boot_float64)
	prometheus.MustRegister(combined_idle_cputime_float64)
	prometheus.MustRegister(process_gauge01_float64)
	prometheus.MustRegister(process_gauge02_float64)
	prometheus.MustRegister(process_gauge03_float64)
	prometheus.MustRegister(process_gauge04_float64)
	prometheus.MustRegister(process_gauge05_float64)
	prometheus.MustRegister(process_gauge06_float64)
}
=======
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
>>>>>>> func3

func init() {
	prometheus.MustRegister(random_gauge_float64)
}

func main() {
<<<<<<< HEAD
	 loadavg()
	 meminfo()
	 uptime()
=======

	go counter01()
	go counter02()
	go gauge01()
>>>>>>> func3
	fmt.Println("server started at port 9000")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
=======

>>>>>>> func3
}

