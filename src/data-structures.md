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

## Dynamic Data Structures

* Dictionaries
* TreeSets
* Sequences

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

#### Doubly LinkedList Implementation

[Doubly LinkedList Implementation](https://github.com/golang/go/blob/go1.15.2/src/container/list/list.go)

#### Singly LinkedList Implementation

[Singly LinkedList Implementation](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/linkedlist/linkedlist.go)

#### Singly List Test Run

[List Tests](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/list/list_test.go)

## Queue Operations

[Queue (Wikipedia Definition)](https://en.wikipedia.org/wiki/Queue_%28abstract_data_type%29)

> In computer science, a queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence. By convention, the end of the sequence at which elements are added is called the back, tail, or rear of the queue, and the end at which elements are removed is called the head or front of the queue, analogously to the words used when people line up to wait for goods or services.

#### New

Should return a new queue

#### Enqueue 

Operation should add item to the queue

#### Dequeue

Operation should remove and return first item added to the queue in FIFO order

#### Queue Implementation

[Queue Implementation](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/queue/queue.go)

#### Queue Test Run

[Queue Test Run](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/queue/queue_test.go)

## Tree Operations

[Tree (Wikipedia Definition)](https://en.wikipedia.org/wiki/Tree_%28data_structure%29)

> In computer science, a tree is a widely used abstract data type that simulates a hierarchical tree structure, with a root value and subtrees of children with a parent node, represented as a set of linked nodes.

> A tree data structure can be defined recursively as a collection of nodes (starting at a root node), where each node is a data structure consisting of a value, together with a list of references to nodes (the "children"), with the constraints that no reference is duplicated, and none points to the root.

## Table Operations

[Table (Wikipedia Definition)](https://en.wikipedia.org/wiki/Table_%28information%29)

> A table is an arrangement of data in rows and columns, or possibly in a more complex structure. Tables are widely used in communication, research, and data analysis. Tables appear in print media, handwritten notes, computer software, architectural ornamentation, traffic signs, and many other places. The precise conventions and terminology for describing tables vary depending on the context. Further, tables differ significantly in variety, structure, flexibility, notation, representation and use. In books and technical articles, tables are typically presented apart from the main text in numbered and captioned floating blocks.

A Table has rows and columns

#### Table Implementation

[Table Implementation](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/table/table.go)

#### Table Test Run

[Table Test Run](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/table/table_test.go)

## Dictionary Operations

A dictionary can be thought of as a set of key, value pairs. Dictionaries are often used in stoaring a set of data items.

#### Dictionary Implementation

[Dictionary Implementation](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/dictionary/dictionary.go)

#### Dictionary Test Run

[Dictionary Test Run](https://github.com/jbelmont/algorithms-workshop/blob/master/dataStructures/dictionary/dictionary_test.go)