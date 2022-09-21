package request

import (
	"fmt"
	"github.com/Wwwmmxxx/adtalos-dealer/model"
	"github.com/Wwwmmxxx/adtalos-dealer/pkg/resty"
)

func SSP(token string) {

	// initial param
	url := fmt.Sprintf("https://api.mobrtb.com/ad/xy/%v", token)

	// send post body
	body, err := model.NewAdTalosRequest(model.Video1)
	if err != nil {
		fmt.Println("NewAdTalosRequest failed, err:", err)
		return
	}

	response := &model.AdTalosResponse{}

	http, err := resty.NewResty().R().
		SetResult(&response).
		SetBody(body).
		Post(url)
	if err != nil {
		fmt.Println("Post failed, err:", err)
		return
	}

	fmt.Println(http)
}
