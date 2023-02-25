/*
レスポンスの型用のパッケージ（ログアウト）
*/
package auth

type LogoutResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
