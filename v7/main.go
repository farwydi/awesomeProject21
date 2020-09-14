package main

import (
	"crypto/rand"
	"fmt"
	"github.com/eknkc/basex"
	"log"
)

func checkWriteErr(_ int, err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	enc, err := basex.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	checkErr(err)

	botName := "super-long-name-for-bot"

	payload := make([]byte, 2)
	checkWriteErr(rand.Read(payload))

	finalURL := enc.Encode(payload)
	fmt.Println(finalURL, "hash enc")

	finalURL = "https://" + botName + ".magicbots.org/" + finalURL
	fmt.Println(finalURL, "Final URL")
}
