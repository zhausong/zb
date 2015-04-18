package main

import (
	"bytes"
	"io/ioutil"
)

func loadFavorites() {
	contents, err := ioutil.ReadFile(favoriteFileName)
	if err != nil {
		return
	}
	bs := bytes.NewBuffer(contents)
	for {
		line, err := bs.ReadString('\n')
		if err != nil {
			break
		}
		favoriteItems[line] = true
	}
}

func addFavorite(item string) {
	favoriteItems[item] = true
}
