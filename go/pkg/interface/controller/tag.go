package controller

import "homepage/pkg/domain/model"

// TagsResponse タグのレスポンス
type TagsResponse struct {

}

// TagResponse タグのレスポンス
type TagResponse struct {
	ID int
	Name string
	CreatedAt string
}

func convertToTagResponse(tag *model.Tag) *TagResponse {
	return &TagResponse{
		ID: tag.ID,
		Name: tag.Name,
		CreatedAt: tag.CreatedAt,
	}
}