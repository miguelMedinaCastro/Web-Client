package main

import "fmt"

func main(){
	fmt.Println("HTTP Client")
	method := ReadInput("escolha (GET, POST, PUT, DELETE): ")
	url := ReadInput("url: ")

	headers := ReadHeaders()
	
	var body string
	if method != "GET" && method != "DELETE" {
		body = ReadBody()
	}

	err := SendRequest(method, url, headers, body)
	if err != nil {
		fmt.Println("erro: ", err)
	}
}
