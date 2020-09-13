package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
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
	// Params ID
	checkErr(binary.Write(buf, binary.LittleEndian, paramsID))
	// Hash bot name
	checkErr(binary.Write(buf, binary.LittleEndian, adler32.Checksum([]byte(botName))))
	fmt.Println(base64.RawURLEncoding.EncodeToString(buf.Bytes()), "Payload")
	payload := buf.Bytes()

	// Add verify mac
	mac := hmac.New(fnv.New128, keyHmac)
	checkWriteErr(mac.Write(payload))
	payload = mac.Sum(payload)

	// aes crypto
	block, err := aes.NewCipher(keyAes)
	checkErr(err)
	nonce := make([]byte, 12)
	checkWriteErr(io.ReadFull(rand.Reader, nonce))
	aesgcm, err := cipher.NewGCM(block)
	checkErr(err)
	payload = aesgcm.Seal(nil, nonce, payload, nil)

	// Add sum
	buf = bytes.NewBuffer(payload)
	checkErr(binary.Write(buf, binary.LittleEndian, crc32.ChecksumIEEE(buf.Bytes())))
	payload = buf.Bytes()

	finalURL := base64.RawURLEncoding.EncodeToString(payload)
	fmt.Println(finalURL, "Payload with protect and hmac")

	finalURL = "https://" + botName + ".magicbots.org/lp-" + finalURL
	fmt.Println(finalURL, "Final URL")
}
