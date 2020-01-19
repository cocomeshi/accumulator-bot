package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cocomeshi/accumulator-bot/internal"
)

func main() {
	fmt.Println("Test")
	apiKey, err := readKey()
	if err != nil {
		fmt.Println("APIアクセスキーの取得に失敗しました。")
	}
	internal.Fetch(apiKey)
}

func readKey() (string, error) {
	f, err := os.Open("access-key.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	return string(buf), err
}
