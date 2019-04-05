// Add a function that will insert a new page, after a given page.
// Add a function that will delete a page.

package main

import "fmt"

type storyPage struct { //this is a linked list
	text     string
	nextPage *storyPage
}

/*
func (receiver *storyPage) function() {
	receiver.nextPage.function() //this is a function with a receiver.
} */

/*
func playStory(page *storyPage) {
	if page == nil { //checking for a null pointer
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
} */

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(text string) {
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page1.addToEnd("You are alone, and you need to find the sacred helmet before the bad guys do.")
	page1.addToEnd("You see a troll ahead.")

	page1.playStory()

	//Functions - has return value - may also execute commands
	//Procedures - has no return value, just executes commands
	//Methods - functions attached to a stuct/object/etc
}
