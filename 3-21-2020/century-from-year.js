function century(year) {
  console.log(year);
  let rumus = 100;
  let result = year / rumus;
  let bulatkan = Math.ceil(result);
  return bulatkan;
}

console.log(century(1705), 18, "Testing for year 1705");
console.log(century(1900), 19, "Testing for year 1900");
console.log(century(1601), 17, "Testing for year 1601");
console.log(century(2000), 20, "Testing for year 2000");
console.log(century(89), 1, "Testing for year 89");
