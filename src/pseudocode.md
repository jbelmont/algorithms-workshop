# Algorithms Workshop - Pseudocode

## Sections:

* [Pseudocode Definition](#pseudocode-definition)

## Pseudocode Definition

[Definition of Pseudocode (Wikipedia)](https://en.wikipedia.org/wiki/Pseudocode)

> In computer science, pseudocode is an informal high-level description of the operating principle of a computer program or other algorithm. It uses the structural conventions of a normal programming language, but is intended for human reading rather than machine reading. Pseudocode typically omits details that are essential for machine understanding of the algorithm, such as variable declarations, system-specific code and some subroutines. The programming language is augmented with natural language description details, where convenient, or with compact mathematical notation. The purpose of using pseudocode is that it is easier for people to understand than conventional programming language code, and that it is an efficient and environment-independent description of the key principles of an algorithm. It is commonly used in textbooks and scientific publications that are documenting various algorithms, and also in planning of computer program development, for sketching out the structure of the program before the actual coding takes place.

The pseudocode that you use can be a more formal syntax which is known as a *Pidgin* code.

Pidgin code is often used by more mathematically inclined individuals but is not strictly necessary.

I have personally seen people write *Pseudocode* like english language and some who write pseudocode in their day to day programming language.

#### Pseudocode Example

```
algorithm quicksort(A, lo, hi) is
    if lo < hi then
        p := partition(A, lo, hi)
        quicksort(A, lo, p - 1)
        quicksort(A, p + 1, hi)

algorithm partition(A, lo, hi) is
    pivot := A[hi]
    i := lo
    for j := lo to hi do
        if A[j] < pivot then
            swap A[i] with A[j]
            i := i + 1
    swap A[i] with A[hi]
    return i
```

#### Pidgin Code Example

This is using a mathematical style symbols:

\\( \int x dx = \frac{x^2}{2} + C \\)
