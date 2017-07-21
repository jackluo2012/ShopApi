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
 *     Initial: 2017/07/21        Yang Zhengtian
 */

package handler

import (
	"ShopApi/log"
	"ShopApi/models"
	"github.com/labstack/echo"
	"ShopApi/general"
	"ShopApi/general/errcode"
)

type Pid struct {
	Pid uint64 `json:"pid"`
}

func CreateC(c echo.Context) error {
	var (
		err error
		cate models.CreateCat
	)
	if err = c.Bind(&cate); err != nil {
		log.Logger.Error("Create crash with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	err = models.CategoriesService.Create(cate)
	if err != nil {
		log.Logger.Error("Create crash with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func GetCategories(c echo.Context) error {
	var (
		err        error
		pid        Pid
		categories []models.Categories
	)

	if err = c.Bind(&pid); err != nil {
		log.Logger.Error("Bind with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	categories, err = models.CategoriesService.GetCategories(pid.Pid)
	if err != nil {
		log.Logger.Error("error:", err)
		return general.NewErrorWithMessage(errcode.ErrGetCategories, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, categories)
}