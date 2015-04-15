package main

import (
	"flag"
	"fmt"
	"os"
)

var opt struct {
	showGroups       bool
	groupId          string
	showHostsOfGroup bool
	showItemsOfHost  bool
	hostId           string
	showItemData     bool
	itemId           string
}

func parseFlags() {
	flag.BoolVar(&opt.showGroups, "groups", false, "show host groups")
	flag.BoolVar(&opt.showHostsOfGroup, "hosts", false, "show hosts of a group id")
	flag.StringVar(&opt.groupId, "gid", "", "group id")
	flag.BoolVar(&opt.showItemsOfHost, "items", false, "show items of a host id")
	flag.StringVar(&opt.hostId, "hid", "", "host id")
	flag.BoolVar(&opt.showItemData, "data", false, "show item history data")
	flag.StringVar(&opt.itemId, "iid", "", "item id")
	flag.Parse()
}

func handleCli() {
	if opt.showGroups {
		for _, group := range Groups() {
			fmt.Printf("id: %s name: %s\n", group.GroupId, group.Name)
		}
		os.Exit(0)
	}

	if opt.showHostsOfGroup {
		if opt.groupId == "" {
			panic("group id must be provided")
		}
		for _, host := range HostsByGroupId(opt.groupId) {
			fmt.Printf("id: %s name: %s status: %+v\n", host.HostId, host.Name, host.Status)
		}
		os.Exit(0)
	}

	if opt.showItemsOfHost {
		if opt.hostId == "" {
			panic("host id must be provided")
		}
		for _, item := range ItemsOfHost(opt.hostId) {
			fmt.Printf("id: %s name: %s key: %s\n", item.ItemId, item.Name, item.Key)
		}
		os.Exit(0)
	}

	if opt.showItemData {
		for _, data := range ItemHistory(opt.itemId, 12*6) {
			fmt.Printf("clock: %d value: %d\n", data.clock, data.value)
		}
		os.Exit(0)
	}
}
