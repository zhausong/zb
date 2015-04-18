package main

import (
	"fmt"
	"github.com/AlekSi/zabbix"
	"os"
	u "os/user"
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

	api = zabbix.NewAPI(apiUrl)
	auth, err := api.Login(user, pass)
	must(err)
	fmt.Printf("auth: %s\n", auth)

	cliGroupId = GroupId(group)

	currentUser, _ := u.Current()
	favoriteFileName = currentUser.HomeDir + "/.zb"
	loadFavorites()
}

func main() {
	cliLoop()
}
