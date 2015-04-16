package main

import (
	"github.com/AlekSi/zabbix"
	"strconv"
)

type historyData struct {
	clock int
	value int
}

// group -> host -> item -> history data
func Groups() zabbix.HostGroups {
	groups, err := api.HostGroupsGet(zabbix.Params{})
	must(err)
	return groups
}

func GroupId(groupName string) string {
	for _, g := range Groups() {
		if g.Name == groupName {
			return g.GroupId
		}
	}

	return emptyStr
}

func GroupName(gid string) string {
	for _, g := range Groups() {
		if g.GroupId == gid {
			return g.Name
		}
	}

	return emptyStr
}

func HostsOfGroup(groupName string) zabbix.Hosts {
	hostGroups := zabbix.HostGroups{zabbix.HostGroup{GroupId: GroupId(groupName)}}
	hosts, err := api.HostsGetByHostGroups(hostGroups)
	must(err)
	return hosts
}

func HostsByGroupId(groupId string) zabbix.Hosts {
	hostGroups := zabbix.HostGroups{zabbix.HostGroup{GroupId: groupId}}
	hosts, err := api.HostsGetByHostGroups(hostGroups)
	must(err)
	return hosts
}

func ItemsOfHost(hostId string) zabbix.Items {
	items, err := api.ItemsGet(zabbix.Params{"hostids": hostId})
	must(err)
	return items
}

func ItemByName(hostId string, name string) zabbix.Item {
	items := ItemsOfHost(hostId)
	for _, item := range items {
		if name == item.Name {
			return item
		}
	}

	return zabbix.Item{}
}

func ItemById(hostId string, itemId string) zabbix.Item {
	items := ItemsOfHost(hostId)
	for _, item := range items {
		if itemId == item.ItemId {
			return item
		}
	}

	return zabbix.Item{}
}

func ItemHistory(id string, limit int) []historyData {
	data, err := api.CallWithError("history.get",
		zabbix.Params{"output": "extend", "itemids": id, "limit": limit,
			"sortfield": "clock", "sortorder": "DESC"})
	must(err)
	d := data.Result.([]interface{})
	r := make([]historyData, len(d))
	j := 0
	for i := len(d) - 1; i >= 0; i-- {
		x := d[i].(map[string]interface{})
		value, _ := strconv.Atoi(x["value"].(string))
		clock, _ := strconv.Atoi(x["clock"].(string))
		r[j] = historyData{value: value, clock: clock}
		j++
	}

	return r
}
