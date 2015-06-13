package lib

import (
	"testing"
)

func TestIpCmdGetNamespace(t *testing.T) {
	ipCmd := &IpCmd{
		Namespace: "ns01",
	}

	if ipCmd.GetNamespace() != "ns01" {
		t.Errorf("IpCmd.GetNamespace failed.")
	}
}

func TestIpCmdSetNamespace(t *testing.T) {
	ipCmd := &IpCmd{
		Namespace: "ns01",
	}

	ipCmd.SetNamespace("ns02")
	if ipCmd.GetNamespace() != "ns02" {
		t.Errorf("IpCmd.SetNamespace failed.")
	}
}
