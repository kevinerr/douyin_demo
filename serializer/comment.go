package serializer

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"`
}
