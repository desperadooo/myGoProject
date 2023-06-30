package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("-->WithServerHeader()")
		w.Header().Set("Server", "HelloServer v0.0.1")
		h(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d\n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s\n", a, b)
	return a + b
}

func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func main() {
	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello", "world!")
	//http.HandleFunc("v1/hello", WithServerHeader(hello))
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
