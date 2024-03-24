package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Error 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	res := Default.Clone()
	if err != nil {
		res.SetMsg(err.Error())
	}
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(GenerateMsgIDFromContext(c))
	res.SetCode(int32(code))
	res.SetSuccess(false)
	c.Set("result", res)
	c.Set("status", code)
	c.AbortWithStatusJSON(code, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	res := Default.Clone()
	res.SetResult(data)
	res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(GenerateMsgIDFromContext(c))
	res.SetCode(http.StatusOK)
	c.Set("result", res)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.Items = result
	res.Total = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

// Custum 兼容函数
//func Custum(c *gin.Context, data gin.H) {
//	data["requestId"] = pkg.GenerateMsgIDFromContext(c)
//	c.Set("result", data)
//	c.AbortWithStatusJSON(http.StatusOK, data)
//}

type Response struct {
	// 数据集
	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Code      int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

type response struct {
	Response
	Result interface{} `json:"result"`
}

type Page struct {
	Total     int `json:"total"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
type page struct {
	Page
	Items interface{} `json:"items"`
}

var Default = &response{}

func (e *response) SetResult(result interface{}) {
	e.Result = result
}

func (e response) Clone() *response {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Message = s
}

func (e *response) SetCode(code int32) {
	e.Code = code
}

func (e *response) SetSuccess(success bool) {
	if !success {
		e.Status = "error"
	}
}
