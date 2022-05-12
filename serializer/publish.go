package serializer

type PublishResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
}
