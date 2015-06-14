package lib

import (
	"os/exec"
	"regexp"
	"strings"
)

type IpCmd struct {
	Namespace string
}

func (ipCmd *IpCmd) GetNamespace() string {
	return ipCmd.Namespace
}

func (ipCmd *IpCmd) SetNamespace(ns string) string {
	ipCmd.Namespace = ns
	return ipCmd.Namespace
}

func (ipCmd *IpCmd) getCmdPrefixArgvs() []string {
	if ipCmd.GetNamespace() != "" {
		return []string{"netns", "exec", ipCmd.GetNamespace()}
	} else {
		return []string{}
	}
}

//ip link show
func (ipCmd *IpCmd) GetInterfacesName() ([]string, error) {
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
