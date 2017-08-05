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
 *     Initial: 2017/08/05       Ai Hao
 */

package errcode

const (
	//Create
	CreateProductSucceed      = 0x0
	ErrCreateInvalidParams    = 0x1

	// GetList
	GetListSucceed            = 0x0
	ErrGetInvalidParams       = 0x1
	ErrCategoryNotFound       = 0x2

	//ChangeStatus
	ErrChangeInvalidParams    = 0x0
	ErrStatusNotFound         = 0x1
	ChangeStatusSucceed       = 0x2

	//GetInfo
	ErrInfoInvalidParams      = 0x0
	GetInfoSucceed            = 0x1

	//ChangeCategory
	ErrCategoryInvalidParams  = 0x0
	ChangeCategorySucceed     = 0x1
)