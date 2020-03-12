package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var rootPath = ""

var numToEng = map[int]string{1: "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "G",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z"}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		// 变更斜杠
		path = strings.TrimPrefix(strings.ReplaceAll(path, "\\", "/"), rootPath)
		//确定文件或者文件夹深度
		lenPath := len(strings.Split(path, "/"))
		//重新拼接文件路径
		if info.IsDir() {
			path += "<dir<" + strconv.Itoa(lenPath)
		} else {
			path += "<file<" + strconv.Itoa(lenPath)
		}

		*files = append(*files, path)
		return nil
	}
}

func main() {

	var files []string
	pwd, _ := os.Getwd()
	root := filepath.ToSlash(pwd)

	//root := "D:/Work/businessProject/梳理实验室培训材料/traindata"
	rootPath = root + "/"

	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}

	eFiles := files[1:]
	f := excelize.NewFile()
	index := f.NewSheet("index")
	f.DeleteSheet("Sheet1")

	var indexColumn = 2
	var sheetColumn = 2
	style, err := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}]}`)
	for _, iFile := range eFiles {
		sFile := strings.Split(iFile, "<")
		lengthA, _ := strconv.ParseInt(sFile[2], 10, 64)
		if sFile[1] == "dir" && lengthA == 1 {
			f.SetCellValue("index", "B"+strconv.Itoa(indexColumn), sFile[0])
			f.SetCellStyle("index", "B"+strconv.Itoa(indexColumn), "B"+strconv.Itoa(indexColumn), style)
			sheetColumn = 2
		} else if lengthA == 2 {
			if sFile[1] == "dir" {
				cellString := strings.Split(sFile[0], "/")
				f.NewSheet(cellString[len(cellString)-1])
				f.SetCellValue("index", "C"+strconv.Itoa(indexColumn), cellString[len(cellString)-1])
				f.SetCellHyperLink("index", "C"+strconv.Itoa(indexColumn), cellString[len(cellString)-1]+"!A1", "Location")
			} else {
				cellString := strings.Split(sFile[0], "/")
				f.SetCellValue("index", "C"+strconv.Itoa(indexColumn), cellString[len(cellString)-1])
				f.SetCellHyperLink("index", "C"+strconv.Itoa(indexColumn), sFile[0], "External")
			}
			f.SetCellStyle("index", "C"+strconv.Itoa(indexColumn), "C"+strconv.Itoa(indexColumn), style)
			f.SetCellStyle("index", "B"+strconv.Itoa(indexColumn), "B"+strconv.Itoa(indexColumn), style)
			indexColumn += 1
			sheetColumn = 2
		} else if lengthA > 2 {
			cellString := strings.Split(sFile[0], "/")
			f.SetCellHyperLink(cellString[1], numToEng[len(cellString)-1]+strconv.Itoa(sheetColumn), sFile[0], "External")
			for i := 2; i < len(cellString); i++ {
				for j := 2; j < sheetColumn+1; j++ {
					f.SetCellStyle(cellString[1], numToEng[i]+strconv.Itoa(j), numToEng[i]+strconv.Itoa(j), style)
				}
			}
			if sFile[1] == "dir" {
				f.SetCellValue(cellString[1], numToEng[len(cellString)-1]+strconv.Itoa(sheetColumn), cellString[len(cellString)-1])
			} else {
				f.SetCellValue(cellString[1], numToEng[len(cellString)-1]+strconv.Itoa(sheetColumn), cellString[len(cellString)-1])
				sheetColumn += 1
			}
		}
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs(rootPath + "index.xlsx"); err != nil {
		fmt.Println(err.Error())
	}

}
