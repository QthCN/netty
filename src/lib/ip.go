package lib

import (
	"os/exec"
	"regexp"
	"strings"

	"common"
)

type IpCmd struct {
	Namespace string
}

func (ipCmd *IpCmd) getCmdPrefixArgvs() []string {
	if ipCmd.GetNamespace() != "" {
		return []string{"netns", "exec", ipCmd.GetNamespace(), "ip"}
	} else {
		return []string{}
	}
}

//netns
func (ipCmd *IpCmd) GetNamespace() string {
	return ipCmd.Namespace
}

func (ipCmd *IpCmd) SetNamespace(ns string) string {
	ipCmd.Namespace = ns
	return ipCmd.Namespace
}

func (ipCmd *IpCmd) AddNamespace(ns string) error {
	argv := []string{"netns", "add", ns}
	c := exec.Command("ip", argv...)
	_, err := c.Output()
	return err
}

func (ipCmd *IpCmd) DeleteNamespace(ns string) error {
	argv := []string{"netns", "del", ns}
	c := exec.Command("ip", argv...)
	_, err := c.Output()
	return err
}

func (ipCmd *IpCmd) ListNamespaces() ([]string, error) {
	argv := []string{"netns", "list"}
	c := exec.Command("ip", argv...)
	d, err := c.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(d), "\n"), nil
}

//link
func (ipCmd *IpCmd) ListInterfaces() ([]string, error) {
	ifName := []string{}
	argv := ipCmd.getCmdPrefixArgvs()
	argv = append(argv, "link", "show")
	c := exec.Command("ip", argv...)
	d, err := c.Output()
	if err != nil {
		return []string{}, err
	}

	output := strings.Split(string(d), "\n")

	ifNameRegexp := regexp.MustCompile(`\d: (\w+): <.*`)
	for _, line := range output {
		matched := ifNameRegexp.FindStringSubmatch(line)
		if len(matched) > 0 {
			ifName = append(ifName, matched[len(matched)-1])
		}
	}

	return ifName, nil
}

func (ipCmd *IpCmd) GetInterfaceDetails(ifName string) (map[string]string, error) {
	details := make(map[string]string)
	argv := ipCmd.getCmdPrefixArgvs()
	argv = append(argv, "link", "show", ifName)
	c := exec.Command("ip", argv...)
	d, err := c.Output()
	if err != nil {
		return nil, err
	}

	output := strings.Split(string(d), "\n")

	if len(output) < 1 {
		return nil, &common.StrError{"No such Interface:" + ifName}
	}

	ifNameRegexp := regexp.MustCompile(`\d: (\w+): <.*> mtu (\d+) .* state (\w+) mode.*`)
	matched := ifNameRegexp.FindStringSubmatch(output[0])
	details["name"] = ifName
	details["mtu"] = matched[len(matched)-2]
	details["state"] = matched[len(matched)-1]

	return details, nil
}

func (ipCmd *IpCmd) SetInterfaceState(ifName string, state string) error {
	state = strings.ToLower(state)
	if state != "up" && state != "down" {
		return &common.StrError{"state must in ('up', 'down')"}
	}

	argv := ipCmd.getCmdPrefixArgvs()
	argv = append(argv, "link", "set", "dev", ifName, state)
	c := exec.Command("ip", argv...)
	_, err := c.Output()
	return err
}

func (ipCmd *IpCmd) AddVethPair(vethName string, peerName string) error {
	if vethName == peerName {
		return &common.StrError{"vethName cannot be same as peerName"}
	}
	argv := ipCmd.getCmdPrefixArgvs()
	argv = append(argv, "link", "add", vethName, "type", "veth", "peer", "name", peerName)
	c := exec.Command("ip", argv...)
	_, err := c.Output()
	return err
}

func (ipCmd *IpCmd) DeleteVethPair(vethName string) error {
	argv := ipCmd.getCmdPrefixArgvs()
	argv = append(argv, "link", "delete", vethName, "type", "veth")
	c := exec.Command("ip", argv...)
	_, err := c.Output()
	return err
}
