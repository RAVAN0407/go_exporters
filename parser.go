package parser

import (
	"fmt"
	"os/exec"
	"strings"
	"strconv"
	"regexp"
)

func Get_mem_info() map[string]float64 { 	//return map
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
