package main

import (
	"github.com/AlekSi/zabbix"
	ui "github.com/funkygao/termui"
	"time"
)

var (
	api *zabbix.API

	hosts zabbix.Hosts
	data  [][]float64

	itemName = "gw.accesslog.count-$1 per 5mins"
	group    = "FFan Gateway"
)

const (
	apiUrl = "http://zabbix.intra.wanhui.cn/zabbix/api_jsonrpc.php"
	user   = "gaopeng27"
	pass   = "gaopeng27"

	panelHeight  = 11
	titleHeight  = 3
	dataSize     = 500
	panelsPerRow = 2
	panelSpan    = 6
	panelOffset  = 0
	uiTheme      = "helloworld"
	lineColor    = ui.ColorYellow | ui.AttrBold
	axesColor    = ui.ColorWhite

	tick = time.Minute * 5

	emptyStr = ""
)
