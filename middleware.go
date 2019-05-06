package frouter

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Middleware func(next http.Handler) http.Handler

func ExampleLoginCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		defer Rec()
		headers := r.Header.Get("Authorization")
		if strings.TrimSpace(headers) == "" {
			fmt.Fprint(w,"not login")
			panic(NotLogin)
		}
		next.ServeHTTP(w,r)
	})
}

func ExampleBeta(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		defer Rec()
		val := r.FormValue("key")
		if val != "abc" {
			fmt.Fprint(w,"no value")
			panic(errors.New("no value"))
		}
		next.ServeHTTP(w,r)
	})
}