# Compatibility

| Dependency | Version | Notes |
|------------|---------|-------|
| moke-kit | `v1.0.5-0.20260811094419-bcdfe55515cd` (#224); tip [#228](https://github.com/GStones/moke-kit/pull/228) on kit `main` | Binder fail-closed; StopServing/CAS/CI on tip |
| platform | `v0.0.0-20260812022153-40241b2322e3` ([#27](https://github.com/moke-game/platform/pull/27) on `main`) | jwt/v5, CAS/chat lifecycle tests, CI lint/vuln |

Validated by [game#22](https://github.com/moke-game/game/pull/22). This PR retargets the platform pin from the #27 tip to the merge commit on `main`.

Tracking: https://github.com/moke-game/game/issues/18
