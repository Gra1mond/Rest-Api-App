package midleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method,r.URL.Path)

		next.ServeHTTP(w,r)
	})
}

func Auth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token:=r.Header.Get("Authorization")

		if token==""{
			http.Error(w,"unauthorized",http.StatusUnauthorized)
			return 
		}
		next.ServeHTTP(w,r)
	})
}