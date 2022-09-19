package model

type AdTalosResponse struct {
	Id  string       `json:"id,omitempty"`
	Ads []AdResponse `json:"ads,omitempty"`
}

type AdResponse struct {
	Width      int     `json:"width"`
	Height     int     `json:"height"`
	AdId       string  `json:"ad_id"`
	CreativeId string  `json:"creative_id"`
	Price      float64 `json:"price,omitempty"`

	// below used for display
	Title          string  `json:"title,omitempty"`
	Subtitle       string  `json:"subtitle,omitempty"`
	Description    string  `json:"description,omitempty"`
	AdvertiserName string  `json:"advertiser_name,omitempty"`
	Ratings        string  `json:"ratings,omitempty"`
	ButtonText     string  `json:"button_text,omitempty"`
	Likes          string  `json:"likes,omitempty"`
	Downloads      string  `json:"downloads,omitempty"`
	Images         []Image `json:"images,omitempty"`
	Icon           Image   `json:"icon,omitempty"`
	Logo           Image   `json:"logo,omitempty"`
	Video          Video   `json:"video,omitempty"`
	VideoCover     Image   `json:"video_cover,omitempty"`
	HtmlSnippet    string  `json:"html_snippet,omitempty"`
	HtmlUrl        string  `json:"html_url,omitempty"`

	// below used for click
	Action             int    `json:"action"`
	TargetUrl          string `json:"target_url"`
	DownloadAppBundle  string `json:"download_app_bundle,omitempty"`
	DownloadAppName    string `json:"download_app_name,omitempty"`
	DownloadAppVersion string `json:"download_app_version,omitempty"`
	DownloadAppSize    int    `json:"download_app_size,omitempty"`
	DeeplinkUrl        string `json:"deeplink_url,omitempty"`

	// below used for report
	WinNoticeTracker                 string   `json:"win_notice_tracker"`
	ImpressionTrackers               []string `json:"impression_trackers"`
	ClickTrackers                    []string `json:"click_trackers"`
	DownloadBeginTrackers            []string `json:"download_begin_trackers,omitempty"`
	DownloadEndTrackers              []string `json:"download_end_trackers,omitempty"`
	InstallBeginTrackers             []string `json:"install_begin_trackers,omitempty"`
	InstallEndedTrackers             []string `json:"install_ended_trackers,omitempty"`
	VideoPlayBeginTrackers           []string `json:"video_play_begin_trackers,omitempty"`
	VideoPlayBreakTrackers           []string `json:"video_play_break_trackers,omitempty"`
	VideoPlayEndedTrackers           []string `json:"video_play_ended_trackers,omitempty"`
	DeeplinkAppNotInstalledTrackers  []string `json:"deeplink_app_not_installed_trackers,omitempty"`
	DeeplinkAppInstalledTrackers     []string `json:"deeplink_app_installed_trackers,omitempty"`
	DeeplinkAppInvokeFailedTrackers  []string `json:"deeplink_app_invoke_failed_trackers,omitempty"`
	DeeplinkAppInvokeSuccessTrackers []string `json:"deeplink_app_invoke_success_trackers,omitempty"`
}

type Image struct {
	Url    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type Video struct {
	Url      string `json:"url"`
	Duration int    `json:"duration,omitempty"`
}
