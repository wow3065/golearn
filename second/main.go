package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var (
	help bool

	h string
	p string
	c string
	U string
	P string
)

func init() {
	flag.BoolVar(&help, "help", false, "this help")
	flag.StringVar(&h, "h", "localhost", "server's ip or hostname.but must be resolve by dns server.")
	flag.StringVar(&p, "p", "22", "connect ")
	flag.StringVar(&c, "c", "/usr/bin/hostname", "linux command")
	flag.StringVar(&U, "U", "root", "user name")
	flag.StringVar(&P, "P", "password", "user password")

	// 另一种绑定方式
	//q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	/**
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")
	*/

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	}
	fmt.Println(*&h)
	sshtest()
}

func usage() {
	fmt.Fprintf(os.Stderr, `second version: second/1.0.1
Usage: second [-hpcUP] help

Options:
`)
	flag.PrintDefaults()
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
