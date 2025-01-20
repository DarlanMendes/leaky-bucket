package service

import "leaky-bucket/api/model"

func GetAll(page int, size int) (*model.Pagination, error) {
	var message []model.Message
	var count int64
	if page == 0 {
		page = 1
	}
	//offset := (page - 1) * size
	pagination := &model.Pagination{
		Content: message,
		Page:    page,
		Size:    size,
		Total:   count,
	}
	return pagination, nil
}
