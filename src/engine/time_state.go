package engine

import "time"

type TimeState struct {
	now, last             float64
	delta                 float64 // in seconds
	frameCount, frameRate int
	frameLast             float64
}

func NewTimeState() *TimeState {
	now := float64(time.Now().UnixMilli())
	return &TimeState{
		now:       now,
		last:      now,
		frameLast: now,
	}
}

func (ts *TimeState) Update() {
	ts.now = float64(time.Now().UnixMilli())
	ts.delta = (ts.now - ts.last) / 1000.0
	ts.last = ts.now
	ts.frameCount++

	if ts.now-ts.frameLast >= 1000.0 {
		ts.frameRate = ts.frameCount
		ts.frameCount = 0
		ts.frameLast = ts.now
	}
}

func (ts *TimeState) Delta() float64 {
	return ts.delta
}

func (ts *TimeState) FrameRate() int {
	return ts.frameRate
}

func (ts *TimeState) Now() float64 {
	return ts.now
}

func (ts *TimeState) Last() float64 {
	return ts.last
}

func (ts *TimeState) FrameCount() int {
	return ts.frameCount
}

func (ts *TimeState) FrameLast() float64 {
	return ts.frameLast
}
