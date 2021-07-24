/*c
Copyright (c) 2021 Naxii.
https://github.com/Naxii-e/get_http_code
*/

package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Host struct {
	Disp string
	Url  string
}

type ReHost struct {
	Disp string    `json:"disp"`
	Url  string    `json:"url"`
	Code int       `json:"code"`
	Time time.Time `json:"time"`
}

func info(msg string) {
	fmt.Println("[INFO]", msg)
}

func ReadCsv(filename string) ([][]string, error) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func main() {
	WelcomeMsg := `
	    __  __________________            _______ ____  _   __
	   / / / /_  __/ ____/ __ \          / / ___// __ \/ | / /
	  / /_/ / / / / /   / / / /_______  / /\__ \/ / / /  |/ / 
	 / __  / / / / /___/ /_/ /_____/ /_/ /___/ / /_/ / /|  /  
	/_/ /_/ /_/  \____/_____/      \____//____/\____/_/ |_/   

	Ver 1.0 - BETA
	Author: Naxii	
	GitHub: https://github.com/Naxii-e
	Keybase: https://keybase.io/naxii_e
	
	SourceCode: https://github.com/Naxii-e/get_http_code
	License: MIT 
	
	Copyright (c) 2021 Naxii.

`
	var intOpt = flag.String("debug", "false", "output debug console")
	var intOpt2 = flag.String("f", "hosts.csv", "file path of csv")
	flag.Parse()
	DebugOptionMode := false
	if *intOpt == "true" {
		DebugOptionMode = true
	}
	fmt.Println(WelcomeMsg)
	if DebugOptionMode == true {
		info("！デバッグモードが有効です！")
	}
	in, err := ReadCsv(*intOpt2)
	if err != nil {
		log.Fatalln("csv ファイルが見つかりません。")
	}
	info("csv ファイルを読み込みました。")
	info("")
	info("---=取得を開始します=---")
	// 非同期処理へ対応・排他制御が必要
	var wg sync.WaitGroup
	var mu sync.Mutex
	var rehost []ReHost
	for _, host := range in {
		wg.Add(1)
		h := host
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			resu := new(ReHost)
			data := Host{
				Disp: h[0],
				Url:  h[1],
			}
			req, err := http.NewRequest("GET", data.Url, nil)
			req.Header.Set("User-Agent", "Mozilla/5.0 (Maenmo; Linux armv71; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 Fennec/10.0.1")
			res, err := http.DefaultClient.Do(req)
			var rescode int
			if err != nil {
				println("[FAIL]   ", data.Disp, "のURLを処理できませんでした。")
				rescode = -1
			} else {
				rescode = res.StatusCode
				fmt.Println("[_OK_]   ", data.Disp, "->", rescode)
			}
			resu.Disp = data.Disp
			resu.Url = data.Url
			rehost = append(rehost, ReHost{Disp: data.Disp, Url: data.Url, Code: rescode, Time: time.Now()})
			json.Marshal(rehost)
		}(&wg)
	}
	wg.Wait()
	result, _ := json.Marshal(rehost)
	var buf bytes.Buffer
	json.Indent(&buf, result, "", "  ")
	indentResult := buf.String()
	//fmt.Println(indentResult)
	info("---=取得を終了しました=---")
	info("")
	info("jsonファイルに書き出しています...")
	if DebugOptionMode == true {
		fmt.Println("==========BEGIN DEBUG==========")
		fmt.Printf("[JSON CONTENTS]\n%s\n", indentResult)
		fmt.Println("==========END DEBUG==========")
	}
	err = ioutil.WriteFile("http_response_results.json", []byte(indentResult), 0644)
	if err != nil {
		log.Fatalln("ファイル生成に失敗しました。")
	}
	info("jsonファイルの書き出しが完了しました。")
	info("10秒後に自動終了します...")
	time.Sleep(10 * time.Second)
}
