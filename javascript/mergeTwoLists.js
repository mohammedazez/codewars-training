/** 

* [1,2,4] [1,3,4]
* [1, 2, 4, 1, 3, 4] -> Join push
* [1, 1, 2, 3, 4, 4] -> sort

 **/

function mergeTwoLists(list1, list2) {
  for (let i = 0; i < list1.length; i++) {
    list2.push(list1[i]);
    list2.sort((a, b) => a - b);
  }

  return list2;
}

console.log(mergeTwoLists([1, 2, 4], [1, 3, 4]));
