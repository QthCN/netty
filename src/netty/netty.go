package main

import (
	"fmt"

	"lib"
)

func main() {
	ipCmd := &lib.IpCmd{
		Namespace: "ns0",
	}

	fmt.Printf(ipCmd.GetNamespace())
	ipCmd.SetNamespace("aa")
	fmt.Printf(ipCmd.GetNamespace())
}
