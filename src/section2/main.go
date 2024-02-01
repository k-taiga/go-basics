package main

import (
	errdemo "main/errDemo"
)

func main() {
	// tracer.DemoTracer()
	// channel.DemoChannel()
	// channel.ChannelClose()
	// channel.DemoSelect()
	// channel.DemoSelectContinuous()
	// context.DemoContextWithTimeout()
	// context.DemoContextWithCancel()
	// context.DemoContextWithDeadline()
	errdemo.DemoErrGroup()
}
