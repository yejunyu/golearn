package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"golearn/lesson11(defer)/filelistserver/filelisting"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		defer func() {
			if r := recover(); r != nil {
				log.Error("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		if err != nil {
			log.Warn("Error handling request: %s",
				err.Error())
			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Msg(), http.StatusBadRequest)
				return
			}
			var code int
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				panic("Unknown error")
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	// 给系统看的
	error
	// 给用户看的
	Msg() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.Handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
