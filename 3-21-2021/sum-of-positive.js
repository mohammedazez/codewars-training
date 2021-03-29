function positiveSum(arr) {
  console.log(arr);
  let result = 0;
  for (i = 0; i < arr.length; i++) {
    console.log(arr[i]);
    if (arr[i] >= 0) {
      result += arr[i];
      console.log(arr[i]);
    }
  }
  return result;
}

console.log(positiveSum([1, 2, 3, 4, 5]), 15);
console.log(positiveSum([1, -2, 3, 4, 5]), 13);
console.log(positiveSum([]), 0);
console.log(positiveSum([-1, -2, -3, -4, -5]), 0);
console.log(positiveSum([-1, 2, 3, 4, -5]), 9);
