package main

import (
	"fmt"
	"testing"
)

func BenchmarkExecutionSegmentScale(b *testing.B) {
	for _, value := range []int64{5523, 67280421310721} {
		value := value
		et := &ExecutionTuple{
			SegmentIndex: 0,
			// not real values but doesn't matter
			Sequence: &ExecutionSegmentSequenceWrapper{
				lcd:     23,
				offsets: [][]int64{[]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			},
		}
		b.Run(fmt.Sprintf("%d", value), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				et.ScaleInt64(value)
			}
		})
	}
}

func BenchmarkExecutionSegmentScaleEmpty(b *testing.B) {
	for _, value := range []int64{5523} {
		value := value
		et := &ExecutionTuple{
			SegmentIndex: 0,
			// not real values but doesn't matter
			Sequence: &ExecutionSegmentSequenceWrapper{
				ExecutionSegmentSequence: []int64{1},
			},
		}
		b.Run(fmt.Sprintf("%d", value), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				et.ScaleInt64(value)
			}
		})
	}
}
