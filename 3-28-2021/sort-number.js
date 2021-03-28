function solution(nums) {
  if (nums === null) {
    return [];
  }
  nums.sort(function (a, b) {
    return a - b;
  });
  return nums;
}

console.log(solution([1, 2, 10, 50, 5]));
// should return [1,2,5,10,50]
solution(null); // should return []
