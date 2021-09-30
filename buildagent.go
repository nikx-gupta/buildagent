package buildagent

import (
	"fmt"
	"net/http"
)

func Run() error {

	http.DefaultServeMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Something Invoked"))
	})

	return http.ListenAndServe("::8080", nil)
}

func main() {
	fmt.Println("This is buildagent")
}
