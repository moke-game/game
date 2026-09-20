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

// Author is a stub middleware that extracts a bearer token and does not
// validate it. Production wiring uses platform AuthMiddlewareModule
// (ValidateToken) via Platform / PlatformClients — do not enable
// CustomAuthModule unless you intentionally replace that provider
// (both export name:"AuthMiddleware").
type Author struct {
	unAuthMethods map[string]struct{}
}

// Auth extracts the bearer token. Replace the body if you use CustomAuthModule.
func (d *Author) Auth(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)
	if _, ok := d.unAuthMethods[method]; ok {
		return context.WithValue(ctx, utility.WithoutTag, true), nil
	}
	token, err := auth.AuthFromMD(ctx, string(utility.TokenContextKey))
	if err != nil {
		return ctx, err
	}
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

// CustomAuthModule is an optional stub for tests and local experiments.
// It does not call ValidateToken. Templates should use Platform or PlatformClients.
var CustomAuthModule = fx.Provide(
	func(_ *zap.Logger) (out sfx.AuthMiddlewareResult, err error) {
		out.AuthMiddleware = &Author{
			unAuthMethods: make(map[string]struct{}),
		}
		return
	},
)
