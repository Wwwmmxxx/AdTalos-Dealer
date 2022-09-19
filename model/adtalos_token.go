package model

type PlacementToken string

const (
	BannerToken                  PlacementToken = "209A03F87BA3B4EB82BEC9E5F8B41383"
	BannerWithDeepLink           PlacementToken = "A48B8D8D3AF6F42F61BC204678D58CD4"
	Insert600_500                PlacementToken = "6ECB2643AFF34D8A1F73AB50F0BE3C91"
	Insert480_320                PlacementToken = "2EF810225D10260506CBB704C96C5325"
	Insert960_640                PlacementToken = "2EF810225D10260506CBB704C96C5325"
	VideoKaiPing720_1280         PlacementToken = "5C3DD65A809B08A2D6CF3DEFBC7E09C7"
	Video1                       PlacementToken = "425CDB8771F6743AD32B8C57DE2AA4C1"
	Pic1_Article1                PlacementToken = "98738D91D3BB241458D3FAE5A5BF7D34"
	Pic1_Article2                PlacementToken = "98738D91D3BB241458D3FAE5A5BF7134"
	Pic1_Article3                PlacementToken = "98738D91D3BB241458D3FAE5A5BF7234"
	Pic1_Icon1_Article1          PlacementToken = "98738D91D3BB241458D3FAE5A5BF7334"
	Pic1_Icon1_Article2          PlacementToken = "98738D91D3BB241458D3FAE5A5BF7434"
	Pic3_Article1                PlacementToken = "98738D91D3BB241458D3FAE5A5BF7534"
	Icon1_Article2               PlacementToken = "98738D91D3BB241458D3FAE5A5BF7634"
	Pic3_Icon1_Article2          PlacementToken = "98738D91D3BB241458D3FAE5A5BF7734"
	Pic1_Icon1_Article2_Button1  PlacementToken = "98738D91D3BB241458D3FAE5A5BF7834"
	Pic                          PlacementToken = "98738D91D3BB241458D3FAE5A5BF7934"
	Video1_Cover1_Icon1_Article2 PlacementToken = "98738D91D3BB241458D3FAE5A5BF7B34"
	Video1_Cover1                PlacementToken = "98738D91D3BB241458D3FAE5A5BF7C34"
)

var placementTokenArray = []PlacementToken{
	BannerToken, BannerWithDeepLink,
	Insert600_500, Insert480_320, Insert960_640,
	VideoKaiPing720_1280, Video1,
	Pic1_Article1, Pic1_Article2, Pic1_Article3,
	Pic1_Icon1_Article1, Pic1_Icon1_Article2,
	Pic3_Article1,
	Icon1_Article2,
	Pic3_Icon1_Article2,
	Pic1_Icon1_Article2_Button1,
	Pic,
	Video1_Cover1_Icon1_Article2,
	Video1_Cover1,
}
