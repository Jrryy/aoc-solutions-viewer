package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strings"
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

func executor(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    splitPath := strings.Split(r.URL.Path, "/")
    day := splitPath[len(splitPath)-1]
    codeFilePath := "./aocs/AoC-2021/" + day + "/code.py"
    inputFilePath := "./aocs/AoC-2021/" + day + "/input.txt"

    // Does the code for thay day exist?
    _, err := os.Open(codeFilePath)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
    } else {
        // Take the uploaded file
        input, inputHeader, err := r.FormFile("File")
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        // Truncate or create the input.txt file
        f, err := os.Create(inputFilePath)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        inputBytes := make([]byte, inputHeader.Size)
        // Read the contents of the uploaded file and close it
        _, err = input.Read(inputBytes)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        err = input.Close()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        // Dump those contents into the recreated input.txt
        _, err = f.Write(inputBytes)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        if err = f.Close(); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
		// Create the command to execute the python code with the new input
        cmd := exec.Command(
            "sh",
            "-c",
            fmt.Sprintf("cd ./aocs/AoC-2021/%s && time python3 code.py", day),
        )
		// Execute and take the output
        output, err := cmd.CombinedOutput()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
		// We only want the solutions and the real execution time
        splitOutput := strings.Split(fmt.Sprintf("%s", output), "\n")
        part1 := splitOutput[0]
        part2 := splitOutput[1]
        timeSpent := "Time spent: " + splitOutput[2][5:]
        if _, err := w.Write([]byte(strings.Join([]string{part1, part2, timeSpent}, "\n"))); err != nil {
            panic(err)
        }
        w.WriteHeader(http.StatusOK)
    }
}

func main() {
    http.HandleFunc("/exec/", executor)
    http.HandleFunc("/AoC-2021/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
