package detail

import (
	"context"
	"genepse_api/src/infra/objstorage"
	"genepse_api/src/infra/orm"
	"log"
	"mime/multipart"
)

type Product struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

type ProductCreator struct {
	UserID int
	Title  string
	URL    string
	Ctx    context.Context
	File   multipart.File
}

func CreateProduct(c *ProductCreator) (res interface{}, err error) {
	imageURL, err := uploadImage(c.Ctx, c.File)
	log.Println("urlは", imageURL)
	if err != nil {
		return
	}

	productID, err := insertProduct(uint(c.UserID), c.Title, c.URL, imageURL)
	if err != nil {
		return
	}
	res = struct {
		ProductID int `json:"product_id"`
	}{productID}
	return
}

func uploadImage(ctx context.Context, file multipart.File) (imageURL string, err error) {
	imageURL, err = objstorage.Upload(ctx, file, objstorage.ProductDir)
	return
}

func insertProduct(userID uint, title, referenceURL, imageURL string) (productID int, err error) {
	product := &orm.Product{
		UserId:       userID,
		Title:        title,
		ReferenceUrl: referenceURL,
		ImageUrl:     imageURL,
	}
	err = product.Insert()
	productID = int(product.Model.ID)
	return
}
