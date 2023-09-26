package dshuf

import (
	"bytes"
	"github.com/drand/drand/chain"
	"github.com/drand/drand/client"
	"github.com/drand/drand/client/http"
)

const quicknetChainInfoJson = `{"public_key":"83cf0f2896adee7eb8b5f01fcad3912212c437e0073e911fb90022d3e760183c8c4b450b6a0a6c3ac6a5776a2d1064510d1fec758c921cc22b0e17e63aaf4bcb5ed66304de9cf809bd274ca73bab4af5a6e9c76a4bc09e76eae8991ef5ece45a","period":3,"genesis_time":1692803367,"hash":"52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971","groupHash":"f477d5c89f21a17c863a7f937c6a6d15859414d2be09cd448d4279af331c5d3e","schemeID":"bls-unchained-g1-rfc9380","metadata":{"beaconID":"quicknet"}}`

var urls = []string{
	"https://api.drand.sh",
	"https://drand.cloudflare.com",
}

func EmptyClient() (client.Client, error) {
	chainInfo, err := chain.InfoFromJSON(bytes.NewBufferString(quicknetChainInfoJson))
	if err != nil {
		return nil, err
	}
	return client.EmptyClientWithInfo(chainInfo), nil
}

func DefaultClient() (client.Client, error) {
	chainInfo, err := chain.InfoFromJSON(bytes.NewBufferString(quicknetChainInfoJson))
	if err != nil {
		return nil, err
	}
	return makeClient(chainInfo)
}

func makeClient(chainInfo *chain.Info) (client.Client, error) {
	chainHashBytes := chainInfo.Hash()

	return client.New(
		client.From(http.ForURLs(urls, chainHashBytes)...),
		client.WithChainHash(chainHashBytes),
	)
}
