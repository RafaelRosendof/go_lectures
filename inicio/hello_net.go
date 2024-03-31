package main

import ( 
  "fmt"
  "net/http"
)

func hello(res http.ResponseWriter, req *http.Request){
  fmt.Fprint(res, "Eai meu caro, como vc está???")
}

func main(){
  http.HandleFunc("/", hello)
  http.ListenAndServe("localhost:4321",nil)
}
