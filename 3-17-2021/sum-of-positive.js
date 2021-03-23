function positiveSum(arr) {
  console.log(arr);

  var sum = 0;
  console.log(sum);

  for (var i = 0; i < arr.length; i++) {
    console.log(arr[i]);

    if (arr[i] > 0) {
      console.log(arr[i]);
      sum += arr[i];
    }
  }
  return sum;
}

console.log(positiveSum([1, -2, 3, 4, 5]));
console.log(positiveSum([]), 0);
console.log(positiveSum([-1, -2, -3, -4, -5]), 0);
console.log(positiveSum([1, 2, 3, 4, 5]), 15);
console.log(positiveSum([-1, 2, 3, 4, -5]), 9);
