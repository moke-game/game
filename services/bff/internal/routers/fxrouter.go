package routers

import (
	"go.uber.org/fx"

	"github.com/moke-game/game/services/bff/internal/biface"
)

type RouterParams struct {
	fx.In

	Routers []biface.IRouter `group:"routers"`
}

type RouterResult struct {
	fx.Out

	Router biface.IRouter `group:"routers"`
}
