/*
レスポンスの型用のパッケージ（出退勤）
*/
package stamp

type UpdateStampResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
