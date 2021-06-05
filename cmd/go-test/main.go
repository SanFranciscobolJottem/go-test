package main

import (
	"github.com/SanFranciscobolJottem/go-test/pkg/discord"
	"github.com/SanFranciscobolJottem/go-test/pkg/server"
)

func main() {
	server.Run()      //  meghivhatom masik package-bol, mert nagy betuvel kezdodik a fuggveny neve.
	discord.Connect() //  meghivhatom masik package-bol, mert nagy betuvel kezdodik a fuggveny neve.
	// discord.doStuff()  - ezt nem hivhatom meg, mert a fuggveny neve kisbetuvel kezdodik, vagyis nem exportalt!
	printStart() //       - meghivhatom kisbetuvel, mert ugyan ebben a package-ben el. Az, hogy masik fajlban, az nem szamit.
}
