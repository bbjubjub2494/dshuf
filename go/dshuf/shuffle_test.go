package dshuf

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"testing"
)

// Byte slice but represented as hex in JSON
type HexBytes []byte

func (b *HexBytes) MarshalJSON() ([]byte, error) {
	s := hex.EncodeToString(*b)
	return json.Marshal(s)
}

func (b *HexBytes) UnmarshalJSON(buf []byte) error {
	var s string
	if err := json.Unmarshal(buf, &s); err != nil {
		return err
	}
	dec, err := hex.DecodeString(s)
	if err == nil {
		*b = dec
	}
	return err
}

// Helper
func unwrap(b []string) [][]byte {
	r := make([][]byte, len(b))
	for i := range b {
		r[i] = []byte(b[i])
	}
	return r
}

type TestCase struct {
	Input       []string `json:"input"`
	Randomness  HexBytes `json:"randomness"`
	Limit       int      `json:"limit"`
	Repetitions bool     `json:"repetitions"`
	Output      []string `json:"output"`
}

func loadTestCase(t *testing.T, name string) (tc TestCase) {
	file, err := os.Open("../../testcases/" + name + ".json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(data, &tc)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func checkListEqual(t *testing.T, expected, actual [][]byte) {
	if len(expected) != len(actual) {
		t.Fatalf("expected: %d got: %d", len(expected), len(actual))
	}
	for i := range actual {
		if !bytes.Equal(expected[i], actual[i]) {
			t.Fatalf("expected: %x got: %x", expected[i], actual[i])
		}
	}
}

func testCase(t *testing.T, name string) {
	tc := loadTestCase(t, name)
	input := unwrap(tc.Input)
	var output [][]byte

	if tc.Repetitions {
		output = make([][]byte, tc.Limit)
		c := ShuffleWithReplacement(tc.Randomness, input)
		for i := 0; i < tc.Limit; i++ {
			output[i] = <-c
		}
	} else {
		ShuffleInplace(tc.Randomness, &input, tc.Limit)
		output = input
	}

	checkListEqual(t, unwrap(tc.Output), output)
}

func TestCase_Basic(t *testing.T) {
	testCase(t, "basic")
}

func TestCase_BasicLess(t *testing.T) {
	testCase(t, "basic_less")
}

func TestCase_BasicMore(t *testing.T) {
	testCase(t, "basic_more")
}

func TestCase_BasicOtherInput(t *testing.T) {
	testCase(t, "basic_other_input")
}

func TestCase_BasicOtherRandomness(t *testing.T) {
	testCase(t, "basic_other_randomness")
}

func TestCase_BasicRepetitions(t *testing.T) {
	testCase(t, "basic_repetitions")
}
