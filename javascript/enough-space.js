function enough(cap, on, wait) {
  let kurangi = cap - (on + wait);
  console.log(kurangi);

  if (kurangi >= 0) {
    return 0;
  }
  if (kurangi < 0) {
    return Math.abs(kurangi);
  }
  return kurangi;
}

console.log(enough(10, 5, 5), 0);
console.log(enough(100, 60, 50), 10);
console.log(enough(20, 5, 5), 0);
