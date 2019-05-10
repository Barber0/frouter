package frouter

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestPOSTJSON(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	r := NewFRouter()
	r.POST("/one/", func(w http.ResponseWriter, r *http.Request) {
		defer Rec()
		res := D{}
		if err:=ReadJSON(r,&res);err!=nil {
			panic(err)
		}
		fmt.Println(res["code"])
		WriteJSON(w,res)
	})

	r.GET("/temp/", func(w http.ResponseWriter, r *http.Request) {
		WriteHTML(w,"./test/index.html")
	})

	r.POST("/fl/", func(w http.ResponseWriter, r *http.Request) {
		defer Rec()
		suc,fail := RespJSON(w,0,-1)
		Check(SaveFile(r,"myfile","./test/resource/one.jpg"),fail)
		suc("上传成功")
	})

	server := &http.Server{
		Addr:		":5000",
		Handler:	r,
	}
	server.ListenAndServe()
}