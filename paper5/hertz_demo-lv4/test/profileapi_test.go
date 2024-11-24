package test

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"go_papers/paper5/hertz_demo-lv2+lv3/api"
	"net/http"
	"net/url"
	"testing"
)

func TestProfile(t *testing.T) {
	h := server.Default()
	h.POST("/profile", api.Profileapi)

	// 创建表单数据
	formData := url.Values{}
	formData.Set("id", "123")
	formData.Set("comefrom", "test golang today")

	//执行POST请求
	w := ut.PerformRequest(h.Engine, "POST", "/profile", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()),
		Len:  len(formData.Encode()),
	}, ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	})

	//验证响应状态码
	assert.DeepEqual(t, http.StatusOK, w.Code)
}
