# Regression found in k6 between go1.16.7 and 1.17 introduced in https://github.com/golang/go/commit/adb467ffd2d82b796de12bdd8effa2cfefe01f29

From the output of the `go test -gcflags "-m"` the difference is that ExecutionTuple.ScaleInt64 gets inlined after the commit in question which reduced the performance by *quite* a lot.
Adding `//go:noinline`  in front reverses the regression, confirming that this is the problem, but why?
