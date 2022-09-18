package parser

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	"regexp"
	"log"
	"encoding/json"
)
type Mpstat_Details struct {
	Sysstat struct {
		Hosts []struct {
			Nodename     string `json:"nodename"`
			Sysname      string `json:"sysname"`
			Release      string `json:"release"`
			Machine      string `json:"machine"`
			NumberOfCpus int    `json:"number-of-cpus"`
			Date         string `json:"date"`
			Statistics   []struct {
				Timestamp string `json:"timestamp"`
				CPULoad   []struct {
					CPU    string  `json:"cpu"`
					Usr    float64 `json:"usr"`
					Nice   float64 `json:"nice"`
					Sys    float64 `json:"sys"`
					Iowait float64 `json:"iowait"`
					Irq    float64 `json:"irq"`
					Soft   float64 `json:"soft"`
					Steal  float64 `json:"steal"`
					Guest  float64 `json:"guest"`
					Gnice  float64 `json:"gnice"`
					Idle   float64 `json:"idle"`
				} `json:"cpu-load"`
			} `json:"statistics"`
		} `json:"hosts"`
	} `json:"sysstat"`
}

type Iostat_Details struct {
	Sysstat struct {
		Hosts []struct {
			Nodename     string `json:"nodename"`
			Sysname      string `json:"sysname"`
			Release      string `json:"release"`
			Machine      string `json:"machine"`
			NumberOfCpus int    `json:"number-of-cpus"`
			Date         string `json:"date"`
			Statistics   []struct {
				Disk []struct {
					DiskDevice string  `json:"disk_device"`
					Tps        float64 `json:"tps"`
					KBReadS    float64 `json:"kB_read/s"`
					KBWrtnS    float64 `json:"kB_wrtn/s"`
					KBDscdS    float64 `json:"kB_dscd/s"`
					KBRead     int     `json:"kB_read"`
					KBWrtn     int     `json:"kB_wrtn"`
					KBDscd     int     `json:"kB_dscd"`
				} `json:"disk"`
			} `json:"statistics"`
		} `json:"hosts"`
	} `json:"sysstat"`
}

func Get_mem_info() map[string]float64 { 	
	prg := "cat"
	arg1 := "/proc/meminfo"
	var arr [5][]string
	var s [5]string
	var mem = make(map[string]float64)
	cmd := exec.Command(prg, arg1)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	split := strings.Split(string(res), "\n")
	space := regexp.MustCompile(`\s+`)
	for i := 0; i < 5; i++ {
		s[i] = space.ReplaceAllString(split[i], " ")
		arr[i] = strings.Split(s[i], " ")
		mem[arr[i][0]],err=strconv.ParseFloat(arr[i][1],64)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return mem

}

func Get_uptime() map[string]float64{
	prg := "cat"
	arg1 := "/proc/uptime"
	cmd := exec.Command(prg, arg1)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	var uptime = make(map[string]float64)  
	split := strings.Split(string(res), "\n")
	split2:= strings.Split(split[0]," ")
	uptime["walk_clock"],err=strconv.ParseFloat(split2[0],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	uptime["combined_idletime"],err=strconv.ParseFloat(split2[1],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	
	return uptime
}

func Get_loadavg() map[string]float64{
	prg := "cat"
	arg1 := "/proc/loadavg"
	var lavg = make(map[string]float64)
	cmd := exec.Command(prg, arg1)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	split := strings.Split(string(res), " ")
	split2 :=strings.Split(split[3],"/")
	split3 :=strings.Split(split[4],"\n")
	lavg["process1"],err=strconv.ParseFloat(split[0],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	lavg["process2"],err=strconv.ParseFloat(split[1],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	lavg["process3"],err=strconv.ParseFloat(split[2],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	lavg["process4"],err=strconv.ParseFloat(split2[0],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	lavg["process5"],err=strconv.ParseFloat(split2[1],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	lavg["process6"],err=strconv.ParseFloat(split3[0],64)
	if err !=nil{
		fmt.Println(err.Error())
	}
	return lavg
}

func Get_mpstat() Mpstat_Details{
	prg := "mpstat"
	arg1 := "-o"
	arg2 := "JSON"
	cmd := exec.Command(prg, arg1,arg2)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	var mp_stat Mpstat_Details
	err = json.Unmarshal(res, &mp_stat)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	return mp_stat

}

func Get_iostat() Iostat_Details{
	prg := "mpstat"
	arg0:= "-d"
	arg1 := "-o"
	arg2 := "JSON"
	cmd := exec.Command(prg,arg0,arg1,arg2)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	var io_stat Iostat_Details
	err = json.Unmarshal(res, &io_stat)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	return io_stat


}