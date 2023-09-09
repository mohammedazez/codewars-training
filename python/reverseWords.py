
def reverseWords(self, s: str) -> str:
    arr = s.split()[::-1]
    return " ".join(arr)


reverseWords("", s="  hello world  ")
