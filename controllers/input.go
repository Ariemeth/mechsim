package controller

import (
	"bufio"
	"os"
	"strings"
)

func NewInput(interval int) *Input {
	newController := Input{}
	newController.isRunning = true
	
	go newController.process()
	return &newController
}

type Input struct {
	keys      []chan string
	isRunning bool
}

func (controller *Input) AddListener(listener chan string) {
	controller.keys = append(controller.keys, listener)
}

func (controller *Input) Stop() {
	controller.isRunning = false

	for _, channel := range controller.keys {
		close(channel)
	}
}

func (controller *Input) process() {
	reader := bufio.NewReader(os.Stdin)

	for ;controller.isRunning; {
		input, _ := reader.ReadString('\n')

		for index := 0 ; index < len(controller.keys); index++{
			controller.keys[index] <- strings.Split(input,"")[0]
		}
	}
}
