package http

import (
	"fmt"
	"net/http"
)

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	fmt.Printf("Pattern is: %s", pattern)
	http.DefaultServeMux.HandleFunc(pattern, handler)
}
