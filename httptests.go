package httptests

import (
  "errors"
  "strings"
  "io/ioutil"
  "net/http"
  "time"
  )

func GetHttpUrl(url string) (string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}

func ExpectHttpUrl(url string, expectedBody string) (error) {
  body, err := GetHttpUrl(url)
  if err != nil {
    return err
  }

  if ! strings.Contains(body, expectedBody) {
    return errors.New("Invalid response")
  }

  return nil
}

func slice(args ...interface{}) []interface{} {
    return args
}

func WaitForUrl(url string, timeout time.Duration) (error) {
  absTimeout := time.Now().Add(timeout)

  _, err := GetHttpUrl(url)
  for err != nil && time.Now().Before(absTimeout) {
    time.Sleep(50 * time.Millisecond)
    _, err = GetHttpUrl(url)
  }

  return err
}
