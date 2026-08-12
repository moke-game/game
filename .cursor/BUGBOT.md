# Bugbot review rules (game)

## Auth topology
- Public game APIs must not embed `utility.WithoutAuth`.
- `GrpcModule` / `HttpModule` / `AllModule` have no auth provider — pair `AuthAllModule` (aggregate) or `AuthMiddlewareModule` (thin) in `main`, never both with a stub.
- Flag handlers that trust `request.uid` instead of `UIDContextKey`.
- MatchFunction / Open Match callbacks may be WithoutAuth — flag if that type is reused for player-facing RPCs.

## TCP
- zinx/TCP is not covered by gRPC AuthMiddleware; flag enabling `AllWithTCPModule` as a public default.

## Deps
- Flag `go.mod` Go version drift vs Docker builder image.
- Flag pins to unmerged platform/kit commits without a PR dependency note.
