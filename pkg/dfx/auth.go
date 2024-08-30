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

// Author is a game auth service
type Author struct {
	unAuthMethods map[string]struct{}
}

// Auth is a game auth method
func (d *Author) Auth(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)
	if _, ok := d.unAuthMethods[method]; ok {
		return context.WithValue(ctx, utility.WithOutTag, true), nil
	}
	token, err := auth.AuthFromMD(ctx, string(utility.TokenContextKey))
	if err != nil {
		return ctx, err
	}
	// TODO check token with your custom auth middleware
	//CheckToken(token)
	_ = token
	return ctx, nil
}

//AddUnAuthMethod(method string)

// AddUnAuthMethod add you want to disable auth method here
func (d *Author) AddUnAuthMethod(method string) {
	if d.unAuthMethods == nil {
		d.unAuthMethods = make(map[string]struct{})
	}
	d.unAuthMethods[method] = struct{}{}
}

// AuthModule is a game auth module
// you can implement your own auth service or auth function
// Auth will check every rpc/http request,
// if you want to disable it for a service, add `utility.WithoutAuth` in struct of your service
//
//	type service struct {
//		utility.WithoutAuth
//	}
//
// or disable it for a method, add the method name by `AddUnAuthMethod`
var AuthModule = fx.Provide(
	func(
		l *zap.Logger,
	) (out sfx.AuthMiddlewareResult, err error) {
		out.AuthMiddleware = &Author{
			unAuthMethods: make(map[string]struct{}),
		}
		return
	},
)
