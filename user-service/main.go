package main

import (
	"fmt"
	"user-service/domain/model"
)

func main (){
  var user = model.User{Id: "asd"}

  fmt.Println(user.Id)
}
