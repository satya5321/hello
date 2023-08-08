package main
import {
    "fmt",
    "net/http",
    "os"
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := os.Getenv("RESPONSE"_
    if len(response) == 0 {
    response = "Hello Openshift for Developers! - Satya"
}

    fmt.fprintln((w, response)
    fmt.Println("Serving an impatient beginner's request.")

}
