package detail

import (
	"context"
	"genepse_api/src/infra/objstorage"
	"genepse_api/src/infra/orm"
	"mime/multipart"
)

type Product struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

// ProductOperator is 作品を操作する構造体
type ProductOperator struct {
	ID       int
	UserID   int
	Title    string
	URL      string
	Ctx      context.Context
	File     multipart.File
	ImageURL string
}

func CreateProduct(c *ProductOperator) (res interface{}, err error) {
	imageURL, err := uploadImage(c.Ctx, c.File)
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

func UpdateProduct(c *ProductOperator) (err error) {
	// TODO ない場合はどうしようかな
	c.ImageURL, err = uploadImage(c.Ctx, c.File)
	if err != nil {
		return
	}

	if err = updateProduct(c); err != nil {
		return
	}
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

func updateProduct(c *ProductOperator) (err error) {
	product := &orm.Product{
		Title:        c.Title,
		ReferenceUrl: c.URL,
		ImageUrl:     c.ImageURL,
	}
	return product.Update(uint(c.ID))
}
