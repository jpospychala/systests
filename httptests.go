package httptests

import (
	"errors"
	"github.com/jpospychala/systests/benchmarks"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Timing struct {
	latency      time.Duration
	responseTime time.Duration
}

func GetHttpUrl(url string, pm *benchmarks.ProgressMonitor) (string, error) {
	pm.Start()
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	pm.Step()

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	pm.Step()
	return string(body), nil
}

func ExpectHttpUrl(url string, expectedBody string, pm *benchmarks.ProgressMonitor) error {
	body, err := GetHttpUrl(url, pm)
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

	pm := &benchmarks.ProgressMonitor{}
	_, err := GetHttpUrl(url, pm)
	for err != nil && time.Now().Before(absTimeout) {
		time.Sleep(50 * time.Millisecond)
		_, err = GetHttpUrl(url, pm)
	}

	return err
}
