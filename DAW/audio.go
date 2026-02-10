package main

const sampleRate = 44100

func render(m model) {
	steps := 16
	stepDur := 60.0 / float64(m.bpm) / 4
	stepSamples := int(stepDur * sampleRate)

	total := steps * stepSamples
	mix := make([]float64, total)

	for i := 0; i < steps; i++ {
		pos := i * stepSamples

		if m.kick[i] {
			add(mix, kick(stepSamples), pos)
		}
		if m.snare[i] {
			add(mix, snare(stepSamples), pos)
		}
		if m.hat[i] {
			add(mix, hihat(stepSamples), pos)
		}
	}

	writeWav("output.wav", mix)
}

func add(dst, src []float64, pos int) {
	for i := 0; i < len(src) && pos+i < len(dst); i++ {
		dst[pos+i] += src[i]
	}
}
