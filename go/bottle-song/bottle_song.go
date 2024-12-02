package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	transTbl := map[int]string{
		10: "Ten",
		9:  "Nine",
		8:  "Eight",
		7:  "Seven",
		6:  "Six",
		5:  "Five",
		4:  "Four",
		3:  "Three",
		2:  "Two",
		1:  "One",
		0:  "no",
	}
	song := []string{}
	for i := takeDown; i > 0; i-- {
		stanza := fmt.Sprintf("%s green bottles hanging on the wall,", transTbl[startBottles])
		if startBottles == 1 {
			stanza = "One green bottle hanging on the wall,"
		}
		song = append(song, stanza)
		song = append(song, stanza)
		song = append(song, "And if one green bottle should accidentally fall,")
		startBottles--
		if startBottles == 1 {
			song = append(song, fmt.Sprintf("There'll be one green bottle hanging on the wall."))
		} else {
			song = append(song, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(transTbl[startBottles])))
		}
		if i > 1 {
			song = append(song, "")
		}
	}
	return song
}
