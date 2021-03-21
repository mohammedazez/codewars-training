function makeNegative(number) {
  console.log(number);
  if (number >= 0) {
    return -Math.abs(number);
  }

  if (number < 0) {
    return number;
  }
  return number;
}

console.log(makeNegative(1));
console.log(makeNegative(-5));
console.log(makeNegative(0));
console.log(makeNegative(0.12)); // return -1 // return -5 // return 0
console.log(makeNegative(42), -42);
