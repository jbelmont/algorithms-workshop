"use strict";

const assert = require('assert');

const quicksort = require('./quicksort').quicksort;

function testQuickSort() {
    let numbers = [7, 2, 1, 6, 8, 5, 3, 4];
    const LOW = 0;
    const HIGH = numbers.length - 1;
    quicksort(numbers, LOW, HIGH);
    let expected = [1, 2, 3, 4, 5, 6, 7, 8];
    assert.deepEqual(numbers, expected, "The numbers should be sorted");
}

testQuickSort();