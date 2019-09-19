package loader

import (
  "os"
)

func vidExists(id string) (bool) {
  fileName := id + ".mp3"

  _, err := os.Stat(fileName);
  return !os.IsNotExist(err)
}
