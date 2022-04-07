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
	memtotal_bytes_total = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Memory",
		Subsystem: "Memory",
		Name:      "memtotal_bytes_total",
		Help:      "Displaying the stats of memory",
	})
	memfree_bytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Memory",
		Subsystem: "Memory",
		Name:      "memfree_bytes",
		Help:      "Displaying the stats of memory",
	})
	memavailable_bytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Memory",
		Subsystem: "Memory",
		Name:      "memavailable_bytes",
		Help:      "Displaying the stats of memory",
	})
	buffers_byte = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Memory",
		Subsystem: "buffers",
		Name:      "buffers_byte",
		Help:      "Displaying the stats of memory ",
	})

	cached_bytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "Memory",
		Subsystem: "Cache",
		Name:      "cached_bytes",
		Help:      "Displaying the stats of memory ",
	})
	wall_clock_since_boot_bytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "uptime",
		Subsystem: "cputime",
		Name:      "wall_clock_since_boot_bytes",
		Help:      "wall clock since boot ",
	})
	idle_cputime_bytes_total = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "uptime",
		Subsystem: "cputime",
		Name:      "idle_cputime_bytes_total",
		Help:      "combined idle cputime ",
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

func collect_meminfo(){

			for {
					mi := parser.Get_mem_info()
					memfree_bytes.Set(mi["MemFree:"])
					memavailable_bytes.Set(mi["MemAvailable:"])
					memtotal_bytes_total.Set(mi["MemTotal:"])
					buffers_byte.Set(mi["Buffers:"])
					cached_bytes.Set(mi["Cached:"])
					time.Sleep(2 * time.Second)

			}



}

func collect_uptime(){
	
		for {
			ut :=parser.Get_uptime()
			wall_clock_since_boot_bytes.Set(ut["walk_clock"])
			idle_cputime_bytes_total.Set(ut["combined_idletime"])
			time.Sleep(2 * time.Second)
		}


}

func collect_loadavg(){
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
	
}

func init() { 
	prometheus.MustRegister(memavailable_bytes)
	prometheus.MustRegister(memfree_bytes)
	prometheus.MustRegister(memtotal_bytes_total)
	prometheus.MustRegister(cached_bytes)
	prometheus.MustRegister(buffers_byte)
	prometheus.MustRegister(wall_clock_since_boot_bytes)
	prometheus.MustRegister(idle_cputime_bytes_total)
	prometheus.MustRegister(process_gauge01_float64)
	prometheus.MustRegister(process_gauge02_float64)
	prometheus.MustRegister(process_gauge03_float64)
	prometheus.MustRegister(process_gauge04_float64)
	prometheus.MustRegister(process_gauge05_float64)
	prometheus.MustRegister(process_gauge06_float64)
}

func main() {
	go collect_loadavg()
	go collect_meminfo()
	go collect_uptime()
	fmt.Println("server started at port 9000")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
