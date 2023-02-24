/*
シーダー用パッケージ
*/
package seed

import (
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
)

var Admin = &dto.Admin{
	ID:       1,
	Name:     "管理者",
	Email:    "admin@example.com",
	Password: "$2a$12$jA9fyFy1qfh379BEfaJM2uGPe9EnqrEZnREv1iaiX8nyCmQz69ERK", // passwordをハッシュ化したもの
}
