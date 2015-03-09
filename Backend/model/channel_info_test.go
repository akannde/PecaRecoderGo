package model

import "testing"

func TestBuildStreamingURL(t *testing.T) {
	validURL := "http://127.0.0.1:7144/pls/1068A408B588A4BC6BA150B950EF6E24?tip=123.456.78.910:7144"
	srvAddr := "127.0.0.1:7144"
	chanID := "1068A408B588A4BC6BA150B950EF6E24"
	broadcastingAddr := "123.456.78.910:7144"
	channel := NewChannelInfo(
		"hoge",
		chanID,
		broadcastingAddr,
		"http://www.google.com",
		"dev",
		"desc",
		12,
		12,
		1024,
		"WMV",
		"11:11",
		"click",
		"comment",
		"SP",
	)
	if channel.BuildStreamingURL(srvAddr) != validURL {
		t.Error("配信中のURLが正しく生成されていません.")
	}
}
