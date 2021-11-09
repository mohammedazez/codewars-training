function repeatStr(n, s) {
  console.log(n);
  console.log(s);
  let tampung = "";
  console.log(tampung);

  for (let i = 0; i < n; i++) {
    console.log(n);
    console.log(s);
    tampung += s;
  }

  return tampung;
}

console.log(repeatStr(4, "*"));
