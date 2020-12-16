package main

import (
	"fmt"
	"github.com/adnanbrq/nanoleaf"
	"github.com/jinzhu/configor"
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
	path := filepath.Dir(os.Args[0]) + "/config/nanoleaf-cmd.yml"

	err := configor.Load(&Config, path)
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
	if !brightness(nano, cmd) {
		if !colorTemp(nano, cmd) {
			effect(nano, cmd)
		}
		if len(args) > 1 {
			brightness(nano, args[1])
		}
	}
}

func brightness(nano *nanoleaf.Nanoleaf, cmd string) (parsed bool) {
	if v, err := strconv.Atoi(cmd); err == nil {
		if v == 0 {
			err = nano.State.SetOn(false)
			if err != nil {
				panic(err)
			}
		} else if v > 0 {
			err = nano.State.SetBrightness(v, Config.Duration)
		}
		return true
	}
	return false
}

func colorTemp(nano *nanoleaf.Nanoleaf, cmd string) (parsed bool) {
	if strings.HasSuffix(cmd, "k") { // set color temp
		colorTemp, err := strconv.Atoi(cmd[:len(cmd)-1])
		if err != nil {
			panic(err)
		}
		err = nano.State.SetColorTemp(colorTemp, false)
		if err != nil {
			panic(err)
		}
		return true
	}
	return false
}

func effect(nano *nanoleaf.Nanoleaf, cmd string) {
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
