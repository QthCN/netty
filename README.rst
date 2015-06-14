========
netty
========

A GO library for common network tools, such as ip/iptables/tc.

ip
---

ip link show ::
    ipCmd := &lib.IpCmd{
        Namespace: "ns0",
    }

    ifNames, e := ipCmd.GetInterfacesName()
    //ifNames is [lo enp0s3 enp0s8 enp0s9 enp0s10]
