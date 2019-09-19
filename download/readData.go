package download

import (
  "errors"
  "io/ioutil"
)

func readData(id string) ([]byte, error) {
  fileName := id + ".mp3"

  data, err := ioutil.ReadFile(fileName)
  if err != nil {
    return nil, errors.New("Could not load mp3 data")
  }

  return data, nil
}
