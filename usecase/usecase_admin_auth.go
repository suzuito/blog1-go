package usecase

import (
	"context"

	"github.com/suzuito/blog1-go/entity/model"
)

// GetAdminAuth ...
func (u *Impl) GetAdminAuth(
	ctx context.Context,
	headerAdminAuth string,
	adminAuth *model.AdminAuth,
) error {
	return nil
}
