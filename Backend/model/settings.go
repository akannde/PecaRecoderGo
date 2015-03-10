package model

import (
	"bytes"
	"os"

	"github.com/BurntSushi/toml"
)

// Settings は設定情報を扱うための構造体
type Settings struct {
	YellowPage   []YellowPageSetting
	Notification NotificationSetting
	Style        StyleSetting
	Player       []PlayerSetting
	Favorite     []FavoriteSetting
	Recoding     []RecodingSetting
}

// NewSettings は新しいSettingインスタンを生成する
func NewSettings() *Settings {
	yellowPageSetting := []YellowPageSetting{}

	// Peercast YellowPage SP
	// http://bayonet.ddo.jp/sp/index.txt
	yellowPageSetting = append(
		yellowPageSetting,
		YellowPageSetting{
			Name:                    "SP",
			IndexURL:                "http://bayonet.ddo.jp/sp/index.txt",
			CheckBitrateURL:         "http://bayonet.ddo.jp/sp/uptest/uptest_n.php",
			EnabledAutoCheckBitrate: false,
			Enabled:                 true,
		},
	)

	// Temporary Yellow Page
	// http://temp.orz.hm/yp/index.txt
	yellowPageSetting = append(
		yellowPageSetting,
		YellowPageSetting{
			Name:                    "TP",
			IndexURL:                "http://temp.orz.hm/yp/index.txt",
			CheckBitrateURL:         "http://temp.orz.hm/yp/uptest/uptest_n.php",
			EnabledAutoCheckBitrate: false,
			Enabled:                 true,
		},
	)

	// 秘密基地TV
	// http://games.himitsukichi.com/hktv/index.txt
	yellowPageSetting = append(
		yellowPageSetting,
		YellowPageSetting{
			Name:                    "HKTV",
			IndexURL:                "http://games.himitsukichi.com/hktv/index.txt",
			CheckBitrateURL:         "http://games.himitsukichi.com/hktv/uptest/uptest_n.php",
			EnabledAutoCheckBitrate: false,
			Enabled:                 true,
		},
	)

	// CaveTube
	// http://rss.cavelis.net/index.txt
	yellowPageSetting = append(
		yellowPageSetting,
		YellowPageSetting{
			Name:                    "CaveTube",
			IndexURL:                "http://rss.cavelis.net/index.txt",
			CheckBitrateURL:         "",
			EnabledAutoCheckBitrate: false,
			Enabled:                 true,
		},
	)

	// Turf-Page(芝)
	// http://peercast.takami98.net/turf-page/index.txt
	yellowPageSetting = append(
		yellowPageSetting,
		YellowPageSetting{
			Name:                    "芝",
			IndexURL:                "http://peercast.takami98.net/turf-page/index.txt",
			CheckBitrateURL:         "http://takami98.sakura.ne.jp/peca-navi/turf-page/index.php",
			EnabledAutoCheckBitrate: false,
			Enabled:                 true,
		},
	)

	var notificationSetting NotificationSetting
	notificationSetting.Position = 2
	notificationSetting.Enabled = false

	var styleSetting StyleSetting
	styleSetting.FontScale = 1.0

	return &Settings{
		YellowPage:   yellowPageSetting,
		Notification: notificationSetting,
		Style:        styleSetting,
	}
}

// LoadSettings は指定パスから設定情報を読み込む関数
func LoadSettings(path string) (*Settings, error) {
	var settings Settings
	_, err := toml.DecodeFile(path, &settings)
	return &settings, err
}

// Encode は SettingモデルをTOML構造へエンコードする関数
func (s *Settings) Encode() (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	err := encoder.Encode(&s)
	return &buffer, err
}

// Save は指定パスに設定ファイルを保存する関数
func (s *Settings) Save(path string) error {
	fp, err := os.Create(path)
	defer fp.Close()
	if err != nil {
		return err
	}
	buffer, err := s.Encode()
	if err != nil {
		return err
	}
	_, err = fp.Write(buffer.Bytes())
	return err
}

// YellowPageSetting はTOML用モデル
type YellowPageSetting struct {
	Name                    string `toml:"name"`                       // YP名
	IndexURL                string `toml:"index_url"`                  // index.txtのあるURL
	CheckBitrateURL         string `toml:"check_bitrate_url"`          // ビットレートチェックのあるURL
	EnabledAutoCheckBitrate bool   `toml:"enabled_auto_check_bitrate"` // 自動ビットレートチェックが有効化
	Enabled                 bool   `toml:"enabled"`                    // YPが有効か
}

// NotificationSetting はTOML用モデル
type NotificationSetting struct {
	Position int64 `toml:"position"` // 通知の出す場所 1:右上 2:右下 3:左上 4:左下
	Enabled  bool  `toml:"enabled"`  // 通知が有効か
}

// StyleSetting はTOML用モデル
type StyleSetting struct {
	FontScale float64 `toml:"font_scale"` // 特大(2.0) 大(1.5) 中(1.0) 小(0.5) の4つのサイズ
}

// PlayerSetting はTOML用モデル
type PlayerSetting struct {
	Name string `toml:"name"` // 表示名
	Path string `toml:"path"` // プレイヤーのフルパス
	Type string `toml:"type"` // WMV, FLV, OPV, RAWなどの動画形式
}

// FavoriteSetting はTOML用モデル
type FavoriteSetting struct {
	Name        string `toml:"name"`         // 表示名
	RegexString string `toml:"regex_string"` // 正規表現文字列
	Enabled     bool   `toml:"enabled"`      // 有効無効
}

// RecodingSetting はTOML用モデル
type RecodingSetting struct {
	OutputDirectory      string `toml:"output_dir"`             // 録画ファイル保存先ディレクトリ
	OutputFilenameFormat string `toml:"output_filename_format"` // ファイル名フォーマット %CHANNEL_NAME%: チャンネル名 %TIMESTAMP%: 録画開始時間
}
