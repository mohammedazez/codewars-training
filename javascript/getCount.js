// Make new variabel array for vowel = aiueo
// Make new variable array for result
// Looping str to get each value
// Check str contain with vowel compare with variable vowel
// assign str to varibel result
// and return length

function getCount(str) {
  let vowel = ["a", "i", "u", "e", "o"];
  let result = [];

  for (let i = 0; i < str.length; i++) {
    if (vowel.includes(str[i])) {
      result += str[i].split("");
    }
  }
  return result.length;
}

getCount("abracadabra");
