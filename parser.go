package parser

import (
	"fmt"
	"os/exec"
	"strings"
)

func get_mem_info() [5][]string {
	prg := "cat"
	arg1 := "/proc/meminfo"
	var arr [5][]string
	cmd := exec.Command(prg, arg1)
	res, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	split := strings.Split(string(res), "\n")
	for i := 0; i < len(split); i++ {
		arr[i] = strings.Split(split[i], ":")
	}

	return arr
}
