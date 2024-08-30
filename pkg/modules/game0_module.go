package modules

import (
	"go.uber.org/fx"

	"github.com/moke-game/game/internal/services/game0"

	"github.com/moke-game/game/pkg/dfx"
)

// GrpcModule  grpc service module
// can inject it to any fxmain.Main(), if you want to start a game grpc service
var GrpcModule = fx.Module("grpcService",
	dfx.SettingsModule,
	dfx.AuthModule,
	game0.GrpcService,
)

// HttpModule  http service module
// can inject it to any fxmain.Main(), if you want to start a game http service
var HttpModule = fx.Module("httpService",
	dfx.SettingsModule,
	dfx.AuthModule,
	game0.GrpcService,
	game0.HttpService,
)

// TcpModule  tcp service module
// can inject it to any fxmain.Main(), if you want to start a game tcp service
var TcpModule = fx.Module("tcpService",
	dfx.SettingsModule,
	game0.TcpService,
)

// AllModule  all service module
// can inject it to any fxmain.Main(), if you want to start game all type services
var AllModule = fx.Module("allService",
	//dfx.AuthModule,
	dfx.SettingsModule,
	game0.GrpcService,
	game0.HttpService,
	game0.TcpService,
)

// GrpcClientModule  grpc client module
// can inject it to any fxmain.Main(), if you want a game grpc client to rpc game service
var GrpcClientModule = fx.Module("grpcClient",
	dfx.SettingsModule,
	dfx.Game0ClientModule,
)
