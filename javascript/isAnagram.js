function isAnagram(s, t) {
  if (s.length !== t.length) {
    return false;
  }
  const countS = {};
  const countT = {};

  for (let i = 0; i < s.length; i++) {
    countS[s[i]] = 1 + (countS[s[i]] || 0);
    countT[t[i]] = 1 + (countT[t[i]] || 0);
  }

  const keysS = Object.keys(countS);

  for (let i = 0; i < keysS.length; i++) {
    const key = keysS[i];
    if (countS[key] !== countT[key]) {
      return false;
    }
  }

  return true;
}

console.log(isAnagram("anagram", "nagaram"));
