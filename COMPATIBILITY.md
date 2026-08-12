# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260812061322-0bee2b36f992` ([#231](https://github.com/GStones/moke-kit/pull/231)) | DocumentBase/NATS tests + create-game smoke/CI |
| platform | `v0.0.0-20260812062354-338ce51d453a` ([#31](https://github.com/moke-game/platform/pull/31) on `main`) | kit #231 tip bump |

`GrpcModule` / `HttpModule` / `AllModule` do not embed auth — pair `AuthAllModule` (aggregate) or `AuthMiddlewareModule` (thin) in `main`.

Validated by [game#27](https://github.com/moke-game/game/pull/27). This PR retargets the platform pin from the #31 tip to the merge commit on `main`.

Tracking: https://github.com/moke-game/game/issues/18
