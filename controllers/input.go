package controller

import (
	"bufio"
	"os"
	"strings"
)

//NewInput is used to create an input controller
func NewInput() *Input {
	newController := Input{}
	newController.isRunning = true
	
	go newController.process()
	return &newController
}

//Input is a controller designed for passing input controls.
type Input struct {
	keys      []chan string
	isRunning bool
}

//AddListener adds a channel that will allow input to be passed
//to a listener.
func (controller *Input) AddListener(listener chan string) {
	controller.keys = append(controller.keys, listener)
}

//RemoveListener removes a channel.
func (controller *Input) RemoveListener(listener chan string){
	
}

//Stop causes the Input controller to stop.
func (controller *Input) Stop() {
	controller.isRunning = false

	for _, channel := range controller.keys {
		close(channel)
	}
}

//process runs a loop awaiting input from the keyboard and passing
//it along to any channel registered.
func (controller *Input) process() {
	reader := bufio.NewReader(os.Stdin)

	for ;controller.isRunning; {
		input, _ := reader.ReadString('\n')

		for index := 0 ; index < len(controller.keys); index++{
			controller.keys[index] <- strings.Split(input,"")[0]
		}
	}
}
