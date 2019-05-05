package frouter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Map map[string]interface{}

func WriteJSON(w http.ResponseWriter, data interface{}) {
	defer Rec()
	raw,err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w,string(raw))
}

func RespJSON(w http.ResponseWriter,CodeSuc int,CodeErr int) (func(data interface{}),func(err error)) {
	return func(data interface{}) {
		WriteJSON(w,Map{
			"code":		CodeSuc,
			"data":		data,
			"msg":		"success",
		})
	}, func(err error) {
		WriteJSON(w,Map{
			"code":		CodeErr,
			"data":		err.Error(),
			"msg":		"failed",
		})
	}
}