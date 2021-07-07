package main

import (
    "log"
    "net/http"
    "os"
	"strings"

    _ "github.com/joho/godotenv/autoload"
)

type securedFileSystem struct {
    fs http.FileSystem
}

func (sfs securedFileSystem) Open(path string) (http.File, error) {
    f, err := sfs.fs.Open(path)
    if err != nil {
        return nil, err
    }

    s, err := f.Stat()
    if s.IsDir() {
        index := strings.TrimSuffix(path, "/") + "/index.html"
        if _, err := sfs.fs.Open(index); err != nil {
            return nil, err
        }
    }

    return f, nil
}

func main() {
    port := os.Getenv("PORT")
    docs_path := os.Getenv("DOCS_PATH")
    log.Println("port:", port)
    log.Println("docs_path:", docs_path)
    fs := http.FileServer(securedFileSystem{http.Dir(docs_path)})
    http.Handle("/", http.StripPrefix("/", fs))
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}