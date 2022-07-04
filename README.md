# GoConvenience

This module contains two packages, I first wanted to make it a functional library but soon 
I realized that Go was not Java 8 (generics and fully object-oriented). 
So I dialed down my ambitions to make this into a 
convenience library leveraging what Go has best to offer.

## Convenience
Contains some methods to loop over, map, filter and reduce slices. Was supposed to be something like
Java 8 streams but soon I found out that methods in Go cannot be generic.

## Assertions
Contains methods based on Hamcrest assertions, the file Convenience/interfaces_test.go has examples on
how to use them
