package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: DIR QUERY-FILE-PATH")
	}

	config := &packages.Config{
		Mode: packages.NeedName | packages.NeedImports | packages.NeedFiles,
		Dir:  os.Args[1],
	}
	pkgs, err := packages.Load(config, fmt.Sprintf("file=%s", os.Args[2]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("packages.Load returned %d packages\n", len(pkgs))
	for _, pkg := range pkgs {
		m, err := pkg.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(m))
	}
}
