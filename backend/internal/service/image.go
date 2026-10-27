package service

import (
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"quiz.com/quiz/internal/collection"
)

type ImageService struct {
	imageCollection *collection.ImageCollection
}

func Image(imageCollection *collection.ImageCollection) *ImageService {
	return &ImageService{
		imageCollection: imageCollection,
	}
}

func (s *ImageService) UploadImage(filename string, contentType string, reader io.Reader) (primitive.ObjectID, error) {
	return s.imageCollection.Upload(filename, contentType, reader)
}

func (s *ImageService) DownloadImage(id primitive.ObjectID) (*gridfs.DownloadStream, string, int64, error) {
	return s.imageCollection.Download(id)
}

func (s *ImageService) DeleteImage(id primitive.ObjectID) error {
	return s.imageCollection.Delete(id)
}
