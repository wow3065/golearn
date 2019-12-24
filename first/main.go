package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
    "flag"
	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("CPU数量：" + strconv.Itoa(runtime.NumCPU()))
	fmt.Println("系统类型：" + runtime.GOOS)
	fmt.Println("处理器架构：" + runtime.GOARCH)
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("主机名：" + hostname)
	}
	fmt.Println("this is person test.")
	
	sshtest()
}

func sshtest() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("123123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.FixedHostKey(ssh.PublicKey),
	}
	client, err := ssh.Dial("tcp", "10.200.72.26:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls -l /home"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
