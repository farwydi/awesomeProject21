package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"hash/adler32"
	"io"
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
	keyHmac := make([]byte, 32)
	checkWriteErr(rand.Read(keyHmac))
	keyAes := make([]byte, 32)
	checkWriteErr(rand.Read(keyAes))

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
	// Hash bot name
	checkErr(binary.Write(buf, binary.LittleEndian, adler32.Checksum([]byte(botName))))
	fmt.Println(base64.RawURLEncoding.EncodeToString(buf.Bytes()), "landingID+paramsID+Checksum")
	payload := buf.Bytes()

	// aes crypto
	block, err := aes.NewCipher(keyAes)
	checkErr(err)
	aesgcm, err := cipher.NewGCM(block)
	checkErr(err)
	nonce := make([]byte, aesgcm.NonceSize())
	checkWriteErr(io.ReadFull(rand.Reader, nonce))
	payload = aesgcm.Seal(nil, nonce, payload, nil)

	finalURL := base64.RawURLEncoding.EncodeToString(payload)
	fmt.Println(finalURL, "landingID+paramsID+Checksum+aes")

	finalURL = "https://" + botName + ".magicbots.org/lp-" + finalURL
	fmt.Println(finalURL, "Final URL")
}
