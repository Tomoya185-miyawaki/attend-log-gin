/*
レスポンスの型用のパッケージ（ログイン）
*/
package auth

type LoginResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
