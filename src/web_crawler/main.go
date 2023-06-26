package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	reAccount = `(账号|迅雷账号)(；|：)[0-9:]+(| )密码：[0-9a-zA-Z]+`
)

func GetAccountAndPassword(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http.Get error, err = ", err)
	}
	defer resp.Body.Close()
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read resp error, err = ", err)
	}
	str := string(dataBytes)

	re := regexp.MustCompile(reAccount)

	// match time, -1 means all
	results := re.FindAllStringSubmatch(str, -1)
	for _, result := range results {
		log.Println(result[0])
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	GetAccountAndPassword("https://www.ucbug.com/jiaocheng/63149.html?_t=1582307696")
}
