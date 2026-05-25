package bgld

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetBlockheaderDecodesBitgesellCoreResponse(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"result":{"hash":"0000000000000000000000000000000000000000000000000000000000000001","confirmations":2,"height":1,"version":536870912,"versionHex":"20000000","merkleroot":"4d5e6f","time":1700000000,"mediantime":1700000000,"nonce":42,"bits":"1d00ffff","difficulty":1,"chainwork":"0000000000000000000000000000000000000000000000000000000000000002","nTx":1,"previousblockhash":"0000000000000000000000000000000000000000000000000000000000000000","nextblockhash":"0000000000000000000000000000000000000000000000000000000000000002"},"error":null,"id":1}`)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	parts := strings.Split(ts.URL, ":")
	port, err := strconv.Atoi(parts[2])
	if err != nil {
		t.Fatal(err)
	}
	client, err := New(parts[1][2:], port, "user", "pass", false)
	if err != nil {
		t.Fatal(err)
	}

	header, err := client.GetBlockheader("0000000000000000000000000000000000000000000000000000000000000001")
	if err != nil {
		t.Fatalf("GetBlockheader returned error: %v", err)
	}
	if header.Bits != "1d00ffff" {
		t.Fatalf("Bits = %q, want %q", header.Bits, "1d00ffff")
	}
	if header.Previousblockhash != "0000000000000000000000000000000000000000000000000000000000000000" {
		t.Fatalf("Previousblockhash was not decoded: %q", header.Previousblockhash)
	}
	if header.Nextblockhash != "0000000000000000000000000000000000000000000000000000000000000002" {
		t.Fatalf("Nextblockhash was not decoded: %q", header.Nextblockhash)
	}
}
