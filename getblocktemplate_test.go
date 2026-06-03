package bgld

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestGetBlockTemplateReturnsRPCResult(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"result":"template-json","error":null,"id":1}`)
	})
	ts, host, port, err := getNewTestServer(handler)
	if err != nil {
		log.Fatalln(err)
	}
	defer ts.Close()

	bitcoindClient, _ := New(host, port, "x", "fake", false)
	template, err := bitcoindClient.GetBlockTemplate([]string{"proposal"}, "template")
	if err != nil {
		t.Fatalf("GetBlockTemplate error: %v", err)
	}
	if template != "template-json" {
		t.Fatalf("GetBlockTemplate returned %q, want RPC result", template)
	}
}
