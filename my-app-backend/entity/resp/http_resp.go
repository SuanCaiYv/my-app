package resp

import "time"

var notFound = &Result{
	Code: 404,
	Msg:  "Not Fount",
	Data: struct {
	}{},
	Timestamp: time.Now(),
}

var missToken = &Result{
	Code: 461,
	Msg:  "miss token",
	Data: struct {
	}{},
	Timestamp: time.Now(),
}

var authFailed = &Result{
	Code: 462,
	Msg:  "auth failed",
	Data: struct {
	}{},
	Timestamp: time.Now(),
}

type Result struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewOk(data interface{}) *Result {
	return &Result{
		Code:      200,
		Msg:       "",
		Data:      data,
		Timestamp: time.Now(),
	}
}

func NewNotFound() *Result {
	return notFound
}

func NewInternalError(msg string) *Result {
	return &Result{
		Code: 500,
		Msg:  msg,
		Data: struct {
		}{},
		Timestamp: time.Now(),
	}
}

func NewMissToken() *Result {
	return missToken
}

func NewAuthFailed() *Result {
	return authFailed
}

func NewBadRequest(msg string) *Result {
	return &Result{
		Code: 400,
		Msg:  msg,
		Data: struct {
		}{},
		Timestamp: time.Now(),
	}
}

type b struct {
	Result bool `json:"result"`
}

func NewBoolean(result bool) *Result {
	return NewOk(&b{result})
}

type s struct {
	Result string `json:"result"`
}

func NewString(result string) *Result {
	return NewOk(&s{result})
}

type l struct {
	Total       int64       `json:"total"`
	Count       int64       `json:"count"`
	PageNum     int64       `json:"page_num"`
	PageSize    int64       `json:"page_size"`
	NextPageNum int64       `json:"next_page_num"`
	EndPage     bool        `json:"end_page"`
	List        interface{} `json:"list"`
}

func NewList(total, count, pageNum, pageSize, nextPageNum int64, endPage bool, list interface{}) *Result {
	return NewOk(&l{
		Total:       total,
		Count:       count,
		PageNum:     pageNum,
		PageSize:    pageSize,
		NextPageNum: nextPageNum,
		EndPage:     endPage,
		List:        list,
	})
}
