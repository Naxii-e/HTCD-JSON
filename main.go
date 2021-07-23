package main

import (
	"C"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Host struct {
	Disp string
	Url  string
}

type ReHost struct {
	Disp string
	Url  string
	Code int
}

func GetHttpResponse() *C.char {
	in, err := ReadCsv("hosts.csv")
	if err != nil {
		panic(err)
	}

	fmt.Println("---=取得を開始します=---")
	resu := new(ReHost)

	for _, host := range in {
		data := Host{
			Disp: host[0],
			Url:  host[1],
		}

		req, err := http.NewRequest("GET", data.Url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Maenmo; Linux armv71; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 Fennec/10.0.1")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
			log.Fatal("エラーが", data.Disp, "で発生しましたが、飛ばしました。")
		}
		resu.Disp = data.Disp
		resu.Url = data.Url
		resu.Code = res.StatusCode
		fmt.Println(data.Disp, "を取得しました:", res.StatusCode)
	}
	//jsonでpyに渡そうとしたけど失敗
	fmt.Println("---=取得を終了しました=---")
	return nil
}

func ReadCsv(filename string) ([][]string, error) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func main() {
	GetHttpResponse()
}
