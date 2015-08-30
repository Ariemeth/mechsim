package controller

import (
	"bufio"
	"os"
	"strings"
)

//NewInput is used to create an input controller
func NewInput(listener chan string) *input {
	controller := input{}
	controller.isRunning = true
	controller.listener = listener
	controller.isDone = make(chan bool)

	go controller.process()
	return &controller
}

//Input is a controller designed for passing input controls.
type input struct {
	listener  chan string
	isDone    chan bool
	isRunning bool
}

//Stop causes the Input controller to stop.
func (controller *input) Stop() {
	controller.isDone <- true
	controller.isRunning = false
	close(controller.isDone)
}

//process runs a loop awaiting input from the keyboard and passing
//it along to any channel registered.
func (controller *input) process() {

	outChannel := make(chan string)
	defer close(outChannel)
	exitChannel := make(chan bool)

	go getInput(outChannel, exitChannel)

	for {
		select {
		case output := <-outChannel:
			controller.listener <- output
		case <-controller.isDone:
			exitChannel <- true
			return
		}
	}
}

//getInput handles waiting for keyboard input and sending
//the first character out on the out channel.
func getInput(out chan string, isDone chan bool) {
	reader := bufio.NewReader(os.Stdin)
	defer close(isDone)

	for {
		input, _ := reader.ReadString('\n')
		select {
		case out <- strings.Split(input, "")[0]:
		case done := <-isDone:
			if done {
				return
			}
		}
	}
}
