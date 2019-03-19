function sort(arr, l, r) {
  if (l >= r) {
    return;
  }

  let mid = arr[l];

  let i = l;
  let j = r;
  let i_empty = true; // true -> i, false -> j

  while (i < j) {
    if (i_empty) {
      if (arr[j] < mid) {
        arr[i] = arr[j];
        i_empty = false;
        i++;
      } else {
        j--;
      }
    } else {
      if (arr[i] > mid) {
        arr[j] = arr[i];
        i_empty = true;
        j--;
      } else {
        i++;
      }
    }
  }
  arr[i] = mid;

  sort(arr, l, i - 1);
  sort(arr, i + 1, r);

  return;
}

// const ARR = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
const ARR = [10, 5, 9, 4, 8, 3, 7, 2, 6, 0, 0, 0];
sort(ARR, 0, ARR.length - 1);

console.log(ARR);
