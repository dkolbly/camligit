// camliget is a utility to push a git repository into a camlistore.
// The raw objects are pushed directly (i.e., including the header
// part), so if the camlistore uses sha1 refs, the camlistore refs
// will match what git shows
package main

import (
	"os"
	
	"github.com/dkolbly/git"
	"github.com/dkolbly/logging"
	_ "github.com/dkolbly/logging/pretty"
)

var log = logging.New("camligit")

func main() {

	back, err := NewBackend(os.Args[2], os.Args[3])

	if err != nil {
		log.Fatal(err)
	}
	
	src, err := git.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	uploaded := 0
	total := 0
	for ptr := range src.Enumerate() {
		didHave := back.Has(ptr.String())
		if !didHave {
			item := src.Get(&ptr)
			if item == nil {
				log.Fatal("rats")
			}
			buf, err := item.Payload()
			if err != nil {
				log.Fatal(err)
			}
			err = back.Put(ptr.String(), buf)
			if err != nil {
				log.Fatal(err)
			}
			uploaded++
		}
		total++
	}
	log.Info("Uploaded %d out of a total of %d", uploaded, total)
	/*
	for _, arg := range os.Args[4:] {
		log.Info("Processing: %s", arg)
		if p, ok := git.ParsePtr(arg); ok {
			item := src.Get(&p)
			if item == nil {
				log.Fatal("rats")
			}
			buf, err := item.Payload()
			if err != nil {
				log.Fatal(err)
			}
			back.Put(arg, buf)
		}
	}*/
}

