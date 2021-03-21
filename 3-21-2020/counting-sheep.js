function countSheeps(arrayOfSheep) {
  console.log(arrayOfSheep);
  let jumlahTotal = 0;
  for (let i = 0; i < arrayOfSheep.length; i++) {
    if (arrayOfSheep[i] === true) {
      console.log(arrayOfSheep[i]);
      console.log([i]);
      jumlahTotal += arrayOfSheep[i];
      console.log(jumlahTotal);
    }
  }
  return jumlahTotal;
}

let array = [
  true,
  true,
  true,
  false,
  true,
  true,
  true,
  true,
  true,
  false,
  true,
  false,
  true,
  false,
  false,
  true,
  true,
  true,
  true,
  true,
  false,
  false,
  true,
  true,
];

console.log(countSheeps(array));
