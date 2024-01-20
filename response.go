package main

import (
	"fmt"
	"net/http"
	"strings"
)

type response struct {
	status  int
	content string
}

func (res response) serialize() string {
	b := strings.Builder{}

	statusText := http.StatusText(res.status)

	b.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", res.status, statusText))
	b.WriteString("Connection: close\r\n")
	b.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n", len(res.content)))
	if res.content != "" {
		b.WriteString(res.content + "\r\n")
	}

	return b.String()
}

func (res response) serializeToBytes() []byte {
	return []byte(res.serialize())
}
