package main

import (
	staticplugin "github.com/kodmm/UtilityGo/second-chapter/six-chapter/static-link-plugin"
)

func main() {
	for _, p := range staticplugin.Plugins() {
		p.Exec()
	}
}
