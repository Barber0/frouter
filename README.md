# FRouter
## A web router library for Minimalism

## Quick Start

   	func main() {
    	r := frouter.NewFRouter()
    	r.GET("/one/", func(w http.ResponseWriter, r *http.Request) {
    		suc,_ := frouter.RespJSON(w,0,-1)
    		suc(frouter.D{
    			"code":		http.StatusOK,
    		})
    	})

    	api := r.Group("/api/")
    	{
    		api.POST("/two/", func(w http.ResponseWriter, r *http.Request) {
    			frouter.WriteJSON(w,frouter.D{
    				"code":		http.StatusOK,
    				"data":		nil,
    				"msg":		"suc",
    			})
    		})

    		loginOnly := api.Group("/",frouter.ExampleLoginCheck)
    		{
    			loginOnly.GET("/one/", func(w http.ResponseWriter, r *http.Request) {
    				suc,_ := frouter.RespJSON(w,0,-2)
    				suc("成功返回")
    			})
    		}
    	}

    	server := &http.Server{
    		Addr:		":5000",
    		Handler:	r,
    	}
    	server.ListenAndServe()
    }