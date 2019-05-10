# FRouter
A web router library for Minimalism

## Quick Start

    defer Rec()
	r := NewFRouter()
	r.POST("/one/", func(w http.ResponseWriter, r *http.Request) {
		defer Rec()
		res := D{}
		if err:=ReadJSON(r,&res);err!=nil {
			panic(err)
		}
		fmt.Println(res["code"])
		JSON(w,res)
	})

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		HTML(w,"./test/index.html")
	})

	r.POST("/upload/", func(w http.ResponseWriter, r *http.Request) {
		defer Rec()
		Check(SaveFile(r,"myfile","./test/resource/one.jpg"), func(e error) {
			JSON(w,D{
				"code":		-1,
				"msg":		"上传失败",
			})
		})
		JSON(w,D{
			"code":		0,
			"msg":		"上传成功",
		})
	})

	server := &http.Server{
		Addr:		":5000",
		Handler:	r,
	}
	server.ListenAndServe()