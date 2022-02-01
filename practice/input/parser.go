package input

import (
	"errors"
	"os"
	"practice/pizza"
	"strings"
)

func FromFile(name string) ([]pizza.Client, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")

	lines = lines[1 : len(lines)-1]

	if len(lines)%2 != 0 {
		return nil, errors.New("number of lines in file is incorrect")
	}

	clients := make([]pizza.Client, 0, len(lines)/2)
	for i := 0; i+1 <= len(lines)-1; i += 2 {
		likesLine := lines[i]
		likes := strings.Split(likesLine, " ")[1:]

		dislikesLine := lines[i+1]
		dislikes := strings.Split(dislikesLine, " ")[1:]

		client := pizza.Client{
			Likes:    likes,
			Dislikes: dislikes,
		}

		clients = append(clients, client)
	}

	return clients, nil
}
