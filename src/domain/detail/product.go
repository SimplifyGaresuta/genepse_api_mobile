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

func CreateProduct(p *ProductOperator) (res interface{}, err error) {
	// TODO トランザクション
	if p.File != nil {
		p.ImageURL, err = uploadImage(p)
		if err != nil {
			return
		}
	}

	productID, err := insertProduct(p)
	if err != nil {
		return
	}
	res = struct {
		ProductID int `json:"product_id"`
	}{productID}
	return
}

func UpdateProduct(p *ProductOperator) (err error) {
	// TODO トランザクション
	if p.File != nil {
		p.ImageURL, err = uploadImage(p)
		if err != nil {
			return
		}
	}

	if err = updateProduct(p); err != nil {
		return
	}
	return
}

func uploadImage(p *ProductOperator) (imageURL string, err error) {
	imageURL, err = objstorage.Upload(p.Ctx, p.File, objstorage.ProductDir)
	return
}

func insertProduct(p *ProductOperator) (productID int, err error) {
	product := &orm.Product{
		UserId:       uint(p.UserID),
		Title:        p.Title,
		ReferenceUrl: p.URL,
		ImageUrl:     p.ImageURL,
	}
	err = product.Insert()
	productID = int(product.Model.ID)
	return
}

func updateProduct(p *ProductOperator) (err error) {
	product := &orm.Product{
		Title:        p.Title,
		ReferenceUrl: p.URL,
		ImageUrl:     p.ImageURL,
	}
	return product.Update(uint(p.ID))
}
