package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RequestStruct struct {
	Url       string
	WhatApiIs string
	Headers   map[string]string
}

type ResultStruct struct {
	Response []byte
	FromAPI  string
}

var (
	HeadersCDN = map[string]string{
		"Content-Type":              "application/json",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language":           "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7",
		"Cache-Control":             "max-age=0",
		"Connection":                "keep-alive",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"sec-ch-ua":                 `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"Linux"`,
	}

	HeadersVIA = map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language":           "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7",
		"Cache-Control":             "max-age=0",
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}
)

func (rs *RequestStruct) GetUrlResult(ctx context.Context, cep string) (*ResultStruct, error) {

	urlInput := rs.Url

	correctUrl := strings.Replace(urlInput, "CEP_HERE", cep, 1)

	req, err := http.NewRequest("GET", correctUrl, nil)
	if err != nil {
		panic(fmt.Sprintf("error to create the new request, error: %s\n", err.Error()))
	}

	for k, v := range rs.Headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		panic(fmt.Sprintf("error to sends the request, error: %s\n", err.Error()))
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error to read the response body, error: %s", err.Error())
		return nil, err
	}

	r := ResultStruct{
		Response: bodyText,
		FromAPI:  rs.WhatApiIs,
	}

	return &r, nil

}
