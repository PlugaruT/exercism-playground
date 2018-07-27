# Score categories
# Change the values as you see fit


def yacht(x: list) -> int:
    if len(set(x)) == 1:
        return 50
    else:
        return 0


def ones(x: list) -> int:
    return x.count(1)


def twos(x: list) -> int:
    return x.count(2) * 2


def threes(x: list) -> int:
    return x.count(3) * 3


def fours(x: list) -> int:
    return x.count(4) * 4


def fives(x: list) -> int:
    return x.count(5) * 5


def sixes(x: list) -> int:
    return x.count(6) * 6


def full_house(x: list) -> int:
    if {x.count(y) for y in x} == {2, 3}:
        return sum(x)
    else:
        return 0


def four_of_a_kind(x: list) -> int:
    if max({x.count(y) for y in x}) >= 4:
        return max(set(x), key=x.count) * 4
    else:
        return 0


def little_straight(x: list) -> int:
    if sorted(x) == [1, 2, 3, 4, 5]:
        return 30
    else:
        return 0


def big_straight(x: list) -> int:
    if sorted(x) == [2, 3, 4, 5, 6]:
        return 30
    else:
        return 0


def choice(x: list) -> int:
    return sum(x)


def score(dice: list, category) -> int:
    return category(dice)


FIVES = fives
FOURS = fours
THREES = threes
ONES = ones
TWOS = twos
YACHT = yacht
BIG_STRAIGHT = big_straight
LITTLE_STRAIGHT = little_straight
FOUR_OF_A_KIND = four_of_a_kind
FULL_HOUSE = full_house
SIXES = sixes
CHOICE = choice
