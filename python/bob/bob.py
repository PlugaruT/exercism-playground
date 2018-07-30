def is_yelled(string: str) -> bool:
    return string.isupper()


def is_question(string: str) -> bool:
    return string.endswith("?")


def is_yelled_question(string: str) -> bool:
    return is_yelled(string) and is_question(string)


def hey(phrase: str) -> str:
    phrase = phrase.strip()
    if is_yelled_question(phrase):
        return "Calm down, I know what I'm doing!"
    elif is_yelled(phrase):
        return "Whoa, chill out!"
    elif is_question(phrase):
        return "Sure."
    elif not phrase:
        return "Fine. Be that way!"
    else:
        return "Whatever."
