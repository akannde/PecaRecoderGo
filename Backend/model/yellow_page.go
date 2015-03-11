package model

// YellowPageInfo YellowPageの情報
type YellowPageInfo struct {
	ID              int64  `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	IndexTxtURL     string `gorm:"column:index_txt_url"`
	CheckBitrateURL string `gorm:"column:check_bitrate_url"`
}

// NewYellowPageInfo YellowPageInfoを生成するための関数
func NewYellowPageInfo(name string, indexTxtURL string, checkBitrateURL string) *YellowPageInfo {
	return &YellowPageInfo{
		Name:            name,
		IndexTxtURL:     indexTxtURL,
		CheckBitrateURL: checkBitrateURL,
	}
}
