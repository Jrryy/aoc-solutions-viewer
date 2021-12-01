package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    _, err := os.Open("./aocs" + r.URL.Path + "/code.py")
    if err != nil {
        files, err := ioutil.ReadDir("./aocs" + r.URL.Path)
        if err != nil {
            panic(err)
        }
        var filesList []string
        for _, file := range files {
            filesList = append(filesList, file.Name())
        }
        if err := json.NewEncoder(w).Encode(filesList[1:]); err != nil {
            panic(err)
        }
    } else {
        content, err := ioutil.ReadFile("./aocs" + r.URL.Path + "/code.py")
        if err != nil {
            panic(err)
        }
        _, err = w.Write(content)
        if err != nil {
            panic(err)
        }
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
