package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ikshavaku/catalogue/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func LoadTestConfig() utils.GlobalConfig {
	utils.InitConfig()
	cfg := utils.GetConfig()
	return cfg
}

func CreateRequest(method, path string, body any, headers map[string]string) (*http.Request, error) {
	var requestJSON []byte
	var err error
	switch v := body.(type) {
	case []byte:
		requestJSON = v
	case string:
		requestJSON = []byte(v)
	default: // struct or other types
		requestJSON, err = json.Marshal(body)
	}

	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, path, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		request.Header.Add(key, val)
	}
	return request, nil
}

func LoadResponse(response io.Reader, target any) error {
	return json.NewDecoder(response).Decode(target)
}
