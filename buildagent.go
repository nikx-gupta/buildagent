package buildagent

import "net/http"

func Run() {

	http.DefaultServeMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Something Invoked"))
	})

	http.ListenAndServe("::8080", nil)
}
