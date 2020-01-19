package main

import (
	"fmt"
	"four/testuuid"
	"os"
)

func main() {
	fmt.Println("args:", os.Args)
	//fmt.Println("server.port:", (strings.Split(os.Args[1],"="))[1])
	fmt.Println(testuuid.GetUUID())
	fmt.Println("test--"*3)
}
