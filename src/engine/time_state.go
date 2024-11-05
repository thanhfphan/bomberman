package engine

import "time"

type TimeState struct {
	Now, Last             float64
	Delta                 float64 // in seconds
	FrameCount, FrameRate int
	FrameLast             float64
}

func NewTimeState() *TimeState {
	now := float64(time.Now().UnixMilli())
	return &TimeState{
		Now:       now,
		Last:      now,
		FrameLast: now,
	}
}

func (ts *TimeState) Update() {
	ts.Now = float64(time.Now().UnixMilli())
	ts.Delta = (ts.Now - ts.Last) / 1000.0
	ts.Last = ts.Now
	ts.FrameCount++

	if ts.Now-ts.FrameLast >= 1000.0 {
		ts.FrameRate = ts.FrameCount
		ts.FrameCount = 0
		ts.FrameLast = ts.Now
	}
}
