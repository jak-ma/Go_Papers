package test

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"go_papers/paper5/hertz_demo-lv2+lv3/api"
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	h := server.Default()
	h.GET("/search", api.Searchapi)

	w := ut.PerformRequest(h.Engine, "GET", "/search?id=001", &ut.Body{
		Body: bytes.NewBufferString("1"),
		Len:  1,
	}, ut.Header{
		"Connection", "close",
	})
	resp := w.Result()
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
	exresp := `{
	"msg":"success",
	"optime":"2024-11-23 14:28:31",
	"stucomefrom":"beijing",
	"stuid":"001",
	"stuname":"xiaoming",
	"stusex":"m"}`                     //todo 这里一直测试通不过，看着返回的两
	                                  // todo 个字符串都是一样的，问题出在哪呢，是不是断言失败了？
	assert.DeepEqual(t, exresp, string(resp.Body()))
}
