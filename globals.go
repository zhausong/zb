package main

import (
	"github.com/AlekSi/zabbix"
	ui "github.com/funkygao/termui"
	"os"
	"time"
)

var (
	api *zabbix.API

	hosts zabbix.Hosts
	data  [][]float64

	itemName = "gw.accesslog.count-$1 per 5mins"
	group    = "FFan Gateway"

	cliHostId   string
	cliGroupId  string
	cliItemName string

	user = os.Getenv("ZABBIX_USER")
	pass = os.Getenv("ZABBIX_PASS")
)

const (
	apiUrl = "http://zabbix.intra.wanhui.cn/zabbix/api_jsonrpc.php"

	panelHeight     = 11
	titleHeight     = 1
	dataSize        = 95
	panelsPerRow    = 2
	chartPaddingTop = 1
	panelSpan       = 6
	panelOffset     = 0
	uiTheme         = "helloworld"
	lineColor       = ui.ColorYellow | ui.AttrBold
	axesColor       = ui.ColorWhite

	tick = time.Minute * 5

	emptyStr = ""
)
