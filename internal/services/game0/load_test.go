package game0

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	kitmod "github.com/gstones/moke-kit/fxmain/pkg/module"
	mqmod "github.com/gstones/moke-kit/mq/pkg/module"
	"github.com/gstones/moke-kit/mq/pkg/mfx"
	nosql "github.com/gstones/moke-kit/orm/pkg/module"
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"go.uber.org/fx"

	"github.com/moke-game/game/pkg/dfx"
)

// TestServiceInstanceLoads starts the real fx providers used by AllModule
// (settings, stub auth, mock document store, redis, local MQ, *Service).
func TestServiceInstanceLoads(t *testing.T) {
	mr := miniredis.RunT(t)
	t.Setenv("DATABASE_URL", "mock://memory")
	t.Setenv("CACHE_URL", "redis://"+mr.Addr())
	t.Setenv("DB_NAME", "game")
	t.Setenv("GAME_URL", "localhost:8081")

	var (
		svc      *Service
		settings dfx.SettingsParams
		auth     sfx.AuthMiddlewareParams
	)
	app := fx.New(
		kitmod.CoreModule,
		nosql.Module,
		mqmod.Module,
		mfx.LocalModule,
		dfx.SettingsModule,
		dfx.CustomAuthModule,
		ServiceInstance,
		fx.NopLogger,
		fx.Populate(&svc, &settings, &auth),
	)
	if err := app.Err(); err != nil {
		t.Fatalf("graph: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Start(ctx); err != nil {
		t.Fatalf("start: %v", err)
	}
	t.Cleanup(func() {
		stopCtx, stopCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer stopCancel()
		_ = app.Stop(stopCtx)
	})

	if svc == nil || svc.gameHandler == nil {
		t.Fatal("ServiceInstance did not construct *Service")
	}
	if settings.DbName != "game" || settings.GameUrl != "localhost:8081" {
		t.Fatalf("settings not loaded: %+v", settings)
	}
	if auth.AuthMiddleware == nil {
		t.Fatal("CustomAuthModule did not provide AuthMiddleware")
	}
}
