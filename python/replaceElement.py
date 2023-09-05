from typing import List


def replaceElements(self, arr: List[int]) -> List[int]:
    rightMax = -1
    for i in range(len(arr) - 1, -1, -1):
        newMax = max(rightMax, arr[i])
        arr[i] = rightMax
        rightMax = newMax
    return arr


replaceElements("", arr=[17, 18, 5, 4, 6, 1])  # [18, 6, 6, 6, 1, -1]
