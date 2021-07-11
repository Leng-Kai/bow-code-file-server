package main

import (
    "log"
    "net/http"
    "os"

    "github.com/Leng-Kai/bow-code-file-server/pkg/download"
    // "github.com/Leng-Kai/bow-code-file-server/pkg/upload"

    _ "github.com/joho/godotenv/autoload"
)

var port string
var docs_path string

func main() {
    port = os.Getenv("PORT")
    docs_path = os.Getenv("DOCS_PATH")
    log.Println("port:", port)
    log.Println("docs_path:", docs_path)
    fs := http.FileServer(download.SecuredFileSystem{http.Dir(docs_path)})
    http.Handle("/files/", http.StripPrefix("/files", fs))

    // http.HandleFunc("/upload/course/{id}/block", upload.CreateBlockHandler)
    // http.HandleFunc("/upload/course/{id}/block/{bid}", upload.UpdateBlockHandler)

    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}