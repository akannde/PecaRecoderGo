package parser

import (
	"bufio"
	"bytes"
	"html"
	"log"
	"strconv"
	"strings"

	"github.com/PyYoshi/PecaRecoderGo/Backend/model"
)

// ParseIndexTxt はindex.txtをパースして[]ChannelInfoで返す関数
func ParseIndexTxt(data []byte, yp string) (*[]model.ChannelInfo, error) {
	var err error
	channels := []model.ChannelInfo{}
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "<>")
		views, err := strconv.ParseInt(data[6], 10, 64)
		if err != nil {
			views = 0
			err = nil
		}
		relayes, err := strconv.ParseInt(data[7], 10, 64)
		if err != nil {
			relayes = 0
			err = nil
		}
		bitrate, err := strconv.ParseInt(data[8], 10, 64)
		if err != nil {
			bitrate = 0
			err = nil
		}
		if len(data) == 19 {
			channel := model.NewChannelInfo(
				html.UnescapeString(data[0]),
				data[1],
				data[2],
				data[3],
				data[4],
				html.UnescapeString(data[5]),
				views,
				relayes,
				bitrate,
				data[9],
				data[15],
				data[16],
				data[17],
				yp,
			)
			channels = append(channels, *channel)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &channels, err
}
