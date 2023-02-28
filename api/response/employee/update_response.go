/*
レスポンスの型用のパッケージ（従業員）
*/
package employee

type UpdateEmployeeResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
