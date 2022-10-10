package auth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

func GetAuthorization(request *http.Request) string {
	authorization := request.Header.Get("Authorization")
	if authorization == "" { // 文件上传请求
		authorization = request.Header.Get("X-Upload-Token")
	}

	return authorization
}

type AuthorizationInfo struct {
	Uid  int64
	Full map[string]interface{}
}

func GetAuthorizationInfo(request *http.Request) (info AuthorizationInfo, err error) {
	authorization := GetAuthorization(request)
	authVal, _ := base64.StdEncoding.DecodeString(authorization)
	authStr := string(authVal)

	var authSlide map[string]interface{}
	err = json.Unmarshal([]byte(authStr), &authSlide)
	// 格式检测
	if err != nil {
		return
	}
	info = AuthorizationInfo{
		//Uid:cast.
	}
	return
}
