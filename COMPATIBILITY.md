# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260812022140-acb9f313d7fd` ([#228](https://github.com/GStones/moke-kit/pull/228)); create-game tip [#229](https://github.com/GStones/moke-kit/pull/229) | Binder fail-closed + StopServing/CAS/CI; scaffold thin/auth-free modules |
| platform | `v0.0.0-20260812031638-1fe842e07c0c` ([#28](https://github.com/moke-game/platform/pull/28) on `main`) | kit #228 + prod JWT expire guard + buf-push gate |

`GrpcModule` / `HttpModule` / `AllModule` do not embed auth — pair `AuthAllModule` (aggregate) or `AuthMiddlewareModule` (thin) in `main`.

Validated by [game#24](https://github.com/moke-game/game/pull/24). This PR retargets the platform pin from the #28 tip to the merge commit on `main`.

Tracking: https://github.com/moke-game/game/issues/18
