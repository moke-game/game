# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260917090450-382af69044bc` ([#236](https://github.com/GStones/moke-kit/pull/236)) | `fxmain.Main` batteries-included; `WithoutTag`; Go 1.27 / grpc 1.83.2 |
| platform | `v0.0.0-20260917122241-a52357c654c6` ([#35](https://github.com/moke-game/platform/pull/35) on `main`) | kit #236 + shared `platformfx` assembly |

`AllModule` (and `GrpcModule` / `HttpModule`) do not embed auth. Pair:

- aggregate: `modules.Platform` (`AuthAllModule`)
- thin: `modules.PlatformClients` (`AuthMiddlewareModule`)

Or use the composed entries: `fxmain.Main(modules.Aggregate)` / `fxmain.Main(modules.Thin)`.

Tracking: https://github.com/moke-game/game/issues/18
