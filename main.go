package main

import (
	"flag"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Generates Terraform-style IDs. Arguments, if provided, will be used as ID prefixes. Results are newline-separated.")
	}

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println(resource.UniqueId())
		return
	}

	for _, arg := range flag.Args() {
		fmt.Println(resource.PrefixedUniqueId(arg))
	}
}
