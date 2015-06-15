========
netty
========

A GO library for common network tools, such as ip/iptables/tc.

*Notice*: This lib is not productive now. Lack tests.

ip
---

Get Namespaces ::

    ipCmd := &lib.IpCmd{}
    ns, e := ipCmd.ListNamespaces()
    //ns is [ns1 ns0 ]

Add Namespace ::

    ipCmd := &lib.IpCmd{}
    e := ipCmd.AddNamespace("ns")
    // ip netns list will get 'ns' now

Delete Namespace ::

    ipCmd := &lib.IpCmd{}
    e := ipCmd.DeleteNamespace("ns")
    // ip netns list will not get 'ns' now

Get interfaces ::

    ipCmd := &lib.IpCmd{
        Namespace: "ns0",
    }

    ifNames, e := ipCmd.ListInterfaces()
    //ifNames is [lo enp0s3 enp0s8 enp0s9 enp0s10]

Get interface detail ::

    ipCmd := &lib.IpCmd{
        Namespace: "ns0",
    }

    details, e := ipCmd.GetInterfaceDetails("enp0s10")
    //details is map[mtu:1500 state:UP name:enp0s10]

Set interface status ::

    ipCmd := &lib.IpCmd{
        Namespace: "ns0",
    }

    e := ipCmd.SetInterfaceState("enp0s3", "up")
    //enp0s3 is up
    e := ipCmd.SetInterfaceState("enp0s3", "down")
    //enp0s3 is down
