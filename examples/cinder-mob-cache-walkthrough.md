# Cinder Mob Cache Core Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 166 | ship |
| stress | sync drift | 126 | watch |
| edge | local state | 147 | ship |
| recovery | conflict cost | 101 | hold |
| stale | form pressure | 200 | ship |

Start with `stale` and `recovery`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `form pressure` against `conflict cost`, not the raw score alone.
