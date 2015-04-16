package main

import (
	"fmt"
	ui "github.com/funkygao/termui"
	tm "github.com/nsf/termbox-go"
	"time"
)

func drawDashboard() {
	hosts = HostsOfGroup(group)
	data = make([][]float64, len(hosts))

	// title
	titlePanel := ui.NewPar(fmt.Sprintf("%s: %s / %v", group, itemName, tick*dataSize))
	titlePanel.Height = titleHeight
	titlePanel.PaddingLeft = 1
	titlePanel.HasBorder = false
	titlePanel.TextFgColor = ui.ColorCyan

	// charts
	charts := make([]*ui.LineChart, 0)
	for i := 0; i < len(data); i++ {
		chart := ui.NewLineChart()
		chart.Border.Label = hosts[i].Name
		chart.Data = fetchData(i, dataSize)
		chart.Height = panelHeight
		chart.PaddingTop = chartPaddingTop
		chart.AxesColor = axesColor
		chart.LineColor = lineColor

		charts = append(charts, chart)
	}

	draw := func(size int) {
		if size > 0 {
			for i := 0; i < len(data); i++ {
				charts[i].Data = fetchData(i, size)
			}
		}

		ui.Render(ui.Body)
	}

	err := ui.Init()
	must(err)
	defer ui.Close()

	ui.UseTheme(uiTheme)

	// auto layout
	rows := make([]*ui.Row, 1)
	rows[0] = ui.NewRow(ui.NewCol(panelSpan*panelsPerRow, panelOffset, titlePanel))
	for i := 0; i < len(hosts); i += panelsPerRow {
		if i+1 == len(hosts) {
			// the last single panel
			rows = append(rows,
				ui.NewRow(ui.NewCol(panelSpan, panelOffset, charts[i])))
		} else {
			rows = append(rows,
				ui.NewRow(ui.NewCol(panelSpan, panelOffset, charts[i]),
					ui.NewCol(panelSpan, panelOffset, charts[i+1])))
		}

	}
	ui.Body.AddRows(rows...)
	ui.Body.Align()

	// draw the history data
	draw(0)

	evt := make(chan tm.Event)
	go func() {
		for {
			evt <- tm.PollEvent()
		}
	}()

	for {
		select {
		case e := <-evt:
			if e.Type == tm.EventKey && e.Ch == 'q' {
				return
			}
			if e.Type == tm.EventResize {
				ui.Body.Width = ui.TermWidth()
				ui.Body.Align()
			}

		case <-time.After(tick):
			draw(1)
		}
	}
}
