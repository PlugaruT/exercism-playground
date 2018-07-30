class Allergies(object):

    _allergies = [
        "eggs",  # 1
        "peanuts",  # 2
        "shellfish",  # 4
        "strawberries",  # 8
        "tomatoes",  # 16
        "chocolate",  # 32
        "pollen",  # 64
        "cats",  # 128
    ]

    def __init__(self, score: int):
        self.score = score

    # TODO [Tudor] Review this again
    def is_allergic_to(self, item: str):
        return bool(self.score & 1 << self._allergies.index(item))

    @property
    def lst(self):
        return [allergy for allergy in self._allergies if self.is_allergic_to(allergy)]
