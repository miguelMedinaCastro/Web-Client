package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SendRequest(method, url string, headers map[string]string, body string) error {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("\n= RESPOSTA =")
	fmt.Println("Status:", resp.Status)

	fmt.Println("\n- Headers -")
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
	}

	fmt.Println("\n- Corpo -")
	bodyResp, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bodyResp))

	return nil
}

