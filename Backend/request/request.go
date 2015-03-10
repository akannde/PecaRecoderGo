package request

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/golang/glog"
)

// Request はHTTPリクエスト構造体
type Request struct {
	URL       string
	CookieJar *cookiejar.Jar
}

// NewRequest はRequestをインスタンス化する
func NewRequest(url string) *Request {
	cj, _ := cookiejar.New(nil)
	return &Request{
		URL:       url,
		CookieJar: cj,
	}
}

// getResponse は
func (r *Request) getResponse(url string, maxRetry int, retryDelaySec int) (*http.Response, error) {
	client := &http.Client{
		Jar: r.CookieJar,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:38.0) Gecko/20100101 Firefox/38.0")
	req.Header.Set("Referer", url)

	res, err := client.Do(req)
	for retry := 1; (res == nil || res.ContentLength == 0) && retry < maxRetry; {
		if err != nil {
			glog.Errorf("Attempt #%d of %d of doing %s(%s) failed with: %s.\nRetrying in %d seconds \n", retry, maxRetry, req.Method, url, err, retryDelaySec)
		} else {
			glog.Errorf("Attempt #%d of %d of doing %s(%s).\nRetrying in %d seconds \n", retry, maxRetry, req.Method, url, retryDelaySec)
		}
		if retry > 1 {
			time.Sleep(time.Duration(retryDelaySec) * time.Second)
		}
		res, err = client.Do(req)
		retry++
	}

	return res, err
}
