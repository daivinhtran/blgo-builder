package server

import (
	"fmt"
	"net/http"
)

// Serve serves generated content
func Serve(outputPath string, port string) {
	fmt.Println("Serving", "http://localhost:"+port)
	http.Handle("/", http.FileServer(http.Dir(outputPath)))
	http.ListenAndServe(":"+port, nil)
}
