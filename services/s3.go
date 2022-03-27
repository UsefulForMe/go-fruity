package services

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/UsefulForMe/go-ecommerce/config"
	"github.com/UsefulForMe/go-ecommerce/dto"
	"github.com/UsefulForMe/go-ecommerce/errs"
	"github.com/UsefulForMe/go-ecommerce/logger"
	"github.com/UsefulForMe/go-ecommerce/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type S3Service struct {
	S3Client *s3.S3
	Session  *session.Session
}

func getKeyFromLink(link string) (string, error) {
	prefix := config.Cfg.AWSS3Endpoint + "/"
	if !strings.HasPrefix(link, prefix) {
		return "", errors.New("Link not found")
	}

	return strings.Replace(link, prefix, "", 1), nil
}

func uploadToS3(s S3Service, file multipart.FileHeader, userID uuid.UUID) (*dto.UploadFileResponse, *errs.AppError) {

	data, err := utils.MultipartFileToByte(file)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error when convert file to bytes " + err.Error())
	}

	m := regexp.MustCompile(`(\s+|//)`)
	fileName := m.ReplaceAllString(file.Filename, "_")

	contentType := file.Header.Get("Content-Type")
	unixTime := time.Now().Unix()

	key := fmt.Sprintf("%s/%s_%s", userID, fileName, strconv.FormatInt(unixTime, 10))
	inputFile := s3.PutObjectInput{
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
		Bucket:      &config.Cfg.AWSBucket,
		Key:         aws.String(key),
	}
	_, err = s.S3Client.PutObject(&inputFile)
	if err != nil {
		return nil, errs.NewUnexpectedError("Failed to upload file " + err.Error())
	}
	link := fmt.Sprintf("%s/%s", config.Cfg.AWSS3Endpoint, key)
	res := dto.UploadFileResponse{
		Link:     link,
		FileName: fileName,
	}
	return &res, nil
}

func (s S3Service) UploadFile(r dto.UploadFileRequest, userID uuid.UUID) (*dto.UploadFileResponse, *errs.AppError) {
	res, err := uploadToS3(s, r.File, userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s S3Service) UploadFiles(r dto.UploadFilesRequest, userID uuid.UUID) (*dto.UploadFilesResponse, *errs.AppError) {
	responses := make([]dto.UploadFileResponse, 0)
	for _, file := range r.Files {
		res, err := uploadToS3(s, file, userID)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *res)
	}
	res := dto.UploadFilesResponse{
		Files: responses,
	}
	return &res, nil
}

func (s S3Service) DeleteFile(r dto.DeleteFileRequest) (*dto.DeleteFileResponse, *errs.AppError) {

	key, err := getKeyFromLink(r.Link)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error when get key from link " + err.Error())
	}
	_, err = s.S3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &config.Cfg.AWSBucket,
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, errs.NewUnexpectedError("Error when delete file " + err.Error())
	}
	return &dto.DeleteFileResponse{
		Message: "Delete file success",
	}, nil
}

func (s S3Service) DeleteFiles(r dto.DeleteFilesRequest) (*dto.DeleteFilesResponse, *errs.AppError) {

	keys := make([]*s3.ObjectIdentifier, 0)
	for _, link := range r.Links {
		key, err := getKeyFromLink(link)
		if err != nil {
			return nil, errs.NewUnexpectedError("Error when get key from link " + err.Error())
		}
		keys = append(keys, &s3.ObjectIdentifier{
			Key: aws.String(key),
		})

	}
	_, err := s.S3Client.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(config.Cfg.AWSBucket),
		Delete: &s3.Delete{
			Objects: keys,
		},
	})

	if err != nil {
		return nil, errs.NewUnexpectedError("Error when delete files " + err.Error())
	}
	return &dto.DeleteFilesResponse{
		Message: "Delete files success",
	}, nil
}

func NewS3Service() S3Service {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Cfg.AWSRegion), Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{
				Value: credentials.Value{
					AccessKeyID:     config.Cfg.AWSID,
					SecretAccessKey: config.Cfg.AWSSecret,
				},
			},
		)},
	)
	if err != nil {
		logger.Error("Failed to create S3 session " + err.Error())
		panic(err)
	}
	return S3Service{
		Session:  sess,
		S3Client: s3.New(sess),
	}

}
