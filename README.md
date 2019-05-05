# FRouter
## A web router library for Minimalism

## Quick Start

   	r := NewFRouter()
   	r.GET("/test/", func(w http.ResponseWriter, r *http.Request) {
   		fmt.Fprint(w,"ddddddddd")
   	})
   	r.POST("/post/", func(w http.ResponseWriter, r *http.Request) {
   		fmt.Fprint(w,"postfffffff")
   	})

   	api := r.Group("/alpha/",ExampleLoginCheck)
   	{
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