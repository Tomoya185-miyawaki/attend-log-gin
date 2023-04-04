/*
レスポンスの型用のパッケージ（出退勤）
*/
package stamp

type DeleteStampResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
