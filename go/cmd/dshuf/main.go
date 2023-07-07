package main

import (
	"bytes"
	"context"
	"github.com/bbjubjub2494/dshuf/go/dshuf"
	"io"
	"log"
	"os"
)

func fetchRandomness(rn uint64) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := dshuf.DefaultClient()
	if err != nil {
		return nil, err
	}

	rep, err := client.Get(ctx, rn)
	if err != nil {
		return nil, err
	}

	return rep.Randomness(), nil
}

func main() {
	const rn = 1337 // TODO: no hardcoding

	randomness, err := fetchRandomness(rn)
	if err != nil {
		log.Fatal(err)
	}

	// simulate shuf -n 3
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	separator := []byte("\n")
	entries := bytes.Split(input, separator)
	if len(entries[len(entries)-1]) == 0 {
		entries = entries[:len(entries)-1]
	}
	dshuf.ShuffleInplace(randomness, &entries, 3)
	for _, e := range entries {
		os.Stdout.Write(e)
		os.Stdout.Write(separator)
	}
}
