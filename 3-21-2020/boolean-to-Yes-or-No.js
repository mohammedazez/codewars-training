function boolToWord(bool) {
  console.log(bool);
  if (bool === true) {
    return "Yes";
  }

  if (bool === false) {
    return "No";
  }
  return bool;
}

console.log(boolToWord(true), "Yes");
console.log(boolToWord(false), "No");
