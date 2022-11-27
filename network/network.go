package network

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Layer struct {
	Nodes  []float64
	Biases []float64
}

type Connection struct {
	In      *Layer
	Weights [][]uint8
	Out     *Layer
}

type Network struct {
	Name string
	Cons []Connection
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func randWeights(n int) []uint8 {
	floats := randFloats(0, 1, n)
	vals := make([]uint8, n)
	for i := range vals {
		vals[i] = uint8(floats[i] * 255)
	}

	return vals
}

// Creates an new Network with random values.
func RandomNetwork(in, out, min_layers, max_layers, min_nodes, max_nodes int) Network {
	layers := make([]Layer, rand.Intn(max_layers)+min_layers)
	for i := range layers {
		length := rand.Intn(max_nodes) + min_nodes
		if i == 0 {
			length = in
		}
		if i == len(layers)-1 {
			length = out
		}

		layers[i].Nodes = randFloats(0, 1, length)
		layers[i].Biases = randFloats(0, 0.1, length)
	}

	cons := make([]Connection, len(layers)-1)
	for i := range cons {
		cons[i].In = &layers[i]
		cons[i].Out = &layers[i+1]

		cons[i].Weights = make([][]uint8, len(cons[i].Out.Nodes))
		for j := range cons[i].Out.Nodes {
			cons[i].Weights[j] = randWeights(len(cons[i].In.Nodes))
		}
	}

	return Network{
		Name: "Test network",
		Cons: cons,
	}
}

func (net Network) Proc(in []float64) []float64 {
	in_layer := Layer{
		Nodes: in,
	}
	net.Cons[0].In = &in_layer
	for _, con := range net.Cons {
		con.Out.calculate_result(*con.In, con.Weights)
	}

	return net.Cons[len(net.Cons)-1].Out.Nodes
}

func (out *Layer) calculate_result(in Layer, weights [][]uint8) {
	for i := range out.Nodes {
		count := 0
		total := 0.0
		for j := range in.Nodes {
			total = total + (in.Nodes[j] * float64(weights[i][j]))
			count = count + int(weights[i][j])
		}

		out.Nodes[i] = total/float64(count) + out.Biases[i]
	}
}

func (net Network) ToJSON() {
	res, err := json.Marshal(net)
	if err != nil {
		fmt.Printf("Couldn't marshal network: %+v", err)
	}
	fmt.Printf("%s\n", res)
}
