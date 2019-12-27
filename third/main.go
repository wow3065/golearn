package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/elazarl/go-bindata-assetfs"
	"encoding/json"
	"golang.org/x/crypto/ssh"
	"bytes"
)

// index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := Asset("static/index.html") // 根据地址获取对应内容
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err := template.New("tpl").Parse(string(bytes))
	_ = t.ExecuteTemplate(
		w,
		"index",
		map[string]interface{}{"PageTitle": "首页", "Name": "sqrt_cat", "Age": 25},
	)
}

// GetCommandHandler
func GetCommandHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.PostFormValue("ip")
	port := r.PostFormValue("port")
	user := r.PostFormValue("user")
	password := r.PostFormValue("password")
	command := r.PostFormValue("command")
	
	if(ip == ""  ){
		ip = "localhost"
	}

	if(port == "" ){
		port = "22"
	}

	if(user == "" ){
		user = "root"
	}

	if(password == "" ){
		password = ""
	}

	if(command == "" ){
		command = "hostname"
	}

	result := make(map[string]interface{})
	result["data"] = ""

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.FixedHostKey(ssh.PublicKey),
	}
	client, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Panicln("Failed to dial: ", err)
		w.WriteHeader(500)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Panicln("Failed to create session: ", err)
		w.WriteHeader(500)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		log.Panicln("Failed to run: " + err.Error())
		w.WriteHeader(500)
	}else{
		result["data"] = b.String()
	}
	
	rJson, err:= json.Marshal(result)
	if err != nil {
		log.Panicln("error:", err)
	}

	w.Write(rJson)
}

func main() {

	fs := assetfs.AssetFS{
		Asset: Asset,
		AssetDir: AssetDir,
		AssetInfo: AssetInfo,
		}
	//http.Handle("/", http.FileServer(&fs))
	http.Handle("/static/bootstrap4/", http.FileServer(&fs))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/getCommand", GetCommandHandler)
	serverErr := http.ListenAndServe(":8085", nil)

	if nil != serverErr {
		log.Panic(serverErr.Error())
	}
}
