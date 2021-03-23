function rentalCarCost(d) {
  console.log(d);
  let hargaNormal = 40;
  if (d == 1) {
    return hargaNormal;
  } else if (d == 2) {
    return hargaNormal * 2;
  } else if (d >= 3 && d <= 6) {
    return d * hargaNormal - 20;
  } else if (d >= 7) {
    return d * hargaNormal - 50;
  } else {
    return `tidak ada data`;
  }
}

// tidak ada potongan
console.log(rentalCarCost(1), 40);
console.log(rentalCarCost(2), 80);

// potongan 20 dollar
console.log(rentalCarCost(3), 100);
console.log(rentalCarCost(4), 140);
console.log(rentalCarCost(5), 180);
console.log(rentalCarCost(6), 220);

// potongan 50 dollar
console.log(rentalCarCost(7), 230);
console.log(rentalCarCost(8), 270);
console.log(rentalCarCost(9), 310);
console.log(rentalCarCost(10), 350);
