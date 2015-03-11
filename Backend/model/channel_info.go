package model

import "fmt"

// ChannelInfo はindex.txtをパースした時の構造体
type ChannelInfo struct {
	Name             string `gorm:"column:name"`
	ChanID           string `gorm:"column:chan_id;primary_key:yes"`
	BroadcastingAddr string `gorm:"column:broadcasting_addr"`
	ContactURL       string `gorm:"column:contact_url"`
	Genre            string `gorm:"column:genre"`
	Description      string `gorm:"column:description"`
	Views            int64  `gorm:"column:views"`
	Relays           int64  `gorm:"column:relays"`
	Bitrate          int64  `gorm:"column:bitrate"`
	StreamingType    string `gorm:"column:streaming_type"`
	BroadcastingTime string `gorm:"column:broadcasting_time"`
	Status           string `gorm:"column:status"`
	Comment          string `gorm:"column:comment"`
	YellowPage       YellowPageInfo
}

// NewChannelInfo はChannelInfo生成するための関数
func NewChannelInfo(
	name string,
	chanID string,
	broadcastingAddr string,
	contactURL string,
	genre string,
	desc string,
	views int64,
	relayes int64,
	bitrate int64,
	streamingType string,
	broadcastingTime string,
	status string,
	comment string,
	yp YellowPageInfo,
) *ChannelInfo {
	return &ChannelInfo{
		Name:             name,
		ChanID:           chanID,
		BroadcastingAddr: broadcastingAddr,
		ContactURL:       contactURL,
		Genre:            genre,
		Description:      desc,
		Views:            views,
		Relays:           relayes,
		Bitrate:          bitrate,
		StreamingType:    streamingType,
		BroadcastingTime: broadcastingTime,
		Status:           status,
		Comment:          comment,
		YellowPage:       yp,
	}
}

// BuildStreamingURL はストリーミングURLを生成する関数
func (c *ChannelInfo) BuildStreamingURL(srvAddr string) string {
	return fmt.Sprintf("http://%s/pls/%s?tip=%s", srvAddr, c.ChanID, c.BroadcastingAddr)
}

// ChannelInfoByName はNameキーでのソート用構造体
type ChannelInfoByName []ChannelInfo

// Len は[]ChannelInfoの長さを求める関数
func (s ChannelInfoByName) Len() int {
	return len(s)
}

// Less はi番目のChannelInfoのNameとj番目のChannelInfoのNameの比較する関数
func (s ChannelInfoByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// Swap はi番目のChannelInfoとj番目のChannelInfoの順番を入れ替える関数
func (s ChannelInfoByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ChannelInfoByViews はViewsキーでのソート用構造体
type ChannelInfoByViews []ChannelInfo

// Len は[]ChannelInfoの長さを求める関数
func (s ChannelInfoByViews) Len() int {
	return len(s)
}

// Less はi番目のChannelInfoのViewsとj番目のChannelInfoのViewsの比較する関数
func (s ChannelInfoByViews) Less(i, j int) bool {
	return s[i].Views < s[j].Views
}

// Swap はi番目のChannelInfoとj番目のChannelInfoの順番を入れ替える関数
func (s ChannelInfoByViews) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
