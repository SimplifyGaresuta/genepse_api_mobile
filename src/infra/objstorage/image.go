package objstorage

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

// TODO (アイデアベース)imageを構造体にする
func Upload(ctx context.Context, file io.Reader, directory string) (imageURL string, err error) {
	t := time.Now()
	now := t.Format("Mon Jan 2 15:04:05 MST 2006")
	client, err := storage.NewClient(ctx)
	filePath := fmt.Sprintf("%s/%s%s", directory, now, ".jpg")
	blobWriter := client.Bucket(Bucket).Object(filePath).NewWriter(ctx)
	blobWriter.ContentType = "image/jpeg"
	io.Copy(blobWriter, file)
	if err := blobWriter.Close(); err != nil {
		return "", err
	}
	imageURL = fmt.Sprintf("%s/%s/%s", Host, Bucket, filePath)
	return imageURL, err
}
