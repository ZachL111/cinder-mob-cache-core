# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 166, lane `ship`
- `stress`: `sync drift`, score 126, lane `watch`
- `edge`: `local state`, score 147, lane `ship`
- `recovery`: `conflict cost`, score 101, lane `hold`
- `stale`: `form pressure`, score 200, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
