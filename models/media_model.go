package models


import (
    "mime/multipart"
    "github.com/jinzhu/gorm"
)

type File struct {
    File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
    Url string `json:"url,omitempty" validate:"required"`
}

type Video struct {
    gorm.Model
	Video_Url  string    `json:"video_url"`
}