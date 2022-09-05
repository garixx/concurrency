package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWithRange(t *testing.T) {
	start := time.Now()
	stream := GetMockStream()

	tw := producerOption1(stream)
	consumerOption1(tw)

	duration := time.Since(start)
	fmt.Printf("Process took %s\n", duration)
	assert.True(t, duration.Seconds() < 2)
}

// TODO: review panic on channel closed
func TestWithSelect(t *testing.T) {
	start := time.Now()
	stream := GetMockStream()

	done := make(chan struct{})
	//defer close(done)

	tw := producerOption2(done, stream)
	consumerOption2(done, tw)

	duration := time.Since(start)
	fmt.Printf("Process took %s\n", duration)
	//assert.True(t, duration.Seconds() < 2)
}
