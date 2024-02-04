package main

import (
	"main/pipeline"
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
	// errdemo.DemoErrGroup()
	// pipeline.DemoPipeline()
	pipeline.DemoFanoutFanin()
}
