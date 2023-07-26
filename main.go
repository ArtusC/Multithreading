package main

import (
	"context"
	"fmt"
	"time"

	internal "github.com/ArtusC/multithreading/Internal"
)

func main() {
	fmt.Println("Starting challange...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	type Message struct {
		API *internal.ResultStruct
		Msg string
	}

	c1 := make(chan Message)
	c2 := make(chan Message)

	// APICep
	go func() {
		for {
			reqStructCDN := internal.RequestStruct{Url: "https://cdn.apicep.com/file/apicep/CEP_HERE.json", WhatApiIs: "CDN", Headers: internal.HeadersCDN}
			res, err := reqStructCDN.GetUrlResult(ctx, "88040-050")
			if err != nil {
				panic(err)
			}
			msg := Message{res, "Hello from CDN API"}
			c1 <- msg
		}
	}()

	// VIACep
	go func() {
		for {
			reqStructVIA := internal.RequestStruct{Url: "http://viacep.com.br/ws/CEP_HERE/json/", WhatApiIs: "VIACEP", Headers: internal.HeadersVIA}
			res, err := reqStructVIA.GetUrlResult(ctx, "88040-050")
			if err != nil {
				panic(err)
			}
			msg := Message{res, "Hello from VIA Cep API"}
			c2 <- msg
		}
	}()

myloop:
	for {
		select {
		case msg := <-c1: // APICep
			fmt.Printf("Received from CDN:\nMsg: %s\nResult:%v", msg.Msg, string(msg.API.Response))
			break myloop
		case msg := <-c2: // VIACep
			fmt.Printf("Received from VIACep:\nMsg: %s\nResult:%v", msg.Msg, string(msg.API.Response))
			break myloop
		case <-ctx.Done():
			panic(fmt.Sprintf("Request timed out: %v\n", ctx.Err()))
		}
	}
}
