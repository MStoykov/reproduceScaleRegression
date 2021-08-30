# Regression found in k6 between go1.16.7 and 1.17 introduced in https://github.com/golang/go/commit/adb467ffd2d82b796de12bdd8effa2cfefe01f29

From the output of the `go test -gcflags "-m"` the difference is that ExecutionTuple.ScaleInt64 gets inlined after the commit in question which reduced the performance by *quite* a lot.
Adding `//go:noinline`  in front reverses the regression, confirming that this is the problem, but why?

## Benchstat results

```
$ benchstat 1.16.7.bench 27684ea195641ead8a8f08cb345925da889a12ed.bench adb467ffd2d82b796de12bdd8effa2cfefe01f29.bench 1.17.0.bench
name \ time/op                          1.16.7.bench  27684ea195641ead8a8f08cb345925da889a12ed.bench  adb467ffd2d82b796de12bdd8effa2cfefe01f29.bench  1.17.0.bench
ExecutionSegmentScale/5523-8             5.77ns ± 4%                                     6.20ns ± 6%                                    10.51ns ± 2%  10.45ns ± 2%
ExecutionSegmentScale/67280421310721-8   10.9ns ± 5%                                     11.4ns ± 6%                                     58.9ns ± 2%   59.9ns ± 7%
ExecutionSegmentScaleEmpty/5523-8        2.84ns ± 5%                                     2.94ns ± 5%                                     0.48ns ± 5%   0.64ns ± 7%
```
