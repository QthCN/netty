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
	//e := ipCmd.AddVethPair("v0", "v1")
	//e := ipCmd.AddNamespace("blue")
	e := ipCmd.DeleteDefaultGateway("10.0.2.1", "", "enp0s8")
	fmt.Print(e)
	//e := ipCmd.DeleteVethPair("v0")
	//fmt.Print(e)
}
