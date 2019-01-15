package main

import (
	"fmt"
	"os/exec"

	"github.com/gobuffalo/packr"
)

func main() {

	soundsBox := packr.NewBox("./sounds")
	if soundsBox.Has("IEEE_float_mono_32kHz.wav") {
		fmt.Println("It's there.")
	} else {
		fmt.Println("It's not there.")
	}

	args := []string{"-v20", "./sounds/IEEE_float_mono_32kHz.wav"}
	output, err := exec.Command("play", args...).Output()
	if err != nil {
		// Play command was not successful
		fmt.Println("Got an error.")
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(output))
	}

}
