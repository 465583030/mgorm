package main

import (
	"encoding/json"
	"fmt"
	"github.com/elsonwu/restapi/model"
	"net/http"
	// "reflect"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		data, _ := model.UserModel().Find()
		data.LastName = "TTTT"
		data.FirstName = "EEEE"
		err := data.Save()
		if err != nil {
			fmt.Println(err)
			res.Write([]byte("ERROR"))
		}
		output, err := json.Marshal(*data)
		if err != nil {
			fmt.Println("marshal error")
			res.Write([]byte("json marshal error"))
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(output)
	})

	http.ListenAndServe(":8888", nil)
}
