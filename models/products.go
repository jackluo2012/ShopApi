/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/07/21         Ai Hao
 *     Modify : 2017/07/21         Zhu Yaqiang
 *     Modify : 2017/07/21         Yu Yi
 */

package models

import (
	"time"

	"ShopApi/general"
	"ShopApi/orm"
)

type ProductServiceProvider struct {
}

var ProductService *ProductServiceProvider = &ProductServiceProvider{}

type Product struct {
	ID            uint64    `sql:"auto_increment;primary_key;" gorm:"column:id" json:"id"`
	Name          string    `json:"name"`
	TotalSale     uint64    `gorm:"column:totalsale" json:"totalsale"`
	Category      uint64    `json:"categories"`
	Price         float64   `json:"price"`
	OriginalPrice float64   `gorm:"column:originalprice" json:"originalprice"`
	Status        uint64    `json:"status"`
	Size          string    `json:"size"`
	Color         string    `json:"color"`
	ImageID       uint64    `gorm:"column:imageid" json:"imageid"`
	ImageIDs      string    `gorm:"column:imageids" json:"imageids"`
	Remark        string    `json:"remark"`
	Detail        string    `json:"detail"`
	Created       time.Time `json:"created"`
	Inventory     uint64    `json:"inventory"`
}

type ConProduct struct {
	ID            uint64    `gorm:"column:id" json:"id" validate:"numeric"`
	Name          string    `json:"name" validate:"required, alphaunicode, min = 2, max = 18"`
	TotalSale     uint64    `gorm:"column:totalsale" json:"totalsale" validate:"numeric"`
	Category      uint64    `json:"categories" validate:"numeric"`
	Price         float64   `json:"price" validate:"numeric"`
	OriginalPrice float64   `gorm:"column:originalprice" json:"originalprice" validate:"numeric"`
	Status        uint64    `json:"status" validate:"numeric"`
	Size          string    `json:"size"`
	Color         string    `json:"color"`
	ImageID       uint64    `gorm:"column:imageid" json:"imageid" validate:"numeric"`
	ImageIDs      string    `gorm:"column:imageids" json:"imageids"`
	Remark        string    `json:"remark"`
	Detail        string    `json:"detail"`
	Created       time.Time `json:"created"`
	Inventory     uint64    `json:"inventory"`
}

type GetCategories struct {
	Categories uint64 `json:"categories" validate:"required, alphanum, min = 0, max= 30"`
}

type GetProList struct {
	Name          string
	TotalSale     uint64
	Price         float64
	OriginalPrice float64
	Status        uint64
	ImageId       uint64
	Detail        string
	Inventory     uint64
}

type ChangePro struct {
	ID     uint64 `json:"id" validate:"numeric"`
	Status uint64 `json:"status" validate:"numeric"`
}

type ChangeCate struct {
	ID         uint64 `json:"id"`
	Categories uint64 `json:"categories"`
}

func (Product) TableName() string {
	return "products"
}

func (ps *ProductServiceProvider) CreateProduct(pr *ConProduct) error {
	pro := Product{
		Name:			pr.Name,
		TotalSale:		pr.TotalSale,
		Category: 		pr.Category,
		Price:			pr.Price,
		OriginalPrice:	pr.OriginalPrice,
		Size: 			pr.Size,
		Color: 			pr.Color,
		ImageID:		pr.ImageID,
		ImageIDs:		pr.ImageIDs,
		Detail:			pr.Detail,
		Inventory:		pr.Inventory,
	}

	pro.Status = general.ProductOnsale
	pro.Created = time.Now()

	db := orm.Conn
	err := db.Create(&pro).Error

	return err
}

func (ps *ProductServiceProvider) GetProduct(cate uint64) ([]GetProList, error) {
	var (
		ware Product
		list []Product
		s    []GetProList
	)

	db := orm.Conn
	err := db.Model(&ware).Where("categories = ?", cate).Find(&list).Error

	if err != nil {
		return s, err
	}

	for _, c := range list {
		if c.Status == general.ProductOnsale {
			pro := GetProList{
				Name:          c.Name,
				TotalSale:     c.TotalSale,
				Price:         c.Price,
				OriginalPrice: c.OriginalPrice,
				Status:        c.Status,
				ImageId:       c.ImageID,
				Detail:        c.Detail,
				Inventory:     c.Inventory,
			}
			s = append(s, pro)
		}
	}

	return s, nil
}

func (ps *ProductServiceProvider) ChangeProStatus(ID uint64, status uint64) error {
	var (
		pro Product
		err error
	)

	change := map[string]interface{}{"status": status}
	db := orm.Conn

	err = db.Model(&pro).Where("id = ?", ID).Updates(change).Limit(1).Error

	return err
}

func (ps *ProductServiceProvider) GetProInfo(ProID uint64) (*Product, error) {
	var (
		err     error
		ProInfo *Product = &Product{}
	)

	db := orm.Conn

	err = db.Where("id = ?", ProID).First(&ProInfo).Error
	if err != nil {
		return ProInfo, err
	}

	return ProInfo, err
}

func (ps *ProductServiceProvider) ChangeCategories(cate ChangeCate) error {
	var (
		pro Product
	)

	// todo: 复数
	change := map[string]uint64{"categories": cate.Categories}

	db := orm.Conn
	err := db.Model(&pro).Where("id = ?", cate.ID).Update(change).Limit(1).Error

	if err != nil {
		return err
	}

	return nil
}
