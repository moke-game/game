# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260812022140-acb9f313d7fd` ([#228](https://github.com/GStones/moke-kit/pull/228)) | Binder fail-closed + StopServing/CAS/CI |
| platform | tip `1753fc96…` (kit #228 bump + prod JWT expire guard) | Depends on [platform PR](https://github.com/moke-game/platform/compare/main...cursor/issue-23-kit228-jwt-buf-3293) |

`GrpcModule` / `HttpModule` / `AllModule` do not embed auth — pair `AuthAllModule` (aggregate) or `AuthMiddlewareModule` (thin) in `main`.

Tracking: https://github.com/moke-game/game/issues/18
