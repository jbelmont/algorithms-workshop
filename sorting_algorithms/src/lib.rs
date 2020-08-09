#[allow(dead_code)]
fn bubblesort(numbers: Vec<i32>) -> Vec<i32> {
    let length = numbers.len();
    let mut result: Vec<i32> = numbers.into();
    loop {
        let mut swaps = 0;

        for i in 1..length {
            let first = result[i];
            let second = result[i-1];
            if comparator(second, first) {
                result.swap(i - 1, i);
                swaps += 1;
            }
        }

        // If no elements were swapped during a traversal, array is stored
        if swaps == 0 {
            break
        }
    }

    result
}

fn comparator(x: i32, y: i32) -> bool {
    x > y
}

fn swap(numbers: &mut [i32], i: usize, j: usize) {
    let temp = numbers[i];
    numbers[i] = numbers[j];
    numbers[j] = temp;
}

fn partition(numbers: &mut [i32], low: usize, high: usize) -> i32 {
    let mut i = low;
    let pivot = numbers[high];

    for j in low..high {
        if numbers[j] < pivot {
            swap(numbers, i, j);
            i += 1;
        }
    }
    swap(numbers, i, high);
    i as i32
}

#[allow(dead_code)]
fn quicksort(numbers: &mut [i32], low: usize, high: usize) {
    if low >= high {
        return;
    }
    let p = partition(numbers, low, high);
    quicksort(numbers, low, (p-1) as usize);
    quicksort(numbers, (p+1) as usize, high);
}

#[cfg(test)]
mod tests {
    use super::bubblesort;
    use super::quicksort;

    #[test]
    fn test_bubblesort() {
        let numbers = vec![7, 5, 9, 3];
        let actual = bubblesort(numbers);
        let expected: Vec<i32> =vec![3, 5, 7, 9];
        assert_eq!(actual, expected);
    }

    #[test]
    fn test_quicksort() {
        let mut numbers = vec![1, 57, 98, 193, 300, 55, 987];
        let high = numbers.len() - 1;
        quicksort(&mut numbers, 0, high);
        let expected = vec![1, 55, 57, 98, 193, 300, 987];
        assert_eq!(numbers, expected);
    }
}