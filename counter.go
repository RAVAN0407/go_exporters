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

	}()
}

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

func main() {
	 loadavg()
	 meminfo()
	 uptime()
	fmt.Println("server started at port 9000")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
