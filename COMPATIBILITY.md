# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260812022140-acb9f313d7fd` ([#228](https://github.com/GStones/moke-kit/pull/228)); create-game [#229](https://github.com/GStones/moke-kit/pull/229)/[#230](https://github.com/GStones/moke-kit/pull/230) | Binder fail-closed + StopServing/CAS/CI; scaffold thin/smoke |
| platform | tip `98db4775…` (compress + COMPAT) on [#28](https://github.com/moke-game/platform/pull/28) line | Depends on [platform compress PR](https://github.com/moke-game/platform/compare/main...cursor/issue-23-compress-compat-3293) |

`GrpcModule` / `HttpModule` / `AllModule` do not embed auth — pair `AuthAllModule` (aggregate) or `AuthMiddlewareModule` (thin) in `main`.

Tracking: https://github.com/moke-game/game/issues/18
