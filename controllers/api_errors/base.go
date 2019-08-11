package api_errors

/*
# 1xxx: Authentication errors               没有认证
# 2xxx: HTTP params validation errors       参数错误
# 3xxx: Authorization errors                没有权限
# 4xxx: Record errors
# 5xxx: Server errors
# 6xxx: Other errors

*/

import (
	"net/http"

	"github.com/gonrails/gonrails/controllers/helper"
)

func AuthenticationError(request string) *helper.APIError {
	return helper.NewAPIError(http.StatusUnauthorized, 1000, http.StatusText(http.StatusUnauthorized), request)
}

func AuthorizationError(request string) *helper.APIError {
	return helper.NewAPIError(http.StatusForbidden, 1001, http.StatusText(http.StatusForbidden), request)
}

func ServerError(request string) *helper.APIError {
	return helper.NewAPIError(http.StatusInternalServerError, 5000, http.StatusText(http.StatusInternalServerError), request)
}
