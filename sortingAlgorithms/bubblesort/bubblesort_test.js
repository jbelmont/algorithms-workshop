"use strict";

const assert = require('assert');

const bubblesort = require('./bubblesort').bubblesort;

function testBubbleSort() {
    let numbers = [7, 2, 1, 6, 8, 5, 3, 4];
    bubblesort(numbers, numbers.length);
    let expected = [1, 2, 3, 4, 5, 6, 7, 8];
    assert.deepEqual(numbers, expected, "The numbers should be sorted");
}

testBubbleSort();
