package http

import (
	"encoding/json"
	"github.com/EchoGroot/kratos-examples/pkg/lang/mapx"
	"net/http"
	"strings"

	kratoserrors "github.com/go-kratos/kratos/v2/errors"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
)

// ErrorEncoder kratos 返回错误信息处理
// 避免底层错误暴露给调用者
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := kratoserrors.FromError(err)

	// 吞掉具体的错误信息
	if se.Code >= http.StatusInternalServerError {
		se.Message = "内部服务错误"
	}

	codec, _ := kratoshttp.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

const (
	baseContentType = "application"
)

// contentType returns the content-type with base prefix.
func contentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

// NilFilterResponseEncoder 过滤响应 json 中 value 为 nil 的字段，但不过滤其他零值字段，例如 value 为 0、false，或者空数组
func NilFilterResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(kratoshttp.Redirector); ok {
		url, code := rd.Redirect()
		http.Redirect(w, r, url, code)
		return nil
	}

	codec, _ := kratoshttp.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(v)
	if err != nil {
		return err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	mapx.DeleteNil(m)
	data, err = codec.Marshal(m)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", contentType(codec.Name()))
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
