package frouter

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type D map[string]interface{}

func ReadJSON(r *http.Request, out interface{}) error {
	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
		out = nil
		return err
	}
	return json.Unmarshal(body,out)
}

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

func SaveFile(r *http.Request, key string, path string) error {
	file,_,err := r.FormFile(key)
	if err != nil {
		return err
	}
	defer file.Close()
	f,err := os.OpenFile(path,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_,err = io.Copy(f,file)
	return err
}

func WriteHTML(w http.ResponseWriter,path string,data...interface{})  {
	defer Rec()
	tpl,err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}
	if err := tpl.Execute(w, data); err != nil {
		panic(err)
	}
}