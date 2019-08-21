package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tinylib/msgp/msgp"
)

type payload struct {
	// header specifies the first few bytes in the msgpack stream
	// indicating the type of array (fixarray, array16 or array32)
	// and the number of items contained in the stream.
	header []byte

	// off specifies the current read position on the header.
	off int

	// count specifies the number of items in the stream.
	count uint64

	// buf holds the sequence of msgpack-encoded items.
	buf bytes.Buffer
}

func main() {
	http.HandleFunc("/v0.4/traces", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var pload spanLists
		msgp.Decode(r.Body, &pload)

		var topLevelSpans []*span
		var allSpans []*span
		for _, sl := range pload {
			for _, span := range sl {
				allSpans = append(allSpans, span)
			}
		}

		for _, span := range allSpans {
			if span.ParentID != 0 {
				for _, sp := range allSpans {
					if sp.SpanID == span.ParentID {
						sp.childSpans = append(sp.childSpans, span)
					}
				}
			} else {
				topLevelSpans = append(topLevelSpans, span)
			}
		}

		for _, span := range topLevelSpans {
			doPrint(span, 0)
		}
		fmt.Fprintf(w, "")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8126", nil))
}

func doPrint(sp *span, depth int) {
	if depth == 0 {
		fmt.Println("")
	}
	printStr := fmt.Sprintf("%%%ds resource: %%.30q\t|\t name %%s\t|\t duration %%s\n", depth*2)
	fmt.Printf(printStr, "", sp.Resource, sp.Name, time.Duration(sp.Duration))
	for _, chsp := range sp.childSpans {
		doPrint(chsp, depth+1)
	}
}

// Read implements io.Reader. It reads from the msgpack-encoded stream.
func (p *payload) Read(b []byte) (n int, err error) {
	if p.off < len(p.header) {
		// reading header
		n = copy(b, p.header[p.off:])
		p.off += n
		return n, nil
	}
	return p.buf.Read(b)
}
