package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileserver := http.fileserver(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("starting server at port 8080/n")
	if err := http.ListenAndServer(":8080",nil);
	err !=nil{
		log.Fatal(err)
	}
}