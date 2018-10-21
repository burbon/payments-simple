package ps

import (
	"io/ioutil"
	"net/http"
)

func FetchPaymentsFromSource() (body []byte, err error) {
	resp, err := http.Get(config.PaymentsSourceURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
