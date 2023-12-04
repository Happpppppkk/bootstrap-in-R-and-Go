package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/montanaflynn/stats"
)

// need to create a helper function to read the csv file
func readCSVFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []float64
	for _, record := range records {
		for _, value := range record {
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			data = append(data, floatValue)
		}
	}
	return data, nil
}

// input data, lenght of data, sample size, random number
func BS(data []float64, numSamples int, numDatasets int, rnd *rand.Rand) [][]float64 {
	bsDatasets := make([][]float64, numDatasets)
	for i := 0; i < numDatasets; i++ {
		resampledData := make([]float64, numSamples)
		for j := 0; j < numSamples; j++ {
			index := rnd.Intn(len(data))
			resampledData[j] = data[index]
		}
		bsDatasets[i] = resampledData
	}
	return bsDatasets
}

func main() {
	data, err := readCSVFile("data1.csv")
	if err != nil {
		log.Println("Error reading CSV file:", err)
		return
	}

	rd := rand.New(rand.NewSource(465))
	sampleSize := 100000

	startTime := time.Now()

	var mStart, mEnd runtime.MemStats
	runtime.ReadMemStats(&mStart)

	bootstrapSamples := BS(data, len(data), sampleSize, rd)

	// compute means*combine with bootstrap mb
	bootstrapMeans := make([]float64, 0, sampleSize)
	for _, sample := range bootstrapSamples {
		mean, err := stats.Mean(sample)
		if err != nil {
			log.Println("Error computing mean:", err)
			continue
		}
		bootstrapMeans = append(bootstrapMeans, mean)

	}

	// Compute confidence interval
	lp, _ := stats.Percentile(bootstrapMeans, 2.5)
	up, _ := stats.Percentile(bootstrapMeans, 97.5)
	elapsedTime := time.Since(startTime)
	runtime.ReadMemStats(&mEnd)
	fmt.Printf("95%% Confidence Interval for the Mean: [%f, %f]\n", lp, up)
	fmt.Printf("Execution Time: %s\n", elapsedTime)
	fmt.Printf("Memory Usage: %d bytes\n", mEnd.Alloc-mStart.Alloc)
}
