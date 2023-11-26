package integration

import (
	"net"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, nil))
}

const chainHash = "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971"

func mockDrandEndpoint() (endpoint string, err error) {
	http.HandleFunc("/"+chainHash+"/info", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"public_key":"83cf0f2896adee7eb8b5f01fcad3912212c437e0073e911fb90022d3e760183c8c4b450b6a0a6c3ac6a5776a2d1064510d1fec758c921cc22b0e17e63aaf4bcb5ed66304de9cf809bd274ca73bab4af5a6e9c76a4bc09e76eae8991ef5ece45a","period":3,"genesis_time":1692803367,"hash":"52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971","groupHash":"f477d5c89f21a17c863a7f937c6a6d15859414d2be09cd448d4279af331c5d3e","schemeID":"bls-unchained-g1-rfc9380","metadata":{"beaconID":"quicknet"}}`))
	})

	http.HandleFunc("/"+chainHash+"/public/", func(w http.ResponseWriter, r *http.Request) {
		roundNumber := path.Base(r.URL.Path)
		switch roundNumber {
		case "1":
			w.Write([]byte(`{"round":1,"randomness":"1466a6cd24e327188770752f6134001c64d6efcc590ccc26b721611ad96f165a","signature":"b55e7cb2d5c613ee0b2e28d6750aabbb78c39dcc96bd9d38c2c2e12198df95571de8e8e402a0cc48871c7089a2b3af4b"}`))
		case "2":
			w.Write([]byte(`{"round":2,"randomness":"5782d6987841c654515a0e72b2d1ebb4e741234042c37cb19608ae50d93fb60c","signature":"b6b6a585449b66eb12e875b64fcbab3799861a00e4dbf092d99e969a5eac57dd3f798acf61e705fe4f093db926626807"}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	go http.Serve(listener, nil)

	return "http://" + listener.Addr().String(), nil
}
func TestScript(t *testing.T) {
	addr, err := mockDrandEndpoint()
	if err != nil {
		t.Fatal(err)
	}
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
		Setup: func(env *testscript.Env) error {
			env.Setenv("DSHUF_ENDPOINT", addr)
			return nil
		},
	})
}
