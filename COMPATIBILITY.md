# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260918072054-388d523a53ef` ([#238](https://github.com/GStones/moke-kit/pull/238)) | kit #236 assembly + #238 module bump; `fxmain.Main` batteries-included; `WithoutTag`; Go 1.27 / grpc 1.83.2 / gorm 1.25.12 |
| platform | `v0.0.0-20260920031848-1c84eb3aae26` ([#37](https://github.com/moke-game/platform/pull/37) on `main`) | kit #238 + shared `platformfx` assembly |

`AllModule` (alias of `HttpModule`) and `GrpcModule` do not embed auth. Pair:

- aggregate: `modules.Platform` (`AuthAllModule`)
- thin: `modules.PlatformClients` (`AuthMiddlewareModule`)

Or use the composed entries: `fxmain.Main(modules.Aggregate)` / `fxmain.Main(modules.Thin)`.

Game settings/clients reuse `platformfx.ProvideFromEnv` / `platformfx.NewClient` instead of a local copy.

Pins (not latest-on-purpose):

- **`google.golang.org/grpc` v1.83.2** — patched 1.83 line. `v1.84.0` reintroduces [GO-2026-6443](https://pkg.go.dev/vuln/GO-2026-6443) until 1.85.
- **`gorm.io/gorm` v1.25.12** — stay on 1.25.x. `v1.31.x` `Raw().Scan()` zeros unselected fields ([go-gorm/gorm#7746](https://github.com/go-gorm/gorm/issues/7746)).
- **`github.com/gstones/zinx`** stays on `v1.2.7-0.20240617071724-88bd884d8d08` (HEAD; tagged `@latest` would downgrade to v1.2.6).

Tracking: https://github.com/moke-game/game/issues/18
