package dto

import "mime/multipart"

type UploadFileRequest struct {
	File multipart.FileHeader `form:"file" binding:"required"`
}

type UploadFileResponse struct {
	Link     string `json:"link"`
	FileName string `json:"file_name"`
}

type UploadFilesRequest struct {
	Files []multipart.FileHeader `form:"files" binding:"required"`
}

type UploadFilesResponse struct {
	Files []UploadFileResponse `json:"files"`
}

type DeleteFileRequest struct {
	Link string `json:"link"  binding:"required"`
}

type DeleteFileResponse struct {
	Message string `json:"message" `
}

type DeleteFilesRequest struct {
	Links []string `json:"links"  binding:"required"`
}

type DeleteFilesResponse struct {
	Message string `json:"message" `
}
