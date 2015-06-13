package lib

import (
	"log"
)

type IpCmd struct {
	Namespace string
	*log.Logger
}

func (ipCmd *IpCmd) GetNamespace() string {
	return ipCmd.Namespace
}

func (ipCmd *IpCmd) SetNamespace(ns string) string {
	ipCmd.Namespace = ns
	return ipCmd.Namespace
}
