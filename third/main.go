package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/elazarl/go-bindata-assetfs"
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

func main() {

	
	fs := assetfs.AssetFS{
		Asset: Asset,
		AssetDir: AssetDir,
		AssetInfo: AssetInfo,
		}
	//http.Handle("/", http.FileServer(&fs))
	http.Handle("/static/bootstrap4/", http.FileServer(&fs))
	http.HandleFunc("/", IndexHandler)
	serverErr := http.ListenAndServe(":8085", nil)

	if nil != serverErr {
		log.Panic(serverErr.Error())
	}
}
