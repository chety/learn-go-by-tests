package dependency_injection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler says Hello, world over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

//func main() {
//	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
//}
