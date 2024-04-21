package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/llms/googleai"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/generate", generateHandler).Methods("POST")

    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    htmlContent, err := ioutil.ReadFile("index.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "text/html")
    w.Write(htmlContent)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("prompt")
    prompt := "You are a very good and creative poet, Generate a poem on title"+title+" in case if you can't generate the poem then reply 'I can't dot it' nothing else."
    if title == "" {
        http.Error(w, "Title cannot be empty", http.StatusBadRequest)
        return
    }
    ctx := context.Background()
    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        http.Error(w, "API_KEY environment variable not found.", http.StatusInternalServerError)
        return
    }
    llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
    if err != nil {
        http.Error(w, fmt.Sprintf("Error while initializing language model: %v", err), http.StatusInternalServerError)
        return
    }

    answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error while generating poem : %v", err), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Generated Poem: %s", answer)
}
