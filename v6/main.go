package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	botName := "super-long-name-for-bot"

	buf := bytes.NewBuffer(nil)
	// Increment
	checkErr(binary.Write(buf, binary.LittleEndian, uint8(230)))
	payload := buf.Bytes()

	finalURL := hex.EncodeToString(payload)
	fmt.Println(finalURL, "hash hex")

	finalURL = "https://" + botName + ".magicbots.org/" + finalURL
	fmt.Println(finalURL, "Final URL")
}
