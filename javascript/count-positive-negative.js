function countPositivesSumNegatives(input) {
  let newArray = [];
  let positive = 0;
  let negative = 0;

  if (input == 0 || input == null) {
    return [];
  }

  for (let i = 0; i < input.length; ++i) {
    if (input[i] > 0 && input[i] != 0) {
      positive++;
    }
    if (input[i] < 0) {
      console.log(input[i]);
      negative += input[i];
    }
  }
  let result = newArray.concat(positive, negative);

  return result;
}

console.log(countPositivesSumNegatives([]));
console.log(
  countPositivesSumNegatives(
    [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15],
    [10, -65]
  )
);

console.log(
  countPositivesSumNegatives(
    [0, 2, 3, 0, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14],
    [8, -50]
  )
);
