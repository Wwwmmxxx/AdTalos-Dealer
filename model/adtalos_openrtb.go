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
			Ver:            "",
			Layout:         0,
			AdUnit:         0,
			Context:        0,
			ContextSubType: 0,
			PlcmtType:      0,
			PlcmtCnt:       0,
			Seq:            0,
			Assets: []request.Asset{
				{
					ID:       0,
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

	_, err := json.Marshal(native)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return nil
	}

	return &openrtb2.BidRequest{
		ID: uuid.NewString(),
		Imp: []openrtb2.Imp{{
			ID:     uuid.NewString(),
			Metric: nil,
			Banner: &openrtb2.Banner{
				Format:   nil,
				W:        nil,
				H:        nil,
				WMax:     0,
				HMax:     0,
				WMin:     0,
				HMin:     0,
				BType:    nil,
				BAttr:    nil,
				Pos:      nil,
				MIMEs:    nil,
				TopFrame: 0,
				ExpDir:   nil,
				API:      nil,
				ID:       "",
				VCm:      0,
				Ext:      nil,
			},
			//Native: &openrtb2.Native{
			//	Request: string(nativeInJson),
			//	Ver:     "",
			//	API:     nil,
			//	BAttr:   nil,
			//	Ext:     nil,
			//},
			PMP:               nil,
			DisplayManager:    "",
			DisplayManagerVer: "",
			Instl:             0,
			TagID:             string(pt),
			BidFloor:          0,
			BidFloorCur:       "",
			ClickBrowser:      0,
			Secure:            nil,
			IframeBuster:      nil,
			Exp:               0,
			Ext:               nil,
		}},
		Site:    nil,
		App:     nil,
		Device:  nil,
		User:    nil,
		Test:    0,
		AT:      0,
		TMax:    0,
		WSeat:   nil,
		BSeat:   nil,
		AllImps: 0,
		Cur:     nil,
		WLang:   nil,
		BCat:    nil,
		BAdv:    nil,
		BApp:    nil,
		Source:  nil,
		Regs:    nil,
		Ext:     nil,
	}

}
