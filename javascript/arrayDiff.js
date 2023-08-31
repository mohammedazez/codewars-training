function arrayDiff(a, b) {
  let res = [];
  if (a.length == 0) {
    return res;
  }
  if (b.length == 0) {
    return (res = a);
  }

  let cba = 0;
  for (let i = 0; i < a.length; i++) {
    if (b[i] !== undefined) {
      cba = b[i];
    }
    if (a[i] !== cba) {
      res += a[i];
    }
  }
  return res.split("").map(Number);
}

console.log(arrayDiff([1, 2], [1]), [2], "a was [1,2], b was [1]");
console.log(arrayDiff([1, 2, 2], [1]), [2, 2], "a was [1,2,2], b was [1]");
console.log(arrayDiff([1, 2, 2], [2]), [1], "a was [1,2,2], b was [2]");
console.log(arrayDiff([1, 2, 2], []), [1, 2, 2], "a was [1,2,2], b was []");
console.log(arrayDiff([], [1, 2]), [], "a was [], b was [1,2]");
console.log(arrayDiff([1, 2, 3], [1, 2]), [3], "a was [1,2,3], b was [1,2]");
