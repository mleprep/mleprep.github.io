# Yaegi (vendored, patched)

A copy of [github.com/traefik/yaegi](https://github.com/traefik/yaegi) v0.16.3 (Apache License 2.0, see LICENSE),
limited to the packages the judge builds: `interp`, `internal/unsafe2` and `stdlib/generic` (tests left out).
`../go.mod` points the module path here with a `replace` directive.

Changes from upstream (search for "Patched"):

- `interp/cfg.go`: a multiple assignment whose right side includes a call, channel receive or composite literal
  (`k, res := len(p), []int{}`) skipped the assignment of the other pairs.
- `interp/run.go`: a multiple assignment with an untyped `nil` on the right (`cur.Right, cur.Left = cur.Left, nil`)
  panicked with "reflect: New(nil)".
- `interp/type.go`: typing `*p` (dereference of a pointer value) as a pointer type crashed the compiler on
  composite literals such as `[]int{(*h)[0].v, hi}`.
- `interp/cfg.go`: `for i := 0; ; { ... }` (init statement, no condition, no post statement) re-ran the init
  statement on every iteration.

To upgrade: copy the same directories from the new release, re-apply the patches, then run the judge self-check
(`go run .` in `tools/gojudge`) and rebuild with `sh tools/gojudge/build.sh`.
