package main

import "math"

func kick(n int) []float64 {
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		t := float64(i) / sampleRate
		f := 120 * math.Exp(-t*8)
		out[i] = math.Sin(2*math.Pi*f*t) * math.Exp(-t*10)
	}
	return out
}

func snare(n int) []float64 {
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		t := float64(i) / sampleRate
		out[i] = (rand()*2-1) * math.Exp(-t*20) * 0.5
	}
	return out
}

func hihat(n int) []float64 {
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		t := float64(i) / sampleRate
		out[i] = (rand()*2-1) * math.Exp(-t*50) * 0.3
	}
	return out
}

var seed uint32 = 1

func rand() float64 {
	seed = seed*1664525 + 1013904223
	return float64(seed%1000) / 1000
}
