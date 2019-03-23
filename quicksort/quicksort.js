"use strict";

// temporary variable swap technique in JavaScript
function swap(numbers, i, j) {
    let temp = numbers[i];
    numbers[i] = numbers[j];
    numbers[j] = temp;
 }

function partition(numbers, low, high) {
    let i = low;
    let pivot = numbers[high];
    for (let j = low; j < high; j++) {
        if (numbers[j] < pivot) {
            swap(numbers, i, j);
            i++
        }
    }
    swap(numbers, i, high);
	return i
}

function quicksort(numbers, low, high) {
	if (low < high) {
		let p = partition(numbers, low, high)
		quicksort(numbers, low, p-1)
		quicksort(numbers, p+1, high)
    }
    return numbers;
}

exports.quicksort = quicksort;