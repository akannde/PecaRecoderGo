PeercastRecorderのFLV版っぽいなにかのメモ書き
======================================

Windows, Linux, OSXで動作することを目指す

50時間くらいでプロトタイプをつくる

# フロントエンド

## 環境

- Qt5.4 WebEngine(Chromium)を利用
- C++
- QML
- 通知 https://github.com/Snorenotify/Snorenotify

## 仕様

- 日本語入力可能(fcitxを利用する場合は別途プラグインが必要)
- バックエンドから提供されるHTMLを描画してUIを構成
- バックエンドからもらった配信者URLを規定のブラウザで開く
- バックエンドからもらった動画URI(https, mms, rtmp)を設定されているプレイヤーで開く
- 最小化するとシステムトレイへ格納
- 実行時にコマンド引数を渡すと開発者モードで起動してDeveloperToolsを開けるように ※5.4.1ではQWebEngineSettings::DeveloperExtrasEnabledは未実装
- 各プラットフォームにあった通知 Qt
- 新バージョンがでたら通知する機能

# バックエンド

## 環境

- Go 1.4
- go-martini/martini: web framework
- robfig/cron: cronっぽく処理を実行できる
- mattn/sqlite
- jinzhu/gorm: ORM
- glog: logging tool

## 仕様

- cronのように一定時間で処理を繰り返す仕組みを基板にする
- index.txt周辺の処理 取得 パーズなど
- 録画(WMV, FLV) コンテナが正常かどうか判断するためにヘッダパーザが必要かも
- フロントエンドにHTMLを提供(コマンドライン引数でhtmlディレクトリパスを指定. 存在しない場合はHTTPサーバは稼働しない)
- Linuxで使用する場合も考えてコマンドライン引数でデーモン化(単体動作)できるように. pidファイルの実装
- bind addressは127.0.0.1またはlocalhostで固定
- なるべくサードパーティに頼らない(大嘘)
- 新しいバージョンがあるか確認 github pagesでバージョン通知ファイルを管理
- YPの帯域チェックの自動化(グレーな機能かも…)

# UI

- Bower
- Angular 1.3.x
- angular-material
- jQueryは使わない

## JS

- 移植性を考えてサーバとのアクセスをする部分はモジュール化 => 極力サードパーティ製ライブラリを使用しない

## 開発環境構築

``npm install -g`` で ``bower`` ``grunt-cli`` を入れておく

```

# 設定できる項目

- YellowPageの管理
- お気に入りチャンネルなどの管理(regex, 全角半角無視みたいなお手軽機能など)
- コンテンツタイプごとにプレイヤーを設定
- フォントサイズ
- 通知 オンオフ

# ライセンス

Qtを利用しているものはGPL v3
