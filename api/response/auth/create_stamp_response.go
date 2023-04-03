/*
レスポンスの型用のパッケージ（出退勤）
*/
package auth

type CreateStampResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
