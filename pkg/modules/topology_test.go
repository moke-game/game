package modules

import (
	"testing"

	kitmod "github.com/gstones/moke-kit/fxmain/pkg/module"
	"go.uber.org/fx"
)

func TestAggregateGraph(t *testing.T) {
	t.Parallel()
	if err := fx.ValidateApp(kitmod.AppModule, Aggregate); err != nil {
		t.Fatalf("Aggregate graph: %v", err)
	}
}

func TestThinGraph(t *testing.T) {
	t.Parallel()
	if err := fx.ValidateApp(kitmod.AppModule, Thin); err != nil {
		t.Fatalf("Thin graph: %v", err)
	}
}

func TestTransportGraphs(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		opt  fx.Option
	}{
		{"GrpcModule", GrpcModule},
		{"HttpModule", HttpModule},
		{"TcpModule", TcpModule},
		{"AllModule", AllModule},
		{"AllWithTCPModule", AllWithTCPModule},
		{"GrpcClientModule", GrpcClientModule},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if err := fx.ValidateApp(kitmod.AppModule, Infra, PlatformClients, tc.opt); err != nil {
				t.Fatal(err)
			}
		})
	}
}
