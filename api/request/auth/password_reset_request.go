/*
フロントからのリクエストのstruct用のパッケージ
*/
package auth

type PasswordResetRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
