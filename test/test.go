package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	payload, err := base64.StdEncoding.DecodeString("eyJhcHBfa2V5IjoiMjY0ZjUzYjM0YWZkNWY4OGI0YjVjNTIzMDY5Y2M3OTgiLCJhcHBfc2VjcmV0IjoiN2JiNWFmMWZiNDc2MDk0YzQxMzg4ODM0YzM1OGE0OWYiLCJleHAiOjE2MTg2NTQwNzEsImlzcyI6InNreSJ9")
	fmt.Println(string(payload), err)
}
