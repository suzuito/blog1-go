package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/internal/entity"
)

// GetAdminAuth ...
func (u *Impl) GetAdminAuth(
	ctx context.Context,
	headerAdminAuth string,
	adminAuth *entity.AdminAuth,
) error {
	return nil
}
