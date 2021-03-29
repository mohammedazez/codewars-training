class SmallestIntegerFinder {
  findSmallestInt(args) {
    console.log(args);
    let find = Math.min(...args);
    console.log(find);
    return find;
  }
}

var sif = new SmallestIntegerFinder();
console.log(sif.findSmallestInt([78, 56, 232, 12, 8]));
console.log(sif.findSmallestInt([78, 56, 232, 12, 18]));
console.log(sif.findSmallestInt([78, 56, 232, 412, 228]));
console.log(sif.findSmallestInt([78, 56, 232, 12, 0]));
console.log(sif.findSmallestInt([1, 56, 232, 12, 8]));
