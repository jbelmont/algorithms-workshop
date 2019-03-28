"use strict";

// temporary variable swap technique in JavaScript
function swap(numbers, i, j) {
    let temp = numbers[i];
    numbers[i] = numbers[j];
    numbers[j] = temp;
 }

 function bubblesort(numbers, numOfElements) {
   let swapTotal = 0;

   for (let i = 0; i < numOfElements; i++) {

    let numberOfSwaps = 0;

    for (let j = 0; j < numOfElements-1; j++) {
      if (numbers[j] > numbers[j+1]) {
        swap(numbers, j, j+1);
        numberOfSwaps++;
      }
    }

    swapTotal += numberOfSwaps;

    // If no elements were swapped during a traversal, array is sorted
    if (numberOfSwaps === 0) {
      break;
    }
   }

   return swapTotal;
 }

exports.bubblesort = bubblesort;
