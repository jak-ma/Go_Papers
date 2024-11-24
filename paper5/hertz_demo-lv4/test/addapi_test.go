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

func TestAddapi(t *testing.T) {
	h := server.Default()
	h.POST("/addapi", api.Addapi)

	// 创建表单数据
	formData := url.Values{}
	formData.Set("id", "004")
	formData.Set("name", "testyaml")
	formData.Set("sex", "m")
	formData.Set("comefrom", "golang")

	// 执行 POST 请求
	w := ut.PerformRequest(h.Engine, http.MethodPost, "/addapi", &ut.Body{
		Body: bytes.NewBufferString(formData.Encode()),
		Len:  len(formData.Encode()),
	}, ut.Header{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	})

	// 验证响应状态码
	assert.DeepEqual(t, http.StatusOK, w.Code)

	// 验证响应内容
	expectedResponse := `{"message":"add success","status":200}`
	assert.DeepEqual(t, expectedResponse, w.Body.String())
}
