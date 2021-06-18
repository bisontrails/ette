package ethereum

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	cfg "github.com/itzmeanjan/ette/app/config"
)

func DebugTransaction(hash string, parameters string) (string, error) {
	client := &http.Client{}
	json := `{"id": 1, "method": "debug_traceTransaction", "params": ["` + hash + `", ` + parameters + `]}`
	log.Println("Request with:", json)
	jsonByte := []byte(json)
	req, _ := http.NewRequest("POST", cfg.Get("RPCUrl"), bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func DebugBlockByNumber(block string, parameters string) (string, error) {
	client := &http.Client{}
	json := `{"id": 1, "method": "debug_traceBlockByNumber", "params": ["` + block + `", ` + parameters + `]}`
	log.Println("Request with:", json)
	jsonByte := []byte(json)
	req, _ := http.NewRequest("POST", cfg.Get("RPCUrl"), bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
