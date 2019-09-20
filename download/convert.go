package download

import (
  "os"
  "errors"
  "os/exec"
)

func convertWebmToMP3(id string, data []byte) (error) {
  rawFileName := id + ".webm"
  outputFileName := id + ".mp3"

    // Download if does not exists
  if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
    f, err := os.Create(rawFileName)
    f.Write(data)
    f.Close()

    cmd := exec.Command(
      "ffmpeg",
      "-i",
      rawFileName,
      "-vn",
      "-ab",
      "96k",
      "-ar",
      "44100",
      "-y",
      outputFileName)
    err = cmd.Run()
    if err != nil {
      return errors.New("Could not convert to mp3")
    }
  }

  os.Remove(rawFileName)

  return nil
}
