# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260812061322-0bee2b36f992` ([#231](https://github.com/GStones/moke-kit/pull/231)) | DocumentBase/NATS tests + create-game smoke/CI |
| platform | tip `196750cd…` (kit #231 bump) | Depends on [platform kit#231 PR](https://github.com/moke-game/platform/compare/main...cursor/post-merge-kit231-3293); retarget to merge after land |

`GrpcModule` / `HttpModule` / `AllModule` do not embed auth — pair `AuthAllModule` (aggregate) or `AuthMiddlewareModule` (thin) in `main`.

Tracking: https://github.com/moke-game/game/issues/18
