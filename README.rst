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

Add veth pair ::

    ipCmd := &lib.IpCmd{
        Namespace: "",
    }

    e := ipCmd.AddVethPair("v0", "v1")
    //veth pair v0 and v1 set up now

Delete veth pair ::

    ipCmd := &lib.IpCmd{
        Namespace: "",
    }

    e := ipCmd.DeleteVethPair("v0")
    //veth v0 and it's peer are removed

Add interface into namespace ::

    ipCmd := &lib.IpCmd{
        Namespace: "",
    }

    e := ipCmd.AddInterfaceIntoNamespace("v0", "blue")
    //interface v0 now in namespace blue 

Add IP ::

    ipCmd := &lib.IpCmd{
        Namespace: "",
    }

    e := ipCmd.AddIPAddress("veth0", "10.5.5.5/24")
    //veth0 has IP address "10.5.5.5/24" now

Delete IP ::

    ipCmd := &lib.IpCmd{
        Namespace: "",
    }

    e := ipCmd.DeleteIPAddress("veth0", "10.5.5.5/24")
    //IP address "10.5.5.5/24" is removed

Get IP ::

    ipCmd := &lib.IpCmd{
        Namespace: "",
    }

    ipAddresses, e := ipCmd.GetIPAddress("veth0")
    //ipAddresses is [10.5.5.5/24 10.6.5.5/24 10.7.5.5/24] now
