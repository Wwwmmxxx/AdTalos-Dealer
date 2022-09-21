package request

import (
	"fmt"
	"github.com/Wwwmmxxx/adtalos-dealer/model"
	"github.com/Wwwmmxxx/adtalos-dealer/pkg/resty"
	"github.com/mxmCherry/openrtb/v15/openrtb2"
)

func OpenRTB25(token string) {

	// initial param
	url := fmt.Sprintf("https://api.mobrtb.com/ad/openrtb2/%v", token)

	body := model.NewOpenRTB(model.Pic)

	response := &openrtb2.BidResponse{}

	http, err := resty.NewResty().R().
		SetResult(&response).
		SetBody(body).
		SetHeader("x-openrtb-version", "2.5").
		Post(url)
	if err != nil {
		fmt.Println("OpenRTB2.5 Post failed, err:", err)
		return
	}

	fmt.Println(http)

}
