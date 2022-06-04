package repository

import (
	"github.com/RaymondCode/simple-demo/model"
)

type CommentRepository struct {
}

//insert one comment
func (c CommentRepository) CreatComment(comment *model.Comment) error {
	err := model.DB.Create(comment).Error
	return err
}

//get video id by comment id
func (c CommentRepository) GetVideoId(commentId int64) (videoId int64, err error) {
	comment := model.Comment{}
	err = model.DB.Where("id=?", commentId).First(&comment).Error
	return comment.VideoId, err
}

//delete one comment
func (c CommentRepository) DeleteComment(commentId int64) error {
	comment := model.Comment{Id: commentId}
	err := model.DB.Delete(&comment).Error
	return err
}

//select comments of a video
func (c CommentRepository) SelectComments(videoId int64, comments *[]model.Comment) error {
	err := model.DB.Where("video_id=?", videoId).Order("create_time DESC").Find(&comments).Error
	return err
}

//get the author of the commit
func (c CommentRepository) GetAuthorId(commentId int64) (authorId int64, error error) {
	var comment model.Comment
	err := model.DB.Where("id=?", commentId).First(&comment).Error
	return comment.UserId, err
}
