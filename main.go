package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/jcocozza/gotomata/common/totalistic"
	// "github.com/jcocozza/gotomata/common/continuous"
	// "github.com/jcocozza/gotomata/common/elementary"
)

const (
	width  = 300
	height = 300
	depth  = 300
	steps  = 300
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

	//initCfg2 := continuous.SetCenterConfig(width)
	//continuous.MainContinuous(width, steps, 10, initCfg2)

	initCfg := totalistic.SetCenterConfig(width)
	totalistic.MainTotalistic(912, width, steps, 10, initCfg)
}
