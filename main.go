package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Gewad/GoML/mnist"
	"github.com/Gewad/GoML/network"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	net := network.RandomNetwork(784, 10, 2, 5, 10, 30)

	res, err := mnist.ReadTestSet(".\\.test_data\\mnist")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	first := res.Data[266]
	data := make([]float64, 784)
	for i := range first.Image {
		for j := range first.Image[i] {
			data[28*i+j] = float64(first.Image[i][j] / 255)
		}
	}

	//net.ToJSON()
	result := net.Proc(data)
	fmt.Printf("%+v\n", result)

	chosen := maxIndexFromSlice(result)
	if first.Digit == chosen {
		fmt.Printf("The neural network chose the right value!\n")
	} else {
		fmt.Printf("The neural network was wrong: chose %d instead of %d\n", chosen, first.Digit)
	}
	fmt.Printf("Confidence: %f", result[chosen])
}

func maxIndexFromSlice(sl []float64) int {
	maxIndex := 0
	max := 0.0

	for i, val := range sl {
		if val > max {
			max = val
			maxIndex = i
		}
	}

	return maxIndex
}
