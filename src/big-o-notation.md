# Algorithms Workshop - Big-O Notation

## What is Big-O Notation

> In computer science, big O notation is used to classify algorithms according to how their run time or space requirements grow as the input size grows. In analytic number theory, big O notation is often used to express a bound on the difference between an arithmetical function and a better understood approximation; a famous example of such a difference is the remainder term in the prime number theorem.

When talking about data structures and algorithms you will frequently encounter Big O Notation for a given algorithm.

## Big-O, Little-O, Theta, and Omega

When talking about algorithms Big O, Little o, Omega and Theta are formal notational methods that state the growth of an algorithm in terms of efficiency and storage.

#### Big-O notation

Big-O notation is denoted by \\( \mathcal{O}(n) \\). 

#### Little-O notation

Little-O notation is denoted by \\( o(n) \\)

#### Theta 

Theta notation is denoted by \\( \Theta(n) \\)

#### Omega

Omega notation is denoted by \\( \Omega(n) \\)

## Difference between Big-O, Little-O, Theta and Omega

The differences between Big-O, Little-O, Theta, and Omega center around constraints and growth rates.

#### Informal explanation for Big-O, Little-O, Theta and Omega

Big-O: \\( T(n) \\; is \\; \mathcal{O}(f(n)) \\; \rightarrow \\) f(n) describes upper bound for T(n) 

Little-O: \\( T(n) \\; is \\; o(f(n)) \\; \rightarrow \\) f(n) describes the lower bound for T(n)

Theta: \\( T(n) \\; is \\; \Theta(f(n)) \\; \rightarrow \\) f(n) describes the exact bound for T(n)

Omega: \\( T(n) \\; is \\; \Omega(f(n)) \\; \rightarrow \\) f(n) is the upper bound for T(n) but T(n) can never be equal to f(n)