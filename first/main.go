package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println("CPU数量：" + strconv.Itoa(runtime.NumCPU()))
	fmt.Println("系统类型：" + runtime.GOOS)
	fmt.Println("处理器架构：" + runtime.GOARCH)
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("主机名：" + hostname)
	}

}
