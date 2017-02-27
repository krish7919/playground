package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {





	postData := `{
  "asset": {
    "data": {
      "dummy": "dummy"
    }
  },
  "id": "7fa883f47d0cbcd6e238747f245d98eb84d7df10ff794c97668633ec19070942",
  "inputs": [{
    "fulfillment": "cf:4:4dby6UPos5lSZFTP-DEJQMXXvIgsB0mu-9wob5WCJAk9fLb02CnGVwbq74kGPuNkXExWw70LQh3QzDnHT_tJZIJQ7_7B82UyLAccFl6OugzuqVPmIsYX12XYBR_VH64H",
    "fulfills": null,
    "owners_before": [
      "GCaqN8h8jHPY34Lb2H2dVaojSGRBSTaqxRHzAinXtUVn"
    ]
  }],
  "metadata": {
    "dummy": "dummy"
  },
  "operation": "CREATE",
  "outputs": [{
    "amount": 1,
    "condition": {
      "details": {
        "bitmask": 32,
        "public_key": "GCaqN8h8jHPY34Lb2H2dVaojSGRBSTaqxRHzAinXtUVn",
        "signature": null,
        "type": "fulfillment",
        "type_id": 4
      },
      "uri": "cc:4:20:4dby6UPos5lSZFTP-DEJQMXXvIgsB0mu-9wob5WCJAk:96"
    },"public_keys": [
      "GCaqN8h8jHPY34Lb2H2dVaojSGRBSTaqxRHzAinXtUVn"
    ]
  }],
  "version": "0.9"
}`

	//resp, err := SendHttpRequest("http://localhost:59984/api/v1/transactions",
	//	"POST", "application/json", []byte(postData))

	resp, err := PostTx(postData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}

func SendHttpRequest(url, method, contentType string, data []byte) ([]byte,
	error) {
	var reqBody io.Reader
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	client := &http.Client{
		Transport: transport,
	}
	if data == nil {
		reqBody = nil
	} else {
		reqBody = bytes.NewReader(data)
	}
	httpReq, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		fmt.Println("debug", fmt.Sprintf("HTTP req create: ", err))
		return nil, err
	}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("debug", fmt.Sprintf("GET/POST failed: %s, ResponseCode: %s", err,
			httpResp.Status))
		return nil, err
	}
	defer httpResp.Body.Close()
	if resp, err := ioutil.ReadAll(httpResp.Body); err != nil {
		fmt.Println("debug", fmt.Sprintf("HTTP resp read: %s", err))
		return nil, err
	} else {
		return resp, nil
	}
}

func PostTx(tx string) (string, error) {
    url := "http://localhost:59984/api/v1/transactions"
	kv := map[string]string{"content_type": "application/json"}
	//kv := map[string]string{"content-type": "application/json"}

	txBytes := bytes.NewBufferString(tx)

	response, err := HttpPostRequest(url, txBytes, kv)
	if err != nil {
		return "", err
	}
	fmt.Println(response)
	rd, err := ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(rd))
	// data := make(map[string]interface{})
	// UnmarshalJSON(rd, data)
	// id := data["id"].(string)
	id := ""
	return id, nil
}

func ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

func HttpPostRequest(url string, body io.Reader, kv map[string]string) (*http.Response, error) {
	return HttpRequest(http.MethodPost, url, body, kv)
}
func HttpRequest(method, url string, body io.Reader, kv map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range kv {
		req.Header.Set(k, v)
	}
	cli := new(http.Client)
	return cli.Do(req)
}
