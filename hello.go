package main
import {
    "fmt",
    "net/http",
    "os"
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := os.Getenv("RESPONSE")
    if len(response) == 0 {
    response = "Hello Openshift for Developers! - Satya"
}

    fmt.Fprintln((w, response)
    fmt.Println("Serving an impatient beginner's request.")

}
