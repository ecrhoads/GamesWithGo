package main

// HOMEWORK
// NPC - move around graph
// items picked up or placed down
// accept natural language as input
// build your own game -- see what you can accomplish just with this system

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)

	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println()
			fmt.Println(choice.cmd, choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func printArray(a []string) {
	for _, e := range a {
		fmt.Print(e)
	}
}

func main() {

	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You are in a large chamber, deep underground.
	You see three passages leading out. A north passage leads into darkness.
	To the south, a passage appears to head upward. The eastern passage appears
	flat and well traveled`}

	darkRoom := storyNode{text: "\n It is pitch black. You cannot see a thing."}
	darkRoomLit := storyNode{text: "\n The dark passage is now lit by your lantern. A Troll stands in the way."}
	grue := storyNode{text: "\n While stumbling around in the darkness, you are eaten by a Grue."}
	trap := storyNode{text: "\n You head down the well traveled path when suddenly a trap door opens and you fall into a pit."}
	treasure := storyNode{text: "\n You arrive at a small chamber, filled with treasure!"}
	game := storyNode{text: "The troll evades your attack. He rumbles toward you and raises his spiked club."}
	death := storyNode{text: "The troll swings his club down on you and crushes you like a bug."}
	victory := storyNode{text: "You slay the troll."}
	lastRoom := storyNode{text: "The passage way has two paths. One has a terrible odor. The other appears to have been well traveled."}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addChoice("A", "Attack", &game)
	darkRoomLit.addChoice("R", "Run Away", &start)

	game.addChoice("A", "Attack again", &victory)
	game.addChoice("B", "Try to block", &death)

	victory.addChoice("N", "Continue North", &lastRoom)
	victory.addChoice("S", "Turn back South", &start)

	lastRoom.addChoice("L", "Turn left down the path with the terrible odor.", &treasure)
	lastRoom.addChoice("R", "Turn right down the well traveled path.", &trap)
	lastRoom.addChoice("S", "Turn back South", &start)

	start.play()

	fmt.Println()
	fmt.Println("The End.")
}
