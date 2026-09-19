package main

import (
	"github.com/gstones/moke-kit/fxmain"

	"github.com/moke-game/game/pkg/modules"
)

// Thin topology: game-only process talking to remote platform clients.
//
// Override AUTH_URL to a remote AuthService (e.g. localhost:8082).
// Do not point AUTH_URL at this process's PORT — thin does not host AuthService.
// Interactive client: --auth-host (or AUTH_URL) must match that remote auth.
func main() {
	fxmain.Main(modules.Thin)
}
