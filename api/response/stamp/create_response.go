/*
レスポンスの型用のパッケージ（出退勤）
*/
package stamp

type CreateStampResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
