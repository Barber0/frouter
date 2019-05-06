package frouter

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestAlpha(t *testing.T) {
	r := NewFRouter()
	r.GET("/test/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"ddddddddd")
	})
	r.POST("/post/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"postfffffff")
	})

	api := r.Group("/alpha/",ExampleLoginCheck)
	{
		test := api.Group("/one/",ExampleBeta)
		{
			test.POST("/test/", func(w http.ResponseWriter, r *http.Request) {
				suc,_ := RespJSON(w,0,-2)
				suc(D{
					"id":	20,
					"uname":"alpha",
					"pass":	"one",
				})
			})
		}

		api.GET("/test/",http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
			suc,fail := RespJSON(w,0,-1)
			if r.FormValue("key") == "abc" {
				suc("成功")
			}else {
				fail(errors.New("失败"))
			}
		}))
	}

	server := &http.Server{
		Addr:		":5000",
		Handler:	r,
	}
	server.ListenAndServe()
}