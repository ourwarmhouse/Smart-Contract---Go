package webrtc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRTCIceGatheringState(t *testing.T) {
	testCases := []struct {
		stateString   string
		expectedState RTCIceGatheringState
	}{
		{"unknown", RTCIceGatheringState(Unknown)},
		{"new", RTCIceGatheringStateNew},
		{"gathering", RTCIceGatheringStateGathering},
		{"complete", RTCIceGatheringStateComplete},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			NewRTCIceGatheringState(testCase.stateString),
			testCase.expectedState,
			"testCase: %d %v", i, testCase,
		)
	}
}

func TestRTCIceGatheringState_String(t *testing.T) {
	testCases := []struct {
		state          RTCIceGatheringState
		expectedString string
	}{
		{RTCIceGatheringState(Unknown), "unknown"},
		{RTCIceGatheringStateNew, "new"},
		{RTCIceGatheringStateGathering, "gathering"},
		{RTCIceGatheringStateComplete, "complete"},
	}

	for i, testCase := range testCases {
		assert.Equal(t,
			testCase.state.String(),
			testCase.expectedString,
			"testCase: %d %v", i, testCase,
		)
	}
}
