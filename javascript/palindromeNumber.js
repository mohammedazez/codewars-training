// 122 = 221 --> false
// 121 = 121 --> true
function isPalindrome(x) {
  if (x.toString() === x.toString().split("").reverse().join("")) {
    return true;
  } else {
    return false;
  }
}

console.log(isPalindrome(-121));
