package httptests

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Timing struct {
	latency      time.Duration
	responseTime time.Duration
}

func GetHttpUrl(url string) (string, Timing, error) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return "", Timing{}, err
	}
	latency := time.Now().Sub(start)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	responseTime := time.Now().Sub(start)
	if err != nil {
		return "", Timing{}, err
	}

	return string(body), Timing{latency, responseTime}, nil
}

func ExpectHttpUrl(url string, expectedBody string) error {
	body, _, err := GetHttpUrl(url)
	if err != nil {
		return err
	}

	if !strings.Contains(body, expectedBody) {
		return errors.New("Invalid response")
	}

	return nil
}

func WaitForUrl(url string, timeout time.Duration) error {
	absTimeout := time.Now().Add(timeout)

	_, _, err := GetHttpUrl(url)
	for err != nil && time.Now().Before(absTimeout) {
		time.Sleep(50 * time.Millisecond)
		_, _, err = GetHttpUrl(url)
	}

	return err
}
