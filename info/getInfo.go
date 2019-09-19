package info

import (
  "github.com/rylio/ytdl"
)

func GetInfo(id string) (SongInfo, error) {
  watchUrl := "https://www.youtube.com/watch?v=" + id

  vid, err := ytdl.GetVideoInfo(watchUrl)
  if err != nil {
    return SongInfo{}, err
  }

  result := SongInfo {
    ID: id,
    Title: vid.Title,
  }

  return result, nil
}
