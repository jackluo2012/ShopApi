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
 *     Initial: 2017/07/18        Li Zebang
 *     Modify: 2017/07/20        Yu Yi
 *     Modify: 2017/07/20        Yang Zhengtian
 */

package models

import (
	"time"

	"ShopApi/orm"
	"fmt"
)

type Contact struct {
	ID        uint64    `sql:"auto_increment; primary_key;" json:"id"`
	UserID    uint64    `gorm:"column:userid" json:"userid"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	Address   string    `json:"address"`
	Created   time.Time `json:"created"`
	IsDefault  uint8	`gorm:"column:isdefault" json:"isdefault"`
}

type Addressget struct {
	Province string 	`json:"province"`
	City     string		`json:"city"`
	Street   string		`json:"street"`
	Address  string 	`json:"address"`
}

type ContactServiceProvider struct {
}

var ContactService *ContactServiceProvider = &ContactServiceProvider{}

func (Contact) TableName() string {
	return "contact"
}

func (csp *ContactServiceProvider) AddAddress(contact *Contact) error {
	contact.Created = time.Now()

	db := orm.Conn

	err := db.Create(contact).Error
	if err != nil {
		return err
	}

	return nil
}

func (us *ContactServiceProvider) ChangeAddress(id *uint64, name, phone, province, city, street, address *string) error {

	changmap := map[string]interface{}{"name": *name, "phone": *phone, "province": *province, "city": *city, "street": *street, "address": *address}

	// todo: 传入的结构先声明
	db := orm.Conn
	err := db.Model(&Contact{}).Where(&Contact{ID: *id}).Updates(changmap).Limit(1).Error

	if err != nil {
		return err
	}

	return nil
}

func (us *ContactServiceProvider) GetAddress(userid uint64) ([]Addressget, error) {
	var (
		cont  Contact
		list  []Contact
		s     []Addressget
	)

	db := orm.Conn
	err := db.Model(&cont).Where("userid=?", userid).Find(&list).Error
	if err != nil {
		return s, err
	}

	for i, c := range list{
		s[i].Province = c.Province
		s[i].City = c.City
		s[i].Street = c.Street
		s[i].Address = c.Address
	}

	return s, nil
}

func(us *ContactServiceProvider) AlterDefalt(id uint64) error{
	var(
		s	Contact
		a	int8
		con	Contact
	)
	db := orm.Conn
	err := db.Where("id=?",id).Find(&s).Error
	if err != nil {
	return err
	}
	if s.IsDefault == 0{
		a = 1
	}
	updater := map[string]interface{}{"isdefault": a}
	err = db.Model(&con).Where("id=?",id).Update(updater).Limit(1).Error
	if err != nil {
		return err
	}

	return nil
}
