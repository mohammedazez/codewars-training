function solution(str) {
  console.log(str);
  let word = "";
  console.log(word);
  for (let i = str.length - 1; i >= 0; i--) {
    word += str[i];
    console.log(str[i]);
    console.log(word);
  }
  return word;
}

console.log(solution("world"), "dlrow");
// console.log(solution("hello"), "olleh");
// console.log(solution(""), "");
// console.log(solution("h"), "h");
