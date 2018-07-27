def is_valid(checksum: int) -> bool:
    return checksum % 11 == 0


def verify(isbn: str) -> bool:
    isbn = isbn.replace("-", "")
    check_sum = 0

    if not len(isbn) == 10:
        return False

    for index, digit in enumerate(isbn):
        if index == len(isbn) - 1 and digit == "X":
            digit = "10"
        if digit.isdigit():
            check_sum += int(digit) * (10 - index)
        else:
            break
    return is_valid(check_sum)
