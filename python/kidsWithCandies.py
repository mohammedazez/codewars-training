from typing import List


def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
    maxNumber = max(candies)
    res = []
    for i in range(len(candies)):
        candies[i] += extraCandies
        if candies[i] >= maxNumber:
            candies[i] = True
            res.append(candies[i])
        else:
            candies[i] = False
            res.append(candies[i])
    return res


kidsWithCandies("", candies=[2, 3, 5, 1, 3], extraCandies=3)
