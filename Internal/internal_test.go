//go:build unit
// +build unit

package internal_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	internal "github.com/ArtusC/multithreading/Internal"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	testCase := []struct {
		name            string
		reqStruct       internal.RequestStruct
		resultExpected  map[string]interface{}
		fromApiExpected string
		expectedError   error
	}{
		{
			name:      "test wirh cdn api",
			reqStruct: internal.RequestStruct{Url: "https://cdn.apicep.com/file/apicep/CEP_HERE.json", WhatApiIs: "CDN", Headers: internal.HeadersCDN},
			resultExpected: map[string]interface{}{
				"address":    "Servidão Martinho Leandro dos Santos",
				"city":       "Florianópolis",
				"code":       "88040-050",
				"district":   "Pantanal",
				"ok":         true,
				"state":      "SC",
				"status":     200.0,
				"statusText": "ok",
			},
			fromApiExpected: "CDN",
			expectedError:   nil,
		},
		{
			name:      "test with viacep api",
			reqStruct: internal.RequestStruct{Url: "http://viacep.com.br/ws/CEP_HERE/json/", WhatApiIs: "VIACEP", Headers: internal.HeadersVIA},
			resultExpected: map[string]interface{}{
				"cep":         "88040-050",
				"logradouro":  "Servidão Martinho Leandro dos Santos",
				"complemento": "",
				"bairro":      "Pantanal",
				"localidade":  "Florianópolis",
				"uf":          "SC",
				"ibge":        "4205407",
				"gia":         "",
				"ddd":         "48",
				"siafi":       "8105",
			},
			fromApiExpected: "VIACEP",
			expectedError:   nil,
		},
	}

	for _, test := range testCase {

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		res, err := test.reqStruct.GetUrlResult(ctx, "88040-050")

		convertedRes, errorUn := unmarshalResult(res.Response)
		fmt.Println(convertedRes)
		assert.Nil(t, errorUn)

		if res.FromAPI == "CDN" && convertedRes["status"] != 200.0 {
			if convertedRes["status"] == 403.0 {
				fmt.Println("Erro 403")
				resErrorExpected := map[string]interface{}{"code": "download_cap_exceeded", "message": "Cannot download file, download bandwidth or transaction (Class B) cap exceeded. See the Caps & Alerts page to increase your cap.", "status": 403.0}
				assert.Equal(t, resErrorExpected, convertedRes)
			}

			if convertedRes["status"] == 429.0 {
				fmt.Println("Erro 429")
				resErrorExpected := map[string]interface{}{"status": 429.0, "ok": false, "message": "Blocked by flood cdn", "statusText": "bad_request"}

				assert.Equal(t, resErrorExpected, convertedRes)
			}

		} else {

			assert.Equal(t, test.resultExpected, convertedRes)

			assert.Equal(t, test.fromApiExpected, res.FromAPI)

			assert.Equal(t, test.expectedError, err)

		}

	}
}

func unmarshalResult(result []byte) (data map[string]interface{}, err error) {
	errUn := json.Unmarshal(result, &data)
	if errUn != nil {
		errorMsg := fmt.Sprintf("could not unmarshal json: %s\n", errUn)
		return nil, errors.New(errorMsg)
	}
	return data, nil
}
