package main
import (
	"fmt"
	"time"
	"io"
	"net/http"
)

const Form = `
	     <html>
	     <body>
	     <form action = "#" method = "post" name = "test">
		<input type = "text" name = "num" />
		<input type = "submit" value = "submit" /> 
	     </form>
	     </body>
	     </html>
	     `

func Welcome(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome, Now is %s</h1>", time.Now());
}

func HandleForm(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-type", "text/html");
	switch request.Method {
		case "GET":
			io.WriteString(w, Form);
		case "POST":
			io.WriteString(w, request.FormValue("num"));
	}
}

func main() {
	http.HandleFunc("/welcome", Welcome);
	http.HandleFunc("/form", HandleForm);
	if err := http.ListenAndServe("localhost:80", nil); nil != err {
		fmt.Println(err.Error());
		panic("WebServer Start Failed!");
	}
}
