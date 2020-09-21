# Algorithms Workshop - Data Structures

## Definition of a Data Structure

[Data Structure (Wikipedia Definition)](https://en.wikipedia.org/wiki/Data_structure)

> In computer science, a data structure is a data organization, management, and storage format that enables efficient access and modification. More precisely, a data structure is a collection of data values, the relationships among them, and the functions or operations that can be applied to the data.

## Types of Data Structures

#### Linear Data Structures

A data structure is linear if its elements form a sequence:

* Lists
* Sets
* Tuples
* Queues
* Stacks

## NonLinear Data Structures

A data structure is non-linear when an element is connected to many other elements

* Trees
* Tables
* Hash
* Containers

## Set Operations

[Set (Wikipedia Definition)](https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29)

> In computer science, a set is an abstract data type that can store unique values, without any particular order. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.

#### Union 

The union of s and t returns all the elements of set s and set t

#### Intersection

The intersection of s and t return the all the elements that belong to A and to B

#### Add

Add is a method that adds an element to the set

#### Remove

Remove is a method that removes an element from the set

#### Set Implementation

[Set Implementation](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/set/set.go)

#### Set Test Run

[Set Tests](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/set/set_test.go)

## Stack Operations

[Stack (Wikipedia Definition)](https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29)

> In computer science, a stack is an abstract data type that serves as a collection of elements, with two main principal operations:
>     push, which adds an element to the collection, and
>     pop, which removes the most recently added element that was not yet removed.

#### Push 

Add an element to the stack

#### Pop 

Remove most recent element from the stack and return it.

#### Stack Implementation

[Stack Implementation](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/stack/stack.go)

#### List Test Run

[Stack Tests](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/stack/stack_test.go)

## List Operations

[List (Wikipedia Definition)](https://en.wikipedia.org/wiki/List_%28abstract_data_type%29)

> In computer science, a list or sequence is an abstract data type that represents a countable number of ordered values, where the same value may occur more than once. An instance of a list is a computer representation of the mathematical concept of a tuple or finite sequence; the (potentially) infinite analog of a list is a stream. Lists are a basic example of containers, as they contain other values. If the same value occurs multiple times, each occurrence is considered a distinct item.

#### Append

Operation to append an element to a list

#### Prepend

Operation to prepend an element to a list

#### Head

Operation to find the first element of a list

#### Empty

Operation to determine whether or not a list is empty

#### List Implementation

[List Implementation](https://github.com/golang/go/blob/go1.15.2/src/container/list/list.go)

#### List Test Run

[List Tests](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/list/list_test.go)