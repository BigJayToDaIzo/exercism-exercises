package twelve

import "fmt"

func Verse(i int) string {
	days := map[int][]string{
		1:  {"first", "a Partridge in a Pear Tree."},
		2:  {"second", "two Turtle Doves, and "},
		3:  {"third", "three French Hens, "},
		4:  {"fourth", "four Calling Birds, "},
		5:  {"fifth", "five Gold Rings, "},
		6:  {"sixth", "six Geese-a-Laying, "},
		7:  {"seventh", "seven Swans-a-Swimming, "},
		8:  {"eighth", "eight Maids-a-Milking, "},
		9:  {"ninth", "nine Ladies Dancing, "},
		10: {"tenth", "ten Lords-a-Leaping, "},
		11: {"eleventh", "eleven Pipers Piping, "},
		12: {"twelfth", "twelve Drummers Drumming, "},
	}
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[i][0])
	for v := i; v > 0; v-- {
		verse += fmt.Sprintf("%s", days[v][1])
	}
	return verse
}

func Song() string {
	song := ""
	for verse := 1; verse <= 12; verse++ {
		if verse == 12 {
			song += Verse(verse)
		} else {
			song += Verse(verse) + "\n"
		}
	}
	return song
}
