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
 *     Initial: 2017/07/21        Zhu Yaqiang
 *     Modify: 2017/07/22     Xu Haosheng    添加购物车
 */

package models

import (
	"time"

	"ShopApi/orm"
)

type CartsServiceProvider struct {
}

var CartsService *CartsServiceProvider = &CartsServiceProvider{}

type CartsDel struct {
	ID    uint64 `gorm:"column:id" json:"id"`
	ProID uint64 `json:"productid"`
}

type CartPro struct {
	ID    uint64 `json:"id"`
	Count uint64 `json:"count"`
	Size  string `json:"size"`
	Color string `json:"color"`
}

type Image struct {
	ID   uint64 `json:"id"`
	Url  *string `json:"url"`
	Image *string `json:"image"`
	Type  *string `json:"type"`
	Title *string `json:"title"`
}

type Images struct {
	Url  *string `json:"url"`
	Image *string `json:"image"`
	Type  *string `json:"type"`
	Title *string `json:"title"`
}

func (Image) TableName() string {
	return "image"
}

type Carts struct {
	ID        uint64    `sql:"primary_key;" gorm:"column:id" json:"id"`
	ProductID uint64    `gorm:"column:productid" json:"productid"`
	Name      string    `json:"name"`
	Count     uint64    `json:"count"`
	Size      string    `json:"size"`
	Color     string    `json:"color"`
	UserID    uint64    `gorm:"column:userid" json:"userid"`
	ImageID uint64    `gorm:"column:imageid"json:"imageid"`
	Status    uint64    `json:"status"`
	Created   time.Time `json:"created"`
}

<<<<<<< HEAD
type Cart struct {
	ProductID uint64    `gorm:"column:imageID" json:"productid"`
	ImagineID uint64    `gorm:"column:imageid" json:"imageid"`
	Count     uint64    `json:"count"`
	Status    uint64    `json:"status"`
	Created   time.Time `json:"created"`
}

type Browse struct {
	Name      string    `json:"name"`
	Size      string    `json:"size"`
	Color     string    `json:"color"`
	Count     uint64    `json:"count"`
	Status    uint64    `json:"status"`
	Created   time.Time `json:"created"`
	Url  *string `json:"url"`
	Image *string `json:"image"`
	Type  *string `json:"type"`
	Title *string `json:"title"`
}

func (Carts) TableName() string {
	return "carts"
}

func (cs *CartsServiceProvider) WhetherInCart(CartsID uint64) error {
	var (
		err  error
		cart Carts
	)
=======
func (cs *CartsServiceProvider) CreateInCarts(carts Carts, userID uint64) error {
	cartsPutIn := Carts {
		UserID:                       userID,
		ProductID:                 carts.ProductID,
		Name:                        carts.Name,
		Count:                        carts.Count,
		Size:                            carts.Size,
		Color:                          carts.Color,
		ImageID:                    carts.ImageID,
		Status:                        carts.Status,
		Created:                     time.Now(),
	}
>>>>>>> 9b3a594449314b1bc769c4216c23fdbba9529f8f

	db := orm.Conn

	err := db.Create(&cartsPutIn).Error
	if err != nil {
		return err
	}

	return nil
}

// 状态1表示商品在购物车，状态0表示商品不在购物车
func (cs *CartsServiceProvider) CartsDelete (ID uint64, ProID uint64) error {
	var (
		cart Carts
		err  error
	)

	db := orm.Conn

	err = db.Where("id = ? and productid = ?", ID, ProID).First(&cart).Error
	if err != nil {
		return err
	}

	err = db.Model(&cart).Where("id = ? and productid = ?", ID, ProID).Update("status", 0).Limit(1).Error
	if err != nil {
		return err
	}

	return nil
}

func (cs *CartsServiceProvider) AlterCartPro(CartsID uint64, Count uint64, Size string, Color string) error {
	var (
		cart Carts
	)
	updater := map[string]interface{}{
		"count": Count,
		"size":  Size,
		"color": Color,
	}

	db := orm.Conn
	err := db.Model(&cart).Where("id = ?", CartsID).Update(updater).Limit(1).Error
	if err != nil {
		return err
	}

	return nil
}




