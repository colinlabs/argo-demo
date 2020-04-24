package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
		<div>
			<h2>GOOS: %s</h2>
			<h2>Version: %s</h2>
		</div>
	`
	content := fmt.Sprintf(tmpl, runtime.GOOS, runtime.Version())

	fmt.Fprintf(w, content)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Start web server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Error")
	}
}
