// "Hey fellow warriors" -> "Hey wollef sroirraw"
// ["Hey", "fellow", "warriors"]
// "Hey" "fellow" "warriors"
// lebih dari 5 di balik "fellow" -> "wollef"  "sroirraw""
// ["Hey", "wollef", "sroirraw"]
// "Hey wollef sroirraw"

function spinWords(str) {
  let split = str.split(" ");
  let res = [];
  for (let i = 0; i < split.length; i++) {
    if (split[i].length >= 5) {
      res.push(split[i].split("").reverse().join(""));
    } else {
      res.push(split[i]);
    }
  }
  return res.join(" ");
}

console.log(spinWords("Hey fellow warriors"));
