package main

import (
	"encoding/binary"
	"os"
)

const (
	bitDepth = 16
	channels = 1
)

func writeWav(filename string, samples []float64) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numSamples := len(samples)
	byteRate := sampleRate * channels * bitDepth / 8
	blockAlign := channels * bitDepth / 8
	dataSize := numSamples * bitDepth / 8


	file.WriteString("RIFF")
	binary.Write(file, binary.LittleEndian, uint32(36+dataSize))
	file.WriteString("WAVE")


	file.WriteString("fmt ")
	binary.Write(file, binary.LittleEndian, uint32(16))         
	binary.Write(file, binary.LittleEndian, uint16(1))           
	binary.Write(file, binary.LittleEndian, uint16(channels))
	binary.Write(file, binary.LittleEndian, uint32(sampleRate))
	binary.Write(file, binary.LittleEndian, uint32(byteRate))
	binary.Write(file, binary.LittleEndian, uint16(blockAlign))
	binary.Write(file, binary.LittleEndian, uint16(bitDepth))


	file.WriteString("data")
	binary.Write(file, binary.LittleEndian, uint32(dataSize))


	for _, s := range samples {
		if s > 1 {
			s = 1
		}
		if s < -1 {
			s = -1
		}
		val := int16(s * 32767)
		binary.Write(file, binary.LittleEndian, val)
	}
}
