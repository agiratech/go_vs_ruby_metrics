package main

import (
  "fmt"
  "net/http"
  "github.com/agiratech/go_vs_ruby_metrics/goCsv"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("./public")))
  http.HandleFunc("/go", goHandler)
  http.HandleFunc("/get-records", goApi)
  http.ListenAndServe(":3000", nil)
}

func goHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-type", "application/json")
  jsonMsg, err := goCsv.Import()
  if err != nil {
    http.Error(w, "Oops", http.StatusInternalServerError)
  }
  fmt.Fprintf(w, jsonMsg)
}

func goApi(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-type", "application/json")
  name :=  r.FormValue("name")
  err,jsonMsg := goCsv.SearchByName(name)
  if err != nil {
    http.Error(w, "Oops", http.StatusInternalServerError)
  }
  fmt.Fprintf(w, jsonMsg)
}