package dshuf

import (
	"github.com/zeebo/blake3"

	"encoding/binary"
	"math/big"
)

const SAMPLE_LEN = 24

type prng blake3.Digest

func makePRNG(randomness []byte, es [][]byte) *prng {
	h, _ := blake3.NewKeyed(randomness)
	for _, e := range es {
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(len(e)))
		h.Write(buf[:])
		h.Write(e)
	}
	return (*prng)(h.Digest())
}

func (self *prng) nextBigInt() *big.Int {
	var sample [SAMPLE_LEN]byte
	(*blake3.Digest)(self).Read(sample[:])
	return new(big.Int).SetBytes(sample[:])
}

func ShuffleInplace(
	randomness []byte,
	input *[][]byte,
	limit int,
) {
	es := *input
	if limit > len(es) {
		limit = len(es)
	}

	prng := makePRNG(randomness, es)
	for i := 0; i < limit; i++ {
		r := prng.nextBigInt()
		r.Mod(r, big.NewInt(int64(len(es)-i)))
		j := i + int(r.Uint64())
		es[i], es[j] = es[j], es[i]
	}
	*input = es[:limit]
}

func ShuffleWithReplacement(
	randomness []byte,
	input [][]byte,
) <-chan []byte {
	prng := makePRNG(randomness, input)
	out := make(chan []byte)

	go func() {
		for {
			r := prng.nextBigInt()
			r.Mod(r, big.NewInt(int64(len(input))))
			j := int(r.Uint64())
			out <- input[j]
		}
	}()

	return out
}
