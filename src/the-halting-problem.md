# Algorithms Workshop - The Halting Problem

## Halting Problem Defined

[Halting Problem (Wikipedia Definition)](https://en.wikipedia.org/wiki/Halting_problem)

> In computability theory, the halting problem is the problem of determining, from a description of an arbitrary computer program and an input, whether the program will finish running, or continue to run forever. Alan Turing proved in 1936 that a general algorithm to solve the halting problem for all possible program-input pairs cannot exist. For any program f that might determine if programs halt, a "pathological" program g called with an input can pass its own source and its input to f and then specifically do the opposite of what f predicts g will do. No f can exist that handles this case. A key part of the proof was a mathematical definition of a computer and program, which became known as a Turing machine; the halting problem is undecidable over Turing machines. Turing's proof is one of the first cases of decision problems to be concluded. The theoretical conclusion that it is not solvable is significant to practical computing efforts, defining a class of applications which no programming invention can possibly perform perfectly.

The Halting problem is important because it helps show the difficulty of a given algorithm.

It pretty much shows that sometimes you can only guess at a problem and cannot come up with an actual way to solve the problem.

## Historical Background

[Halting (Wikipedia History)](https://en.wikipedia.org/wiki/Halting_problem#History)

> The halting problem is historically important because it was one of the first problems to be proved undecidable. (Turing's proof went to press in May 1936, whereas Alonzo Church's proof of the undecidability of a problem in the lambda calculus had already been published in April 1936 [Church, 1936].) Subsequently, many other undecidable problems have been described.

## Proof by Contradiction

Alan Turing proved the Halting Problem through Contradiction.

## What is an example of a Halting Problem

Lets say you write a computer program and the program doesn't stop and just hangs.

So you ask yourself will this program *Halt* or stop or will it continue to run forever.