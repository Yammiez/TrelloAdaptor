package main

import (
    "io/ioutil"
	"fmt"
    "log"

    "github.com/VojtechVitek/go-trello"

)

func main() {
    // New Trello Client
	
	// Read appKey and token from file
	dat, err := ioutil.ReadFile("W:\\GoWork\\key\\Trello.com\\builder\\key")
	appKey := string(dat)
	
	dat, err = ioutil.ReadFile("W:\\GoWork\\key\\Trello.com\\builder\\token")
	token := string(dat)

	trelloClient, err := trello.NewAuthClient(appKey, &token)
    if err != nil {
        log.Fatal(err)
    }

    // User @trello
    user, err := trelloClient.Member("builder25")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(user.FullName)

    // @trello Boards
    boards, err := user.Boards()
    if err != nil {
        log.Fatal(err)
    }

/*	
	if len(boards) > 0 {
		board := boards[0]
	} else {
		panic("No boards found!")
	}
*/
	board := boards[0]
	
	// Get TestBoard
	for _, boardIterator := range(boards) {
	    if boardIterator.Name == "TestBoard" {
			fmt.Printf("Found TestBoard!\n")
			board = boardIterator
		}
	}
	
	fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)
	
	// @trello Board Lists
	lists, err := board.Lists()
	if err != nil {
		log.Fatal(err)
	}

	for _, list := range lists {
		fmt.Println("   - ", list.Name)

		// @trello Board List Cards
		cards, _ := list.Cards()
		for _, card := range cards {
			fmt.Println("      + ", card.Name)
		}
	}
	
	// Test adding a card to the first list
	firstList := lists[0]
	firstList.AddCard(trello.Card{Name : "My New Card", Pos : 1})
	fmt.Printf("Card added to the first list!\n")
}