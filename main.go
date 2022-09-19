package main

import (
	"fmt"
	"github.com/Wwwmmxxx/adtalos-dealer/model"
	"github.com/Wwwmmxxx/adtalos-dealer/pkg/resty"
)

func main() {

	// initial param
	token := "1058F1630E33C900BE9753D33DDC59CE"
	url := fmt.Sprintf("https://api.mobrtb.com/ad/xy/%v", token)

	// send post request
	request, err := model.NewAdTalosRequest(model.Video1_Cover1)
	if err != nil {
		fmt.Println("NewAdTalosRequest failed, err:", err)
	}

	response := &model.AdTalosResponse{}

	http, err := resty.NewResty().R().
		SetResult(&response).
		SetBody(request).
		Post(url)
	if err != nil {
		fmt.Println("Post failed, err:", err)
		return
	}

	fmt.Println(http)
}
