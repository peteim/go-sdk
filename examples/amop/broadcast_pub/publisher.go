package main

import (
	"encoding/hex"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/peteim/go-sdk/client"
	"github.com/peteim/go-sdk/conf"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("the number of arguments is not equal 3")
	}
	waitToSend := 5 * time.Second
	if len(os.Args) == 4 {
		i, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("parse to int failed: %v", err)
		}
		waitToSend = time.Duration(i) * time.Second
	}
	endpoint := os.Args[1]
	topic := os.Args[2]
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &conf.Config{IsHTTP: false, ChainID: 1, CAFile: "ca.crt", Key: "sdk.key", Cert: "sdk.crt", IsSMCrypto: false, GroupID: 1,
		PrivateKey: privateKey,
		NodeURL:    endpoint}
	c, err := client.Dial(config)
	if err != nil {
		log.Fatalf("init publisher failed, err: %v\n", err)
	}
	time.Sleep(waitToSend)

	message := "hello, FISCO BCOS, I am broadcast publisher!"
	for i := 0; i < 50; i++ {
		log.Printf("publish message: %s ", message+" "+strconv.Itoa(i))
		err = c.BroadcastAMOPMsg(topic, []byte(message+" "+strconv.Itoa(i)))
		time.Sleep(200 * time.Millisecond)
		if err != nil {
			log.Printf("PushTopicDataRandom failed, err: %v\n", err)
		}
	}
	c.BroadcastAMOPMsg(topic, []byte("Done"))
	c.Close()
}
