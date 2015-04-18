package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cliLoop() {
	fmt.Println("zabbix> Press 'h' for help")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("zabbix> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text = text[:len(text)-1] // strip EOL
		if text == "" {
			continue
		}

		if !handleCliCmd(text) {
			// its internal command
			break
		}
	}

}

func handleCliCmd(s string) bool {
	switch {
	case s == "help" || s == "h":
		fmt.Println("top groups hosts items setitem stat addfavorite favorites go bye")

	case s == "q" || s == "bye" || s == "quit":
		return false

	case s == "go":
		drawDashboard()

	case s == "status" || s == "stat" || s == "s":
		// TODO show name instead of id
		fmt.Printf("group: %s, item: %s\n",
			group, itemName)

	case s == "add" || s == "af":
		addFavorite(itemName)

	case s == "favorites" || s == "f":
		for f, _ := range favoriteItems {
			fmt.Println(f)
		}

	case s == "top":
		hosts := make(map[string]struct{})
		for _, group := range Groups() {
			fmt.Printf("%s\n", group.Name)
			gid := group.GroupId
			for _, host := range HostsByGroupId(gid) {
				fmt.Printf("\t%s\n", host.Name)
				hosts[host.Name] = struct{}{}
			}
		}
		fmt.Printf("-- Total: %d hosts\n", len(hosts))

	case s == "groups" || s == "g":
		for _, group := range Groups() {
			fmt.Printf("id:%s name:%s\n", group.GroupId, group.Name)
		}

	case strings.HasPrefix(s, "hosts"):
		args := strings.SplitN(s, " ", 2)
		var gid string
		if len(args) != 2 {
			if cliGroupId == "" {
				fmt.Println("usage: hosts <group id>")
				return true
			} else {
				gid = cliGroupId
			}
		} else {
			gid = args[1]
			cliGroupId = gid
		}

		group = GroupName(gid)

		for _, host := range HostsByGroupId(gid) {
			fmt.Printf("id:%s name:%s status:%+v\n", host.HostId, host.Name, host.Status)
		}

	case strings.HasPrefix(s, "items"):
		args := strings.SplitN(s, " ", 2)
		var hostId string
		if len(args) != 2 {
			if cliHostId == "" {
				hostId = HostsByGroupId(cliGroupId)[0].HostId
				cliHostId = hostId
			} else {
				hostId = cliHostId
			}
		} else {
			hostId = args[1]
			cliHostId = hostId
		}
		for _, item := range ItemsOfHost(hostId) {
			fmt.Printf("id:%s name:%s key:%s\n", item.ItemId, item.Name, item.Key)
		}

	case strings.HasPrefix(s, "setitem"):
		args := strings.SplitN(s, " ", 2)
		if len(args) != 2 {
			fmt.Println("usage: setitem <item name>")
			return true
		}
		itemName = args[1]
		fmt.Println(itemName)

	default:
		fmt.Printf("unkown command: %s\n", s)

	}

	return true
}
