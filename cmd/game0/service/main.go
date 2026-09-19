package main

import (
	"github.com/gstones/moke-kit/fxmain"

	"github.com/moke-game/game/pkg/modules"
)

// Aggregate topology: game + in-process platform (local/dev).
// fxmain.Main already includes settings, logging, server, Mongo, Redis client, MQ router.
// AUTH_URL should match this process PORT (see .env.example).
func main() {
	fxmain.Main(modules.Aggregate)
}
