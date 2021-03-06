// Project to simulate mechs to learn go.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Ariemeth/mechsim/controllers"
	"github.com/Ariemeth/mechsim/mech"
	"github.com/Ariemeth/mechsim/mech/weapon"
	
	"github.com/Ariemeth/termloop"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	game := termloop.NewGame()
	level := termloop.NewBaseLevel(termloop.Cell{
        Bg: termloop.ColorGreen,
        Fg: termloop.ColorBlack,
        Ch: 'v',
    })
    level.AddEntity(termloop.NewRectangle(10, 10, 50, 20, termloop.ColorBlue))
    game.SetLevel(level)
    game.Start()

	inputChannel := make(chan string)

	inputController := controller.NewInput(inputChannel)

	mech1 := mech.NewMech("Mech1", 2)
	mech2 := mech.NewMech("Mech2", 2)

	weapon1 := weapon.Create(5, 1, "rifle", .75)
	weapon2 := weapon.Create(3, 3, "shotgun", .4)

	mech1.AddWeapon(weapon1)
	mech2.AddWeapon(weapon2)

	mech1.Fire(4, mech2)

	fmt.Println("Mech1 has ", mech1.StructureLeft(), " structure left.")
	fmt.Println("Mech2 has ", mech2.StructureLeft(), " structure left.")

	isRunning := true

	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout

	for isRunning {
		select {
		case msg := <-inputChannel:
			c := exec.Command("cmd", "/c", "cls")
			c.Stdout = os.Stdout
			c.Run()
			if strings.EqualFold(strings.ToLower(msg), "q") {
				isRunning = false
			}
		}
	}

	inputController.Stop()
	close(inputChannel)
}
