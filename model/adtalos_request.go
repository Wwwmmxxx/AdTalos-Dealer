package model

import (
	"fmt"
	"github.com/google/uuid"
)

type AdTalosRequest struct {
	Id        string      `json:"id,omitempty"`
	Version   string      `json:"version" default:"2.0.0"`
	Ads       []AdRequest `json:"ads"`
	App       App         `json:"app"`
	Device    Device      `json:"device"`
	User      *User       `json:"user,omitempty"`
	NeedHttps bool        `json:"need_https,omitempty" default:"false"`
}

type AdRequest struct {
	AdUnitToken string  `json:"ad_unit_token"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	FloorPrice  float64 `json:"floor_price,omitempty"`
	SupportJs   bool    `json:"support_js,omitempty"`
}

type App struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Bundle       string `json:"bundle"`
	DeeplinkMode int8   `json:"deeplink_mode,omitempty"`
}

type Device struct {
	Ip             string  `json:"ip"`
	UserAgent      string  `json:"user_agent"`
	Make           string  `json:"make"`
	Brand          string  `json:"brand"`
	Model          string  `json:"model"`
	Os             string  `json:"os"`
	OsVersion      string  `json:"os_version"`
	ConnectionType string  `json:"connection_type"`
	Orientation    string  `json:"orientation"`
	Plmn           string  `json:"plmn,omitempty"`
	Mac            string  `json:"mac,omitempty"`
	MacMd5         string  `json:"mac_md5,omitempty"`
	Language       string  `json:"language,omitempty"`
	ScreenWidth    int     `json:"screen_width,omitempty"`
	ScreenHeight   int     `json:"screen_height,omitempty"`
	ScreenDpi      int     `json:"screen_dpi,omitempty"`
	ScreenPxratio  float64 `json:"screen_pxratio,omitempty"`
	GeoLongitude   float64 `json:"geo_longitude,omitempty"`
	GeoLatitude    float64 `json:"geo_latitude,omitempty"`
	DeviceType     int     `json:"device_type,omitempty"`
	Ssid           string  `json:"ssid,omitempty"`
	WifiMac        string  `json:"wifi_mac,omitempty"`

	// Below field only used in Android
	Imei                 string `json:"imei,omitempty"`
	ImeiMd5              string `json:"imei_md5,omitempty"`
	AndroidId            string `json:"android_id,omitempty"`
	AndroidIdMd5         string `json:"android_id_md5,omitempty"`
	Imsi                 string `json:"imsi,omitempty"`
	AndroidAdvertisingId string `json:"android_advertising_id,omitempty"`
	Oaid                 string `json:"oaid,omitempty"`
	OaidMd5              string `json:"oaid_md5,omitempty"`
	RomVersion           string `json:"rom_version,omitempty"`
	SysCompilingTime     int64  `json:"sys_compiling_time,omitempty"`

	// Below field only used in IOS
	Idfa        string `json:"idfa,omitempty"`
	IdfaMd5     string `json:"idfa_md5,omitempty"`
	Idfv        string `json:"idfv,omitempty"`
	Caid        string `json:"caid,omitempty"`
	CaidVersion string `json:"caid_version,omitempty"`
	Openudid    string `json:"openudid,omitempty"`
	BootMark    string `json:"boot_mark,omitempty"`
	UpdateMark  string `json:"update_mark,omitempty"`
}

type User struct {
	Age      int      `json:"age,omitempty"`
	Gender   string   `json:"gender,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

func NewAdTalosRequest(tokens ...PlacementToken) (*AdTalosRequest, error) {

	if len(tokens) == 0 {
		return nil, fmt.Errorf("no placementToken in param")
	}

	ads := make([]AdRequest, 0, len(tokens))

	for _, token := range tokens {

		ads = append(ads, AdRequest{
			AdUnitToken: string(token),
			Width:       480,
			Height:      320,
		})
	}

	return &AdTalosRequest{
		Id:      uuid.NewString(),
		Version: "2.0.0",
		Ads:     ads,
		App: App{
			Name:    "Wwwmmxxx",
			Version: "0.0.1",
			Bundle:  "com.sh.sdju",
		},
		Device: Device{
			Ip:             "58.247.129.11",
			UserAgent:      "go-resty/2.7.0 (https://github.com/go-resty/resty)",
			Make:           "Samsung",
			Brand:          "MI4",
			Model:          "iphone13",
			Os:             "ios",
			OsVersion:      "15.1",
			ConnectionType: "4g",
			Orientation:    "landscape",
		},
	}, nil
}
