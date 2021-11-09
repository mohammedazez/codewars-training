function findNeedle(haystack) {
  for (let i = 0; i < haystack.length; i++) {
    console.log(haystack[i]);
    if (haystack[i] == "needle") {
      return `found the needle at position ${i}`;
    }
  }
}

const haystack_2 = [
  "283497238987234",
  "a dog",
  "a cat",
  "some random junk",
  "a piece of hay",
  "needle",
  "something somebody lost a while ago",
];

console.log(findNeedle(haystack_2), "found the needle at position 5");
