package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/spf13/cast"
	"net/http"
)

type AuthorizationInfo struct {
	Uid  int64
	Full map[string]interface{}
}

func AuthorizationDecode(authorization string) (info AuthorizationInfo, err error) {
	authVal, _ := base64.StdEncoding.DecodeString(authorization)
	authStr := string(authVal)

	var full map[string]interface{}
	err = json.Unmarshal([]byte(authStr), &full)
	// 格式检测
	if err != nil {
		return
	}
	info = AuthorizationInfo{
		Uid:  cast.ToInt64(full["uid"]),
		Full: full,
	}
	return
}

func GetAuthorizationInfo(request *http.Request) (AuthorizationInfo, error) {
	authorization := request.Header.Get("Authorization")
	if authorization == "" { // 文件上传请求
		authorization = request.Header.Get("X-Upload-Token")
	}

	return AuthorizationDecode(authorization)
}
