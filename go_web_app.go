package main

import (
	"net/http"
	//"io/ioutil"
	//"bufio"
	//"os"
	//"strings"
)

//func main() {
//	//http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){  //passes anonymous function
//	//	w.Write([]byte("hello World"))
//	//})
//	http.Handle("/", new(MyHandler))
//
//	http.ListenAndServe(":8000", nil) //second arg enables multiplexing
//}
//
//type MyHandler struct {
//	http.Handler
//}
//
//func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	path := "public/" + req.URL.Path
//	f, err := os.Open(path)
//	//data, err := ioutil.ReadFile(string(path)) unbuffered reading of content, huge memory hog
//
//	if err == nil {
//		bufferedReader := bufio.NewReader(f)
//		var contentType string
//		if strings.HasSuffix(path, ".css") {
//			contentType = "text/css"
//		} else if strings.HasSuffix(path, ".html") {
//			contentType = "text/html"
//		} else if strings.HasSuffix(path, ".js") {
//			contentType = "application/javascript"
//		} else if strings.HasSuffix(path, ".png") {
//			contentType = "image/png"
//		} else if strings.HasSuffix(path, ".mp4") {
//			contentType = "video/mp4"
//		} else {
//			contentType = "text/plain"
//		}
//		w.Header().Add("Content-Type", contentType)
//		bufferedReader.WriteTo(w)
//		//w.Write(data)
//	} else {
//		w.Write([]byte("404 - " + http.StatusText(404)))
//	}
//
//}

func main(){
	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}
