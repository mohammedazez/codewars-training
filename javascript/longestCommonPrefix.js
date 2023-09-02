function longestCommonPrefix(strs) {
  if (!strs || strs.length === 0) {
    return "";
  }

  if (strs.length === 1) {
    return strs[0];
  }

  let minLength = Number.MAX_SAFE_INTEGER;
  for (let i = 0; i < strs.length; i++) {
    minLength = Math.min(strs[i].length, minLength);
  }

  let prefix = "";
  outer: for (let i = 0; i <= minLength; i++) {
    const tmp = strs[0].substring(0, i);
    for (let j = 1; j < strs.length; j++) {
      if (!strs[j].startsWith(tmp)) {
        break outer;
      }
    }
    prefix = tmp;
  }

  return prefix;
}
