package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/jcocozza/gotomata/common/continious"
//	"github.com/jcocozza/gotomata/common/elementary"
)

const (
	width  = 150
	height = 300
	depth  = 300
	steps  = 120
)

func main() {
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	// call some main function from common/ or create your own here
	//initCfg := elementary.SetCenterConfig(width)
	//elementary.MainElementary(30, width, steps, initCfg)

	initCfg2 := continious.SetCenterConfig(width)
	continious.MainContinious(width, steps, 10, initCfg2)
}
