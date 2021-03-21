function lovefunc(flower1, flower2) {
  console.log(flower1);
  console.log(flower2);

  // jika flower 2 genap dan flower 1 ganjil maka return true
  if (flower2 % 2 == 0 && flower1 % 2 == 1) {
    return true;
  }
  //   jika flower 1 genap dan flower 2 ganjil maka return true
  if (flower1 % 2 == 0 && flower2 % 2 == 1) {
    return true;
  }
  //   jika flower 2 genap dan flower 1 genap maka return false
  if (flower2 % 2 == 0 && flower1 % 2 == 0) {
    return false;
  }

  //   jika flower 2 ganjil dan flower 1 ganjil maka return false
  if (flower2 % 2 == 1 && flower1 % 2 == 1) {
    return false;
  }
}

console.log(lovefunc(1, 4), true);
console.log(lovefunc(2, 2), false);
console.log(lovefunc(0, 1), true);
console.log(lovefunc(0, 0), false);
