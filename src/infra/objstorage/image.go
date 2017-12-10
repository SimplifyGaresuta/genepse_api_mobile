package objstorage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
	"google.golang.org/appengine/log"
)

const (
	bucket = "genepse-186713.appspot.com"
)

func Upload(req *http.Request) (string, error) {
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	file, err := os.Open("test.jpg")
	userName := "tester"
	fileName := "test.jpg"
	filePath := fmt.Sprintf("%s/%s", userName, fileName)
	blobWriter := client.Bucket(bucket).Object(filePath).NewWriter(ctx)
	blobWriter.ContentType = "image/jpeg"
	io.Copy(blobWriter, file)
	err = blobWriter.Close()
	blobPath := fmt.Sprintf("/gs/%s/%s", bucket, filePath)
	blobKey, err := blobstore.BlobKeyForFile(ctx, blobPath)

	// crop
	opts := image.ServingURLOptions{Secure: false, Crop: true}
	url, err := image.ServingURL(ctx, blobKey, &opts)
	log.Infof(ctx, "url", url)
	return url.String(), err
}

func TestUpload(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	req, err := inst.NewRequest("GET", "/gophers", nil)
	url, err := Upload(req)
	fmt.Printf("test", url, err)
}
