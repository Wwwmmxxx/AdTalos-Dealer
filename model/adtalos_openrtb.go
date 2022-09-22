package model

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mxmCherry/openrtb/v15/native1"
	"github.com/mxmCherry/openrtb/v15/native1/request"
	"github.com/mxmCherry/openrtb/v15/openrtb2"
)

func NewOpenRTB(pt PlacementToken) *openrtb2.BidRequest {

	type NativeRequest struct {
		Native request.Request `json:"native"`
	}

	native := NativeRequest{
		request.Request{
			Ver:            "1.2",
			Layout:         0,
			AdUnit:         0,
			Context:        0,
			ContextSubType: 0,
			PlcmtType:      0,
			PlcmtCnt:       0,
			Seq:            0,
			Assets: []request.Asset{
				{
					ID:       1,
					Required: 1,
					Img: &request.Image{
						Type:  native1.ImageAssetTypeIcon,
						W:     0,
						WMin:  0,
						H:     0,
						HMin:  0,
						MIMEs: nil,
						Ext:   nil,
					},
				},
			},
			AURLSupport:   0,
			DURLSupport:   0,
			EventTrackers: nil,
			Privacy:       0,
			Ext:           nil,
		},
	}

	nativeInJson, err := json.Marshal(native)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return nil
	}

	return &openrtb2.BidRequest{
		ID: uuid.NewString(),
		Imp: []openrtb2.Imp{
			{
				ID: uuid.NewString(),
				Native: &openrtb2.Native{
					Request: string(nativeInJson),
					Ver:     "1.2",
				},
				Instl:       1,
				TagID:       string(pt),
				BidFloor:    0,
				BidFloorCur: "CNY",
			},
		},
		App: &openrtb2.App{
			Name:   "Wwwmmxxx",
			Bundle: "www.digverity.com",
			Ver:    "1.0",
		},
		Device: &openrtb2.Device{
			IP: "58.247.129.11",
			UA: "go-resty/2.7.0 (https://github.com/go-resty/resty)",
			Geo: &openrtb2.Geo{
				Lat: 220,
				Lon: 300,
			},
			Make:           "Samsung",
			Model:          "iphone13",
			OS:             "ios",
			OSV:            "15.1",
			ConnectionType: openrtb2.ConnectionTypeCellularNetwork4G.Ptr(),
		},
	}

}
