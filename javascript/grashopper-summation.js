function summation(num) {
  console.log(num);
  let result = 0;

  for (let i = 1; i < num + 1; i++) {
    console.log(i);
    result += i;
  }
  return result;
}

console.log(summation(1), 1);
console.log(summation(8), 36);
