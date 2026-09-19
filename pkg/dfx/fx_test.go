package dfx

import (
	"context"
	"testing"

	"go.uber.org/fx"
)

func TestProvideFromEnvLoads(t *testing.T) {
	t.Setenv("GAME_URL", "example:9")
	t.Setenv("DB_NAME", "testdb")

	var got SettingsParams
	app := fx.New(
		ProvideFromEnv[SettingsResult](),
		fx.NopLogger,
		fx.Populate(&got),
	)
	if err := app.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = app.Stop(context.Background()) })

	if got.GameUrl != "example:9" || got.DbName != "testdb" {
		t.Fatalf("SettingsParams = %+v, want GAME_URL=example:9 DB_NAME=testdb", got)
	}
}
