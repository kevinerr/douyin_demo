package serializer

type FavoriteActionResponse struct {
	Response
}

type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
