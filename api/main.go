package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "api/src/config"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    response := createResponse(r)

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(response))
}

func createResponse(r *http.Request) string {
    method := r.Method
    url := r.URL
    query := r.URL.Query()

    data := routes.Route(url.Path)
    responseData := map[string]interface{} {
        "data": data,
        "metadata": map[string]interface{} {
            "method": method,
            "url": url.Path,
            "queries": (func() map[string]string {
                list := map[string]string{}
                if query == nil {
                    return list
                }
                for key, val := range query {
                    list[key] = val[0]
                }
                return list
            })(),
        },
    }
    responseJson, _ := json.Marshal(responseData)
    
    return string(responseJson)
}
