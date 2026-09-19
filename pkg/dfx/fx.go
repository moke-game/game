package dfx

import (
	"github.com/gstones/moke-kit/server/pkg/sfx"
	"github.com/gstones/moke-kit/server/tools"
	"github.com/gstones/moke-kit/utility"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// ProvideFromEnv loads T from environment tags and provides it to the fx graph.
// T should be an fx.Out settings result (envconfig + name tags).
func ProvideFromEnv[T any]() fx.Option {
	return fx.Provide(func() (out T, err error) {
		err = utility.Load(&out)
		return
	})
}

// Dial opens a gRPC connection using the process TLS settings.
func Dial(host string, sec sfx.SecuritySettingsParams) (*grpc.ClientConn, error) {
	if sec.MTLSEnable {
		return tools.DialWithSecurity(
			host,
			sec.ClientCert,
			sec.ClientKey,
			sec.ServerName,
			sec.ServerCaCert,
		)
	}
	return tools.DialInsecure(host)
}

// NewClient dials host and wraps the connection with newFn.
func NewClient[T any](
	host string,
	sec sfx.SecuritySettingsParams,
	newFn func(grpc.ClientConnInterface) T,
) (T, error) {
	var zero T
	conn, err := Dial(host, sec)
	if err != nil {
		return zero, err
	}
	return newFn(conn), nil
}
