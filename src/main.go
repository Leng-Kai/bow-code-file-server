package main

import (
    "log"
    "net/http"
    "os"

    "github.com/Leng-Kai/bow-code-file-server/pkg/download"
    // "github.com/Leng-Kai/bow-code-file-server/pkg/upload"

    _ "github.com/joho/godotenv/autoload"
    // "github.com/gorilla/mux"
    // "github.com/rs/cors"
)

var port string
var docs_path string

func main() {
    port = os.Getenv("PORT")
    docs_path = os.Getenv("DOCS_PATH")
    log.Println("port:", port)
    log.Println("docs_path:", docs_path)

    fs := http.FileServer(download.SecuredFileSystem{http.Dir(docs_path)})
    // http.Handle("/files/", http.StripPrefix("/files", fs))
    http.HandleFunc("/files/", download.Handler2HandlerFunc(http.StripPrefix("/files", fs)))

    // http.HandleFunc("/upload/course/{id}/block", upload.CreateBlockHandler)
    // http.HandleFunc("/upload/course/{id}/block/{bid}", upload.UpdateBlockHandler)

    // r := mux.NewRouter()
    // r.HandleFunc("/files/", download.Handler2HandlerFunc(http.StripPrefix("/files", fs))).Methods("GET")
    // http.Handle("/", r)

    // c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000"},        // All origins
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // Allowing only get, just an example
	// 	AllowCredentials: true,
	// })

    err := http.ListenAndServe(":"+port, nil)
    // err := http.ListenAndServe(":"+port, c.Handler(r))
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}