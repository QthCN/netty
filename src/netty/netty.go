package main

import (
	"fmt"

	"lib"
)

func main() {
	ipCmd := &lib.IpCmd{}

	//fmt.Printf(ipCmd.GetNamespace())
	//ipCmd.SetNamespace("")
	//fmt.Printf(ipCmd.GetNamespace())
	//r, e := ipCmd.GetInterfacesName()
	//fmt.Print(r)
	//fmt.Print(e)
	//r, e := ipCmd.GetInterfaceDetails("lo")
	e := ipCmd.SetInterfaceState("enp0s3", "up")
	fmt.Print(e)
}
