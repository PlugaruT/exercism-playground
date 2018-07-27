from datetime import timedelta, datetime


def add_gigasecond(birth_date: datetime) -> datetime:
    return birth_date + timedelta(seconds=10 ** 9)
