// []                                -->  "no one likes this"
// ["Peter"]                         -->  "Peter likes this"
// ["Jacob", "Alex"]                 -->  "Jacob and Alex like this"
// ["Max", "John", "Mark"]           -->  "Max, John and Mark like this"
// ["Alex", "Jacob", "Mark", "Max"]  -->  "Alex, Jacob and 2 others like this"

// index 0     -> "no one likes this"
// index 1        -> 0 + " likes this"
// index >= 2     -> 0 + "and" + 1 +" likes this"
// index >= 3     -> 0 + 1 + and + 2 + " likes this"
// index >= 4     -> 0 + 1 + and + len(index) + "others" + " likes this"

function likes(names) {
  let res = "";
  if (names.length == 0) {
    res = `no one likes this`;
  }
  if (names.length == 1) {
    res = `${names} likes this`;
  }
  if (names.length == 2) {
    res = `${names[0]} and ${names[1]} like this`;
  }
  if (names.length == 3) {
    res = `${names[0]}, ${names[1]} and ${names[2]} like this`;
  }
  if (names.length >= 4) {
    let fit = names.filter((n, i) => console.log(n));
    res = `${names[0]}, ${names[1]} and ${fit.length} others like this`;
  }

  return res;
}

console.log(likes([]), "no one likes this");
console.log(likes(["Peter"]), "Peter likes this");
console.log(likes(["Jacob", "Alex"]), "Jacob and Alex like this");
console.log(likes(["Max", "John", "Mark"]), "Max, John and Mark like this");
console.log(
  likes(["Alex", "Jacob", "Mark", "Max"]),
  "Alex, Jacob and 2 others like this"
);
