package dfx

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Author is an optional custom auth middleware example.
//
// Default wiring uses platform AuthMiddlewareModule (ValidateToken) via
// pkg/modules.AllModule — do not enable CustomAuthModule unless you intentionally
// replace that middleware (both export name:"AuthMiddleware").
type Author struct {
	unAuthMethods map[string]struct{}
}

// Auth extracts the bearer token. Replace the TODO body if you use CustomAuthModule.
func (d *Author) Auth(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)
	if _, ok := d.unAuthMethods[method]; ok {
		return context.WithValue(ctx, utility.WithOutTag, true), nil
	}
	token, err := auth.AuthFromMD(ctx, string(utility.TokenContextKey))
	if err != nil {
		return ctx, err
	}
	// Custom deployments: validate token here (or prefer platform AuthMiddlewareModule).
	_ = token
	return ctx, nil
}

// AddUnAuthMethod marks a full method name as unauthenticated.
func (d *Author) AddUnAuthMethod(method string) {
	if d.unAuthMethods == nil {
		d.unAuthMethods = make(map[string]struct{})
	}
	d.unAuthMethods[method] = struct{}{}
}

// CustomAuthModule is an optional stub middleware for experiments only.
// Production/templates should use:
//
//	auth "github.com/moke-game/platform/services/auth/pkg/module"
//	auth.AuthMiddlewareModule
var CustomAuthModule = fx.Provide(
	func(
		l *zap.Logger,
	) (out sfx.AuthMiddlewareResult, err error) {
		out.AuthMiddleware = &Author{
			unAuthMethods: make(map[string]struct{}),
		}
		return
	},
)

// AuthModule is kept as an alias so older references compile, but it is not
// used by AllModule anymore. Prefer platform AuthMiddlewareModule.
var AuthModule = CustomAuthModule
