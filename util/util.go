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
	if fileJ, err = ioutil.ReadFile("res/j.mp3"); err != nil {
		log.Fatal(err)
	}
	if fileN, err = ioutil.ReadFile("res/n.mp3"); err != nil {
		log.Fatal(err)
	}
	if fileT, err = ioutil.ReadFile("res/t.mp3"); err != nil {
		log.Fatal(err)
	}
	if fileM, err = ioutil.ReadFile("res/m.mp3"); err != nil {
		log.Fatal(err)
	}
	if fileNGM, err = ioutil.ReadFile("res/ngm.mp3"); err != nil {
		log.Fatal(err)
	}

	var decJ *minimp3.Decoder

	if decJ, dataJ, err = minimp3.DecodeFull(fileJ); err != nil {
		log.Fatal(err)
	}
	if _, dataN, err = minimp3.DecodeFull(fileN); err != nil {
		log.Fatal(err)
	}
	if _, dataT, err = minimp3.DecodeFull(fileT); err != nil {
		log.Fatal(err)
	}
	if _, dataM, err = minimp3.DecodeFull(fileM); err != nil {
		log.Fatal(err)
	}
	if _, dataNGM, err = minimp3.DecodeFull(fileNGM); err != nil {
		log.Fatal(err)
	}

	var context *oto.Context
	if context, err = oto.NewContext(decJ.SampleRate, decJ.Channels, 2, 1024); err != nil {
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
