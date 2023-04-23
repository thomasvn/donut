package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Friend struct {
	Name  string   `yaml:"name"`
	Dates []string `yaml:"dates"`
}

type FriendsList struct {
	Friends []Friend `yaml:"friends"`
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: donut [friends_file] <list|pair>")
		return
	}

	// Default to using "friends.yaml" if friends_file not specified
	friendsFile := "friends.yaml"
	command := args[0]

	// Otherwise, use the friends_file provided
	if len(args) == 2 {
		friendsFile = args[0]
		command = args[1]
	}

	switch command {
	case "list":
		listFriends(friendsFile)
	case "pair":
		pairFriend(friendsFile)
	default:
		fmt.Println("Unknown command:", command)
	}
}

// Pretty print your friend file.
func listFriends(friendsFile string) {
	friends := loadFriends(friendsFile)

	if len(friends) == 0 {
		fmt.Println("No friends found.")
		return
	}

	fmt.Println("Friends list:")
	for _, friend := range friends {
		fmt.Printf("- %s (met on %v)\n", friend.Name, friend.Dates)
	}
}

// Continuously propose random friends from your friend list until you accept.
func pairFriend(friendsFile string) {
	friends := loadFriends(friendsFile)

	if len(friends) == 0 {
		fmt.Println("No friends found.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)

	for {
		randomIndex := rand.Intn(len(friends))
		friend := &friends[randomIndex]

		lastMet := "never"
		if len(friend.Dates) > 0 {
			lastMet = friend.Dates[len(friend.Dates)-1]
		}

		fmt.Printf("🍩 Reach out to: %s (last met: %s)? Accept (y), Decline (n): ", friend.Name, lastMet)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if strings.ToLower(response) == "y" {
			friend.Dates = append(friend.Dates, "TMP "+time.Now().Format("2006-01-02"))
			saveFriends(friendsFile, friends)
			fmt.Println("You have chosen to reach out to:", friend.Name)
			break
		} else if strings.ToLower(response) == "n" {
			continue
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}

// Loads data from your friends file.
func loadFriends(friendsFile string) []Friend {
	var friendsList FriendsList

	data, err := os.ReadFile(friendsFile)
	if err != nil {
		fmt.Println("Error reading friends.yaml:", err)
		return nil
	}

	err = yaml.Unmarshal(data, &friendsList)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return nil
	}

	return friendsList.Friends
}

// Writes data to your friends file.
func saveFriends(friendsFile string, friends []Friend) {
	friendsList := FriendsList{Friends: friends}
	data, err := yaml.Marshal(friendsList)
	if err != nil {
		fmt.Println("Error marshalling YAML:", err)
		return
	}

	err = os.WriteFile(friendsFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing", friendsFile+":", err)
	}
}
