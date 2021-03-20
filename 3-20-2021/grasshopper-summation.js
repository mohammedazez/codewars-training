function summation(num) {
  console.log(num);
  let newNumber = 0;
  console.log(newNumber);

  for (let i = 1; i <= num; i++) {
    console.log(i);
    newNumber += i;
    console.log(newNumber);
  }
  return newNumber;
}

// console.log(summation(1));
console.log(summation(8));
