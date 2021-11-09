function removeChar(str) {
  console.log(str);
  let newKata = "";
  console.log(newKata);

  for (let i = 1; i < str.length - 1; i++) {
    console.log(str[i]);
    newKata += str[i];
    console.log(newKata);
  }

  return newKata;
}

console.log(removeChar("eloquent"));
console.log(removeChar("country"));
console.log(removeChar("person"));
console.log(removeChar("place"));
