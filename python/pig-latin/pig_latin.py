def translate(text):
    # Gotta split phrases and iterate
    phrase = text.split(" ")
    vowels = ["a", "e", "i", "o", "u"]
    pl_word = ""
    for word in phrase:
        constsLeadingString = 0
        for i in range(len(word)):
            # find out how many constanants lead the string
            if word[i] not in vowels:
                constsLeadingString += 1
            else:
                break
        # Rule 1
        if constsLeadingString == 0 or word[0:2] in ["xr", "yt"]:
            pl_word += word + "ay "
            continue
        # Rule 3
        if constsLeadingString == 1 and word[:2] == "qu":
            pl_word += word[2:] + "quay "
            continue
        if word[constsLeadingString - 1: constsLeadingString + 1] == "qu":
            pl_word += (
                word[constsLeadingString + 1:]
                + word[: constsLeadingString + 1]
                + "ay "
            )
            continue
        # Rules 2 & 4
        if constsLeadingString > 0:  # Rules share this requirement
            # Rule 4
            if len(word) == 2 and word[-1] == "y":
                pl_word += word[1] + word[0] + "ay "
                continue
            if word[constsLeadingString] == "y":
                pl_word += (
                    word[constsLeadingString + 1:]
                    + word[: constsLeadingString + 1]
                    + "ay "
                )
            # Rule 2
            else:
                pl_word += (
                    word[constsLeadingString:] +
                    word[:constsLeadingString] + "ay "
                )
    return pl_word[:-1]
