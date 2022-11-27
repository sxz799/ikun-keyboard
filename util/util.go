package util

import (
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"io/ioutil"
	"log"
)

var player *oto.Player

var dataJ, dataN, dataT, dataM, dataNGM []byte

func init() {
	var err error
	var fileJ, fileN, fileT, fileM, fileNGM []byte
	fileJ, err = ioutil.ReadFile("res/j.mp3")
	fileN, err = ioutil.ReadFile("res/n.mp3")
	fileT, err = ioutil.ReadFile("res/t.mp3")
	fileM, err = ioutil.ReadFile("res/m.mp3")
	fileNGM, err = ioutil.ReadFile("res/ngm.mp3")
	var dec *minimp3.Decoder
	dec, dataJ, _ = minimp3.DecodeFull(fileJ)
	_, dataN, err = minimp3.DecodeFull(fileN)
	_, dataT, err = minimp3.DecodeFull(fileT)
	_, dataM, err = minimp3.DecodeFull(fileM)
	_, dataNGM, err = minimp3.DecodeFull(fileNGM)
	var context *oto.Context
	if context, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, 1024); err != nil {
		log.Fatal(err)
	}
	player = context.NewPlayer()
}

func PlaySound(str string) {
	switch str {
	case "J":
		player.Write(dataJ)
	case "N":
		player.Write(dataN)
	case "T":
		player.Write(dataT)
	case "M":
		player.Write(dataM)
	case "JNTM":
		player.Write(dataNGM)
	}

}
