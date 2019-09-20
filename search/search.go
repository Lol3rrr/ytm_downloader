package search

type SearchVideo struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Channel string `json:"channel"`
}

type SearchResult struct {
  Videos []SearchVideo `json:"videos"`
}
