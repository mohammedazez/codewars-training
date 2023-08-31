// ulang pangil function recursive
// 770843 -> first
// [7 , 7 , 0, 8, 4, 3] -> split
// 7 + 7 + 0 + 8 + 4 + 3 = 29 -> reduce

// ulang pangil function recursive
// 29 -> result
// [2, 9] -> split
// 2 + 9 -> reduce

// ulang pangil function recursive
// 11 -> result
// [1, 1] -> split
// 1 + 1 -> reduce
// 2
function digital_root(n) {
  if (n < 10) {
    return "ini" + " " + n;
  }
  return digital_root(
    String(n)
      .split("")
      .reduce((s, v) => Number(s) + Number(v))
  );
}

console.log(digital_root(493193));

function digital_root(n) {
  if (n < 10) {
    return "ini" + " " + n;
  }
  return digital_root(
    String(n)
      .split("")
      .reduce((s, v) => Number(s) + Number(v))
  );
}

console.log(digital_root(493193));

// 770843  -->  7 + 7 + 0 + 8 + 4 + 3 = 29  -->  2 + 9 = 11

// 770843 > length lebih dari 1
// [7 , 7 , 0, 8, 4, 3] -> split
// 7 + 7 + 0 + 8 + 4 + 3 = 29 -> reduce

// 29 > length lebih dari 1
// [2, 9] -> split
// 2 + 9 -> reduce

// 11 > length lebih dari 1
// [1, 1] -> split
// 1 + 1 -> reduce
// 2

function digitalRoot(n) {
  let res = 0;
  let split = n
    .toString()
    .split("")
    .map(Number)
    .reduce((acc, cur) => acc + cur);
  res = split;

  if (res.toString().length > 1) {
    let t = res
      .toString()
      .split("")
      .map(Number)
      .reduce((acc, cur) => acc + cur);
    res = t;
  }

  if (res.toString().length > 1) {
    let c = res
      .toString()
      .split("")
      .map(Number)
      .reduce((acc, cur) => acc + cur);
    res = c;
  }
  return res;
}

console.log(digitalRoot(16));
