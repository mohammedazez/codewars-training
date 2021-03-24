// function digitize(n) {
//   console.log(n);
//   let numArr = [];
//   console.log(numArr);

//   while (n > 0) {
//     console.log(numArr);
//     console.log(n);
//     numArr.push(n % 10);
//     n = Math.floor(n / 10);
//   }

//   return numArr;
// }

function digitize(n) {
  console.log(n);
  return (n + "").split("").map(Number).reverse();
}
console.log(digitize(35231), [1, 3, 2, 5, 3]);
