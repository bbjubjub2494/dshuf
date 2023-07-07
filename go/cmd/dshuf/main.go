package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/bbjubjub2494/dshuf/go/dshuf"
	"io"
	"os"
	"time"
)

func fetchRandomness(beacon *uint64) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := dshuf.DefaultClient()
	if err != nil {
		return nil, err
	}

	if beacon == nil {
		info, err := client.Info(ctx)
		if err != nil {
			return nil, err
		}
		rn := client.RoundAt(time.Now().Add(info.Period))
		fmt.Fprintf(os.Stderr, "beacon: %d\n", rn)
		for rep := range client.Watch(ctx) {
			if rep.Round() == rn {
				return rep.Randomness(), nil
			}
		}
		return nil, fmt.Errorf("failed to watch Drand!")
	} else {
		rep, err := client.Get(ctx, *beacon)
		if err != nil {
			return nil, err
		}
		return rep.Randomness(), nil
	}
}

var cli struct {
	HeadCount *int    `short:"n" long:"head-count" placeholder:"COUNT" help:"output at most this many lines"`
	Beacon    *uint64 `short:"b" required:"" help:"round number of beacon to use for randomness"`
	File      string  `arg:"" optional:"" type:"existingfile"`
}

func main() {
	ctx := kong.Parse(&cli)

	var err error
	inputFile := os.Stdin
	if cli.File != "" && cli.File != "-" {
		inputFile, err = os.Open(cli.File)
		ctx.FatalIfErrorf(err)
	}

	randomness, err := fetchRandomness(cli.Beacon)
	ctx.FatalIfErrorf(err)

	input, err := io.ReadAll(inputFile)
	ctx.FatalIfErrorf(err)
	separator := []byte("\n")
	entries := bytes.Split(input, separator)
	if len(entries[len(entries)-1]) == 0 {
		entries = entries[:len(entries)-1]
	}
	n := len(entries)
	if cli.HeadCount != nil {
		n = *cli.HeadCount
	}
	dshuf.ShuffleInplace(randomness, &entries, n)
	for _, e := range entries {
		os.Stdout.Write(e)
		os.Stdout.Write(separator)
	}
}
