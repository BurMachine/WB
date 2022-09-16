package main

import "simple/Command/Command"

func main() {
	tv := &Command.Tv{}
	onCommand := &Command.OnCommand{Device: tv}
	offCommand := &Command.OffCommand{Device: tv}
	onButton := &Command.Button{Command: onCommand}
	onButton.Press()
	offButton := &Command.Button{Command: offCommand}
	offButton.Press()
}
