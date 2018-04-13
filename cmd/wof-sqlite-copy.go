package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/writer"
	"log"
)

func main() {

	var source = flag.String("source", "", "...")
	var target = flag.String("target", "", "...")

	flag.Parse()

	r, err := reader.NewSQLiteReader(*source)

	if err != nil {
		log.Fatal(err)
	}

	w, err := writer.NewSQLiteWriter(*target)

	if err != nil {
		log.Fatal(err)
	}

	for _, path := range flag.Args() {

		fh, err := r.Read(path)

		if err != nil {
			log.Fatal(err)
		}

		err = w.Write(path, fh)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("copied %s\n", path)
	}

}
