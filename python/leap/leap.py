def is_leap_year(year):
    """
    Checks if a year is leap or not
    """
    return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)
