/*
レスポンスの型用のパッケージ（タイムアウト）
*/
package timeout

type TimeoutResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
