package frouter

import (
	"errors"
	"fmt"
	"log"
)

var	(
	MethodNotMatch 			= 	errMsg("controller","method not match")
	ExpectRouterGroupArgs	=	errMsg("middleware","args expected")
	NotLogin				=	errMsg("middleware","invalid token")
)

func errMsg(tag string,content string) error {
	return errors.New(fmt.Sprintf("[%s] %s",tag,content))
}

func Rec() {
	if err := recover(); err != nil {
		log.Printf("[frouter] %v",err)
	}
}