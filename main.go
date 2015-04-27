package main

import (
	"fmt"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchData(idx int, size int) []float64 {
	if len(data[idx]) == dataSize {
		// shift rightwards
		data[idx] = data[idx][1:]
	}

	item := ItemByName(hosts[idx].HostId, itemName)
	his := ItemHistory(item.ItemId, size)
	for j := 0; j < len(his); j++ {
		data[idx] = append(data[idx], float64(his[j].value))
	}
	return data[idx]
}

func init() {
	if user == "" || pass == "" {
		fmt.Println("setenv ZABBIX_USER ZABBIX_PASS first")
		os.Exit(0)
	}

}

func showHostsOfGroup(names []string) {
	groupMapping := map[string]string{
		"lvs":  "FFan LVS",
		"ngix": "FFan Nginx",
		"gw":   "FFan Gateway",

		"java1": "FFan Java1",
		"java2": "FFan Java2",
		"java3": "FFan Java3",

		"php1": "FFan PHP1",
		"php2": "FFan PHP2",
		"php3": "FFan PHP3",

		"kafka": "FFan Kafka",
		"tfs":   "FFan TFS",
		"solr":  "FFan Solr",
		"mc":    "FFan Memcache",
		"redis": "第三方中间件-redis",
	}

	for _, name := range names {
		if groupName, present := groupMapping[name]; present {
			initZabbix()
			hosts := HostsOfGroup(groupName)
			for _, host := range hosts {
				fmt.Printf("%8s: %s\n", name, host.Name)
			}
		} else {
			fmt.Printf("unknown group: %s\nvalid groups:", name)
			for k, _ := range groupMapping {
				fmt.Printf(" %s", k)
			}
			fmt.Println()

		}
	}

}

func main() {
	if len(os.Args) <= 1 {
		cliLoop()
		return
	}

	// not interactive mode, just query host ip in a group
	showHostsOfGroup(os.Args[1:])

}
