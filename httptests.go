package httptests

import (
  "fmt"
  "errors"
  "strings"
  "io/ioutil"
  "net/http"
  )

func getHttpUrl(url string, expectedBody string) (int, error) {
  resp, err := http.Get(url)
  if err != nil {
    return 0, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return 0, err
  }

  if ! strings.Contains(string(body), expectedBody) {
    return 0, errors.New("Invalid response")
  }

  return 0, nil
}
