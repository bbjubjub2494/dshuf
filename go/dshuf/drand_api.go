package dshuf

import (
	"encoding/hex"
	"github.com/drand/drand/client"
	"github.com/drand/drand/client/http"
)

const defaultChainHash = "8990e7a9aaed2ffed73dbd7092123d6f289930540d7651336225dc172e51b2ce"

var urls = []string{
	"https://api.drand.sh",
	"https://drand.cloudflare.com",
}

func DefaultClient() (client.Client, error) {
	return makeClient(defaultChainHash)
}

func makeClient(chainHash string) (client.Client, error) {
	chainHashBytes, err := hex.DecodeString(chainHash)
	if err != nil {
		return nil, err
	}

	return client.New(
		client.From(http.ForURLs(urls, chainHashBytes)...),
		client.WithChainHash(chainHashBytes),
	)
}
