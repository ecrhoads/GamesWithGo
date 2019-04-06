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

type npc struct {
	name        string
	description string
	health      int
}

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
	npc     []*npc
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) addNPC(name string, description string, health int) {
	character := &npc{name, description, health}
	node.npc = append(node.npc, character)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.npc != nil {
		for _, character := range node.npc {
			fmt.Println("\n You are engaged with a character.")
			fmt.Println(" Name: ", character.name, "\n", "Description: ", character.description, "\n", "Health: ", character.health)
		}
	}
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

	abc := []string{"a", "b", "c"}
	abc = append(abc, "d")
	printArray(abc)

	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You are in a large chamber, deep underground.
	You see three passages leading out. A north passage leads into darkness.
	To the south, a passage appears to head upward. The eastern passage appears
	flat and well traveled`}

	darkRoom := storyNode{text: "\n It is pitch black. You cannot see a thing."}
	darkRoomLit := storyNode{text: "\n The dark passage is now lit by your lantern."}
	grue := storyNode{text: "\n While stumbling around in the darkness, you are eaten by The Grue."}
	trap := storyNode{text: "\n You head down the well traveled path when suddenly a trap door opens and you fall into a pit."}
	treasure := storyNode{text: "\n You arrive at a small chamber, filled with treasure!"}
	death := storyNode{text: "\n The Grue kills you."}
	fight := storyNode{text: "\n You wound The Grue with your trusty sword. He lunges toward you, snarling."}
	champion := storyNode{text: "\n You slay The Grue."}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addNPC("The Grue", "A ferocious creature who feasts on human flesh.", 200)
	darkRoomLit.addChoice("A", "Attack", &fight)
	darkRoomLit.addChoice("R", "Run Away", &start)

	fight.addNPC("The Grue", "A ferocious creature who feasts on human flesh.", 100)
	fight.addChoice("R", "Run away", &death)
	fight.addChoice("D", "Dodge the attack", &death)
	fight.addChoice("A", "Attack again", &champion)

	champion.addChoice("N", "Continue heading North", &treasure)
	champion.addChoice("S", "Turn back South", &start)
	start.play()

	fmt.Println()
	fmt.Println("The End.")
}
