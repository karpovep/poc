package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ShouldReturnTrueAfterTimeout(t *testing.T) {
	// Given
	timeout := time.Millisecond * 50
	cancellableTimer := NewCancellableTimer()

	// When
	timedOut := <-cancellableTimer.After(timeout)

	// Then
	assert.Equal(t, true, timedOut, "timer should time out")
}

func Test_ShouldReturnFalseIfTimerCancelled(t *testing.T) {
	// Given
	timeout := time.Second
	cancellableTimer := NewCancellableTimer()
	go func() {
		time.Sleep(time.Millisecond * 50)
		cancellableTimer.Cancel()
	}()

	// When
	timedOut := <-cancellableTimer.After(timeout)

	// Then
	assert.Equal(t, false, timedOut, "timer should be cancelled")
}
