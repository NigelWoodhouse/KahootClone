package collection

import (
	"io"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ImageCollection struct {
	bucket *gridfs.Bucket
}

func Image(database *mongo.Database) (*ImageCollection, error) {
	bucket, err := gridfs.NewBucket(database, options.GridFSBucket().SetName("images"))
	if err != nil {
		return nil, err
	}
	return &ImageCollection{
		bucket: bucket,
	}, nil
}

func (c *ImageCollection) Upload(filename string, contentType string, source io.Reader) (primitive.ObjectID, error) {
	opts := options.GridFSUpload().SetMetadata(bson.M{
		"contentType": contentType,
	})
	return c.bucket.UploadFromStream(filename, source, opts)
}

func (c *ImageCollection) Download(id primitive.ObjectID) (*gridfs.DownloadStream, string, int64, error) {
	stream, err := c.bucket.OpenDownloadStream(id)
	if err != nil {
		return nil, "", 0, err
	}

	file := stream.GetFile()
	contentType := "application/octet-stream"
	var length int64
	if file != nil {
		length = file.Length
		if len(file.Metadata) > 0 {
			var meta struct {
				ContentType string `bson:"contentType"`
			}
			if err := bson.Unmarshal(file.Metadata, &meta); err == nil && meta.ContentType != "" {
				contentType = meta.ContentType
			}
		}
	}

	return stream, contentType, length, nil
}

func (c *ImageCollection) Delete(id primitive.ObjectID) error {
	return c.bucket.Delete(id)
}
