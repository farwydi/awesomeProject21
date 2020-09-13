package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/fnv"
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
	// Params to encrypt
	botName := "super-long-name-for-bot"
	var landingID uint32 = 12555
	var paramsID uint32 = 46587

	buf := bytes.NewBuffer(nil)
	// Landing ID
	checkErr(binary.Write(buf, binary.LittleEndian, landingID))
	fmt.Println(base64.RawURLEncoding.EncodeToString(buf.Bytes()), "landingID")
	// Params ID
	checkErr(binary.Write(buf, binary.LittleEndian, paramsID))
	fmt.Println(base64.RawURLEncoding.EncodeToString(buf.Bytes()), "landingID+paramsID")
	payload := buf.Bytes()

	// To hash
	h := fnv.New64()
	checkWriteErr(h.Write(payload))
	payload = h.Sum(nil)

	finalURL := base64.RawURLEncoding.EncodeToString(payload)
	fmt.Println(finalURL, "hash base64")

	fmt.Println(hex.EncodeToString(payload), "hash hex")

	finalURL = "https://" + botName + ".magicbots.org/lp-" + finalURL
	fmt.Println(finalURL, "Final URL")
}
