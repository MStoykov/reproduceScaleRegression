package main

// ExecutionSegmentSequenceWrapper ...
type ExecutionSegmentSequenceWrapper struct {
	lcd     int64
	offsets [][]int64

	ExecutionSegmentSequence []int64 // not really
}

// ScaleInt64 ...
func (essw *ExecutionSegmentSequenceWrapper) ScaleInt64(segmentIndex int, value int64) int64 {
	start := essw.offsets[segmentIndex][0]
	offsets := essw.offsets[segmentIndex][1:]
	result := (value / essw.lcd) * int64(len(offsets))
	for gi, i := 0, start; i < value%essw.lcd; gi, i = gi+1, i+offsets[gi] {
		result++
	}
	return result
}

// ExecutionTuple ...
type ExecutionTuple struct {
	Sequence     *ExecutionSegmentSequenceWrapper
	SegmentIndex int
}

//

// ScaleInt64 scales the provided value for our execution segment.
func (et *ExecutionTuple) ScaleInt64(value int64) int64 {
	if len(et.Sequence.ExecutionSegmentSequence) == 1 {
		return value // if we don't have any segmentation, just return the original value
	}
	return et.Sequence.ScaleInt64(et.SegmentIndex, value)
}
