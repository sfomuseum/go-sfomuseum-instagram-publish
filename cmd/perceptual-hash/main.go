// perceptual-hash is a command line utility for generating perceptual hashes for one or more files.
// The expected use of the tool is for checking the hashes of existing SFO Museum Instagram images
// and those are included with Instagram export bundles. For example:
//
//	> go run -mod vendor cmd/perceptual-hash/main.go ~/Desktop/120885293_2012351615564945_4864065299023451274_n_17956147654363845.jpg ~/Desktop/3b1bce024e1f35517a8d517a2a8cd169_77e79d2a1a_o.jpg
//	p:b867679231ccc633
//	p:b867679231ccc633
package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/sfomuseum/go-sfomuseum-instagram/hash"
)

func main() {

	flag.Parse()

	for _, uri := range flag.Args() {

		// To do: Some sort of syntax to allow us to use URIs that
		// encode both gocloud/blob.Bucket specifics and the file
		// to fetch in a single string

		r, err := os.Open(uri)

		if err != nil {
			log.Fatalf("Failed to open %s, %v", uri, err)
		}

		h, err := hash.PerceptualHash(r)

		if err != nil {
			log.Fatalf("Failed to derive hash for %s, %v", uri, err)
		}

		fmt.Println(h)
	}
}
