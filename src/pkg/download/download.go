package download

import (
    "log"
    "net/http"
    "strings"
)

type SecuredFileSystem struct {
    Fs http.FileSystem
}

func (sfs SecuredFileSystem) Open(path string) (http.File, error) {
    f, err := sfs.Fs.Open(path)
    if err != nil {
        return nil, err
    }

    s, err := f.Stat()
    if s.IsDir() {
        index := strings.TrimSuffix(path, "/") + "/index.html"
        if _, err := sfs.Fs.Open(index); err != nil {
            return nil, err
        }
    }

    return f, nil
}

func Handler2HandlerFunc(handler http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        log.Println(r)

        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Origin", "http://api.ramen-live.com:3000")
        w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        handler.ServeHTTP(w, r)
    }
}
