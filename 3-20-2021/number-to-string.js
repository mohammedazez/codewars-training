function numberToString(num) {
  console.log(num);
  let string = "";
  console.log(string);
  if (typeof num === "number") {
    console.log(num);
    string = `${num}`;
    console.log(string);
    return string;
  } else {
    return `tidak ada`;
  }
}

console.log(numberToString(123)); // returns '123'
console.log(numberToString(999)); // returns '999';`
