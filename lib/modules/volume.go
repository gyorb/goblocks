package modules

import (
	"fmt"
	"github.com/davidscholberg/go-i3barjson"
	"os/exec"
	"strings"
)

type Volume struct {
	BlockIndex     int `mapstructure:"block_index"`
	UpdateInterval int `mapstructure:"update_interval"`
	UpdateSignal   int `mapstructure:"update_signal"`
}

func (c Volume) GetBlockIndex() int {
	return c.BlockIndex
}

func (c Volume) GetUpdateFunc() func(b *i3barjson.Block, c BlockConfig) {
	return updateVolumeBlock
}

func (c Volume) GetUpdateInterval() int {
	return c.UpdateInterval
}

func (c Volume) GetUpdateSignal() int {
	return c.UpdateSignal
}

func updateVolumeBlock(b *i3barjson.Block, c BlockConfig) {
	fullTextFmt := "V: %s"
	amixerCmd := "amixer"
	amixerArgs := []string{"-D", "default", "get", "Master"}
	out, err := exec.Command(amixerCmd, amixerArgs...).Output()
	if err != nil {
		b.FullText = fmt.Sprintf(fullTextFmt, err.Error())
		return
	}
	outStr := string(out)
	iBegin := strings.Index(outStr, "[")
	if iBegin == -1 {
		b.FullText = fmt.Sprintf(fullTextFmt, "cannot parse amixer output")
		return
	}
	iEnd := strings.Index(outStr, "]")
	if iEnd == -1 {
		b.FullText = fmt.Sprintf(fullTextFmt, "cannot parse amixer output")
		return
	}
	b.FullText = fmt.Sprintf(fullTextFmt, outStr[iBegin+1:iEnd])
}
