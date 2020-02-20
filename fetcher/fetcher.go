package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"spider/config"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", config.GetSetting("Fetcher::UserAgent", 0).(string))

	resp, _ := http.DefaultClient.Do(request)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error:status code:%d", resp.StatusCode)
	}

	//如果页面传来的不是utf8，我们需要转为utf8格式
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(config.GetSetting("Fetcher::PeekSize", 1).(int))
	if err != nil {
		log.Printf("Ftcher error:%v", err)
		return nil
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
