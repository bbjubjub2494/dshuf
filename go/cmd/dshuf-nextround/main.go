package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/bbjubjub2494/dshuf/go/dshuf"
	"time"
)

func roundAt(delay time.Duration) (uint64, error) {
	client, err := dshuf.EmptyClient()
	if err != nil {
		return 0, err
	}
	return client.RoundAt(time.Now().Add(delay)), nil
}

var cli struct {
	Delay time.Duration `short:"d"`
}

func main() {
	ctx := kong.Parse(&cli)

	rn, err := roundAt(cli.Delay)
	ctx.FatalIfErrorf(err)

	fmt.Println(rn)
}
