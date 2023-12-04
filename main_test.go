package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BS function definition remains the same

func TestBS(t *testing.T) {
	assert := assert.New(t)

	data := []float64{1, 2, 3, 4, 5}
	numSamples := 3
	numDatasets := 4
	rnd := rand.New(rand.NewSource(0))

	result := BS(data, numSamples, numDatasets, rnd)

	assert.Equal(numDatasets, len(result), "Number of datasets should match")

	for i, dataset := range result {
		assert.Equal(numSamples, len(dataset), "Dataset %d should have the correct number of samples", i)
	}
}
