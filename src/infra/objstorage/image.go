package objstorage

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
)

const (
	bucket = "genepse-186713.appspot.com"
)

func Upload(req *http.Request, file multipart.File, directory string) (string, error) {
	t := time.Now()
	now := t.Format("Mon Jan 2 15:04:05 MST 2006")
	fmt.Println("今は", now)
	fmt.Println("ファイルは", file)
	ctx := req.Context()
	fmt.Println("コンテキストは", ctx)
	client, err := storage.NewClient(ctx)
	//file, err := os.Open("test.jpg")
	filePath := fmt.Sprintf("%s/%s%s", directory, now, ".jpg")
	blobWriter := client.Bucket(bucket).Object(filePath).NewWriter(ctx)
	fmt.Println("ライターは", blobWriter)
	blobWriter.ContentType = "image/jpeg"
	io.Copy(blobWriter, file)
	err = blobWriter.Close()
	blobPath := fmt.Sprintf("/gs/%s/%s", bucket, filePath)
	blobKey, err := blobstore.BlobKeyForFile(ctx, blobPath)

	// crop
	opts := image.ServingURLOptions{Secure: false, Crop: true}
	url, err := image.ServingURL(ctx, blobKey, &opts)
	return url.String(), err
}
