/*
レスポンスの型用のパッケージ（パスワードリセット）
*/
package auth

type PasswordResetResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
