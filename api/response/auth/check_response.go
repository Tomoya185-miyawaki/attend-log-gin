/*
認証チェック用のレスポンスの型
*/
package auth

type AuthCheckResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
