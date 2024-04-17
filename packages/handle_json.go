package packages

import (
	"encoding/json"
	"os"
	"log"
	"io"
	"net/http"
	"github.com/joho/godotenv"
)

type Pypilib struct {
	SourceRank int    `json:"sourceRank"`
	PkgName    string `json:"pkgName"`
}

type ApiRes struct {
	Repository_url				string	`json:"repository_url"`
	Forks						int		`json:"forks"`
	Stars						int		`json:"stars"`
	Latest_release_published_at	string	`json:"latest_release_published_at"`
}

type Combined struct {
	Pypilib
	ApiRes
}


func CallApi(url string) ApiRes {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("LoadEnv Error=%s", err.Error())
	}
	API_KEY := os.Getenv("API_KEY")
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("NewRequest err=%s", err.Error())
	}

	// クエリの設定
	q := req.URL.Query()
	q.Add("api_key", API_KEY)
	req.URL.RawQuery = q.Encode()

	// ヘッダーの設定(必要に応じて)
	// req.Header.Add("Content-Type", "application/json")

	// http.requestの送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Client.Do err=%s", err.Error())
	}
	defer resp.Body.Close()

	// レスポンスの内容表示
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll err=%s", err.Error())
	}

	response := ApiRes{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("json.Unmarshal err=%s", err.Error())
	}

	return response
}

func LoadJson(filePath string) []Pypilib {
	// JSONファイルの読み込み
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// JSONデータをパースして構造体に格納
	var packages []Pypilib
	if err := json.Unmarshal(jsonData, &packages); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// パッケージの情報を表示
	// fmt.Println(packages)
	return packages
}

func WriteOutJson(outPath string, jsonPtr *[]Combined) {

}
