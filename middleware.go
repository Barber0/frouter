package frouter

import (
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
		}else {
			next.ServeHTTP(w,r)
		}
	})
}