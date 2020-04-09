package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cocomeshi/accumulator-bot/internal"
)

func main() {
	apiKey, err := readKey()
	if err != nil {
		fmt.Println("APIアクセスキーの取得に失敗しました。")
	}
	internal.Exec(apiKey)
	internal.AdditionalUpdate(apiKey)
}

func readKey() (string, error) {
	f, err := os.Open("googleapi-accesskey.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	return string(buf), err
}
