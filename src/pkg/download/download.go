package download

import (
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
        handler.ServeHTTP(w, r)
    }
}