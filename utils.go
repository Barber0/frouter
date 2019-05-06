package frouter

import (
	"encoding/json"
	"net/http"
)

type D map[string]interface{}

func WriteJSON(w http.ResponseWriter, data interface{}) {
	defer Rec()
	raw,err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type","text/json")
	w.Write(raw)
}

func RespJSON(w http.ResponseWriter,CodeSuc int,CodeErr int) (func(data interface{}),func(err error)) {
	return func(data interface{}) {
		WriteJSON(w,D{
			"code":		CodeSuc,
			"data":		data,
			"msg":		"success",
		})
	}, func(err error) {
		WriteJSON(w,D{
			"code":		CodeErr,
			"data":		err.Error(),
			"msg":		"failed",
		})
	}
}