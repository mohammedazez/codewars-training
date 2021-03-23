function squareSum(numbers) {
  console.log(numbers);
  let result = 0;

  for (let i = 0; i < numbers.length; i++) {
    let kali = numbers[i] * numbers[i];
    result += kali;
    console.log(kali);
    console.log(numbers[i]);
    console.log(result);
  }
  return result;
}

console.log(squareSum([1, 2]), 5);
console.log(squareSum([0, 3, 4, 5]), 50);
