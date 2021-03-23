var stringToNumber = function (str) {
  if (typeof str == "string") {
    return parseInt(str);
  }
  return result;
};

console.log(stringToNumber("1234"), 1234);
console.log(stringToNumber("605"), 605);
console.log(stringToNumber("1405"), 1405);
console.log(stringToNumber("-7"), -7);
