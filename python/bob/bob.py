def response(hey_bob):
    if hey_bob.isspace() or hey_bob == "":
        return "Fine. Be that way!"
    hey_bob = hey_bob.strip()
    isQuestion = hey_bob[-1] == "?"
    isShouting = hey_bob.isupper()
    if isQuestion and not isShouting:
        return "Sure."
    elif isQuestion and isShouting:
        return "Calm down, I know what I'm doing!"
    elif isShouting:
        return "Whoa, chill out!"
    return "Whatever."
