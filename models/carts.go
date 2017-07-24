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
 *     Initial: 2017/07/21       Zhu Yaqiang
 *     Modify : 2017/07/22       Xu Haosheng    添加购物车
 *     Modify : 2017/07/23       Wang Ke
<<<<<<< HEAD
<<<<<<< HEAD
 *     Modify : 2017/07/24       Ma Chao
=======
 *     Modify : 2017/07/23 		 Ma Chao
>>>>>>> 3d64342f474377a51e79106608c0092cd3e45f66
=======
 *     Modify : 2017/07/23       Ma Chao
>>>>>>> c9f6e54f8a86371d05c52278ef277322011623e7
 */

package models

import (
	"time"

	"ShopApi/general"
	"ShopApi/orm"
)

type CartsServiceProvider struct {
}

var CartsService *CartsServiceProvider = &CartsServiceProvider{}

<<<<<<< HEAD
type Browse struct {
	Name    string    `json:"name"`
	Count   uint64    `json:"count"`
	Size    string    `json:"size"`
	Color   string    `json:"color"`
	Status  uint8     `json:"status"`
	Created time.Time `json:"created"`
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Image   string    `json:"image"`
	Url     string    `json:"url"`
}

type Cart struct {
	ProductID uint64    `gorm:"column:productid" json:"productid"`
	ImageID   uint64    `gorm:"column:imageid"json:"imageid"`
	Name      string    `json:"name"`
	Size      string    `json:"size"`
	Color     string    `json:"color"`
	Status    uint64    `json:"status"`
	Created   time.Time `json:"created"`
	Count     uint64    `json:"count"`
}

type Images struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Image string `json:"image"`
	Url   string `json:"url"`
}

type Image struct {
	Type  uint64 `json:"type"`
	Title string `json:"title"`
	Image string `json:"image"`
	Url   string `json:"url"`
	ID    uint64 `json:"id"`
}

func (Image) TableName() string {
	return "image"
}

type CartPro struct {
	ID    uint64 `json:"id"`
	Count uint64 `json:"count"`
	Size  string `json:"size"`
	Color string `json:"color"`
}

=======
>>>>>>> 3d64342f474377a51e79106608c0092cd3e45f66
type Carts struct {
	ID        uint64    `sql:"primary_key;" gorm:"column:id" json:"id"`
	ProductID uint64    `gorm:"column:productid" json:"productid"`
	Name      string    `json:"name"`
	Count     uint64    `json:"count"`
	Size      string    `json:"size"`
	Color     string    `json:"color"`
	UserID    uint64    `gorm:"column:userid" json:"userid"`
	ImageID   uint64    `gorm:"column:imageid"json:"imageid"`
	Status    uint8     `json:"status"`
	Created   time.Time `json:"created"`
}

type ConCarts struct {
	ID        uint64    `gorm:"column:id" json:"id" validate:"numeric"`
	ProductID uint64    `gorm:"column:productid" json:"productid" validate:"numeric"`
	Name      string    `json:"name" validate:"required, alphaunicode, min = 2, max = 18"`
	Count     uint64    `json:"count" validate:"numeric"`
	Size      string    `json:"size"`
	Color     string    `json:"color"`
	UserID    uint64    `gorm:"column:userid" json:"userid"`
	ImageID   uint64    `gorm:"column:imageid"json:"imageid" validate:"numeric"`
	Status    uint8     `json:"status" validate:"required, numeric, max = 1"`
	Created   time.Time `json:"created"`
}

func (cs *CartsServiceProvider) CreateInCarts(carts *ConCarts, userID uint64) error {
	cartsPutIn := Carts{
		UserID:    userID,
		ProductID: carts.ProductID,
		Name:      carts.Name,
		Count:     carts.Count,
		Size:      carts.Size,
		Color:     carts.Color,
		ImageID:   carts.ImageID,
		Created:   time.Now(),
	}

	db := orm.Conn
	err := db.Create(&cartsPutIn).Error

	return err
}

// 状态0表示商品在购物车，状态1表示商品不在购物车
func (cs *CartsServiceProvider) CartsDelete(ID uint64, ProID uint64) error {
	var (
		cart Carts
		err  error
	)

	db := orm.Conn
	err = db.Model(&cart).Where("id = ? AND productid = ?", ID, ProID).Update("status", general.ProductNotInCart).Limit(1).Error

	return err
}

func (cs *CartsServiceProvider) AlterCartPro(CartsID uint64, Count uint64) error {
	var (
		cart Carts
	)

	updater := map[string]interface{}{
		"count": Count,
	}

	db := orm.Conn
	err := db.Model(&cart).Where("id = ?", CartsID).Update(updater).Limit(1).Error
	if err != nil {
		return err
	}

	return nil
}

<<<<<<< HEAD
func (cs *CartsServiceProvider) BrowseCart(UserID uint64) ([]ConCarts, error) {
	var (
		err         error
		carts       []ConCarts
		browseCart  []ConCarts
		browse      []ConCarts
=======
func (cs *CartsServiceProvider) BrowseCart(UserID uint64) ([]Carts, error) {
	var (
		err    error
		carts  []Carts
		browse []Carts
>>>>>>> 3d64342f474377a51e79106608c0092cd3e45f66
	)

	db := orm.Conn
	err = db.Where("userid = ?", UserID).Find(&carts).Error
	if err != nil {
		return browse, err
	}

	for _, v := range carts {
<<<<<<< HEAD
		add := ConCarts {
			ImageID:  v.ImageID,
		}
		browseCart = append(browseCart, add)

		add1 := ConCarts{
			ImageID:  v.ImageID,
			Status:  v.Status,
			Created: v.Created,
			Count:   v.Count,
			Name:    v.Name,
			Color:   v.Color,
			Size:    v.Size,
=======
		add1 := Carts{
			ImageID:   v.ImageID,
			Status:    v.Status,
			Created:   v.Created,
			Count:     v.Count,
			Name:      v.Name,
			Color:     v.Color,
			Size:      v.Size,
			ProductID: v.ProductID,
>>>>>>> 3d64342f474377a51e79106608c0092cd3e45f66
		}
		browse = append(browse, add1)
	}

	return browse, err
}
<<<<<<< HEAD

=======
>>>>>>> 3d64342f474377a51e79106608c0092cd3e45f66
