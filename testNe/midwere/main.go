package main

import (
	"log"
	"net/http"
	"time"
)

var limitCh = make(chan struct{}, 100)

type middleware func(handler http.Handler) http.Handler
type Router struct {
	middleWareChain []middleware
	mux             map[string]http.Handler
}

func timeMiddlerWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		next.ServeHTTP(writer, request)
		timeElapsed := time.Since(now)
		log.Printf("request %s use %d ms \n", request.URL.Path, timeElapsed.Milliseconds())
	})
}

func limitMiddlerWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		limitCh <- struct{}{}
		next.ServeHTTP(writer, request)
		<-limitCh
	})
}
func NewRouter() *Router {
	return &Router{
		middleWareChain: make([]middleware, 0, 10),
		mux:             make(map[string]http.Handler, 10),
	}
}

// 添加中间件
func (s *Router) Use(m middleware) {
	s.middleWareChain = append(s.middleWareChain, m)
}

// 自定义路由
func (s *Router) Add(path string, handler http.Handler) {
	var mergeHandler = handler
	for i := 0; i < len(s.middleWareChain); i++ {
		mergeHandler = s.middleWareChain[i](mergeHandler)
	}
	s.mux[path] = mergeHandler
}
func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Golang"))
}

func main() {
	router := NewRouter()
	router.Use(timeMiddlerWare)
	router.Use(limitMiddlerWare)
	router.Add("/", http.HandlerFunc(get))
	for path, handle := range router.mux {
		http.Handle(path, handle)
	}
	//http.HandleFunc("/", get)
	http.ListenAndServe(":8080", nil)
}
