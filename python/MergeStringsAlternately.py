import itertools


def mergeAlternately(self, word1: str, word2: str) -> str:
    zipLong = itertools.zip_longest(word1, word2, fillvalue="")
    result = ""
    for q in zipLong:
        result += ''.join(q)
    return result


mergeAlternately(",", word1="ab", word2="pqrs")
