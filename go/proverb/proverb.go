// Proverb package proverbiates
package proverb

// Proverb returns a proverb post proverbiated
func Proverb(rhyme []string) []string {
	proverb := []string{}
	if len(rhyme) == 0 {
		return proverb
	}
	if len(rhyme) == 1 {
		return []string{"And all for the want of a " + rhyme[0] + "."}
	}
	for i, str := range rhyme {
		if i == len(rhyme)-1 {
			proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")
		} else {
			proverb = append(proverb, "For want of a "+str+" the "+rhyme[i+1]+" was lost.")
		}
	}
	return proverb
}
