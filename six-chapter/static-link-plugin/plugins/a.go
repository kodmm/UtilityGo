package a

import (
	"fmt"

	staticplugin "github.com/kodmm/UtilityGo/second-chapter/six-chapter/static-link-plugin"
)

type pluginA struct{}

func (p pluginA) Exec() {
	fmt.Println("pllugin A")
}

func init() {
	staticplugin.Register("a", &pluginA{})
}
