========
netty
========

A GO library for common network tools, such as ip/iptables/tc.

ip
---

Get interfaces ::

    ipCmd := &lib.IpCmd{
        Namespace: "ns0",
    }

    ifNames, e := ipCmd.GetInterfacesName()
    //ifNames is [lo enp0s3 enp0s8 enp0s9 enp0s10]

Get interface detail ::

    ipCmd := &lib.IpCmd{
        Namespace: "ns0",
    }

    details, e := ipCmd.GetInterfaceDetails("enp0s10")
    //details is map[mtu:1500 state:UP name:enp0s10]
