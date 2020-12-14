package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"nanoleaf"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var Config = struct {
	Ip       string `required:"true"`
	Token    string `required:"true"`
	Duration int    `default:"0"`
}{}

func main() {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	err = configor.Load(&Config, path+"/config/lights.yml")
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]
	nano := nanoleaf.NewNanoleaf("http://" + Config.Ip + ":16021/api/v1")
	nano.SetToken(Config.Token)

	if len(args) == 0 { // toggle on/off without args
		on, err := nano.State.IsOn()
		if err != nil {
			panic(err)
		}
		err = nano.State.SetOn(!on)
		if err != nil {
			panic(err)
		}
		return
	}

	cmd := args[0]
	if v, err := strconv.Atoi(cmd); err == nil { // set brightness
		if v == 0 {
			err = nano.State.SetOn(false)
			if err != nil {
				panic(err)
			}
		} else if v > 0 {
			err = nano.State.SetBrightness(v, Config.Duration)
		}
	} else {
		if strings.HasSuffix(cmd, "k") { // set color temp
			colorTemp, err := strconv.Atoi(cmd[:len(cmd)-1])
			if err != nil {
				panic(err)
			}
			err = nano.State.SetColorTemp(colorTemp, false)
			if err != nil {
				panic(err)
			}
		} else { // set effect if existing, otherwise list
			effects, err := nano.Effects.List()
			if err != nil {
				panic(err)
			}
			for _, effect := range effects {
				if cmd == effect {
					err = nano.Effects.Set(cmd)
					if err != nil {
						panic(err)
					}
					return
				}
			}
			for _, effect := range effects {
				fmt.Println(effect)
			}
		}
	}
}
