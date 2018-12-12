## Introduction to Algorithm Design

#### Algorithm

An algorithm is a procedure to accomplish a specific task.
An algorithm is the idea behind any reasonable computer program.

To be interesting, an algorithm must solve a general, well-specified problem. An
algorithmic problem is specified by describing the complete set of instances it must
work on and of its output after running on one of these instances.

For example, the
algorithmic problem known as sorting is defined as follows:
```
Problem: Sorting
Input: A sequence of n keys a1, . . . , an.
Output: The permutation (reordering) of the input sequence such that a'1 ≤ a'2 ≤ ··· ≤ a'n−1 ≤ a'n.
```
Determining that
you are dealing with a general problem is your first step towards solving it.


Three desirable properties for a good algorithm:
*correct and efficient while being easy to implement*

#### Algorithm Correctness

#### Robot Tour optimization
```
Problem: Robot Tour Optimization
Input: A set S of n points in the plane.
Output: What is the shortest cycle tour that visits each point in the set S?
```

Several algorithms might come to mind to solve this problem. Perhaps the most
popular idea is the *nearest-neighbor heuristic*. Starting from some point p0, we walk
first to its nearest neighbor p1. From p1, we walk to its nearest unvisited neighbor,
thus excluding only p0 as a candidate. We now repeat this process until we run
out of unvisited points, after which we return to p0 to close off the tour.

```
NearestNeighbor(P)
    Pick and visit an initial point p0 from P
    p = p0
    i = 0
    While there are still unvisited points
        i = i + 1
        Select pi to be the closest unvisited point to pi−1
        Visit pi
    Return to p0 from pn−1
```
It is simple to understand and implement.
It makes sense to visit nearby points before we visit faraway points to reduce
the total travel time.
Against all these positives there is only one problem. This algorithm is completely
wrong.

How can it be wrong? The algorithm always finds a tour, but it doesn’t
necessarily find the shortest possible tour.

Always walking to the closest
point is too restrictive, since it seems to trap us into making moves we didn’t
want.

A different idea might be to repeatedly connect the closest pair of endpoints
whose connection will not create a problem, such as premature termination of the
cycle.

Each vertex begins as its own single vertex chain. After merging everything
together, we will end up with a single chain containing all the points in it. Connecting the final two endpoints gives us a cycle. At any step during the execution
of this closest-pair heuristic, we will have a set of single vertices and vertex-disjoint chains available to merge. In pseudocode:

```
ClosestPair(P)
    Let n be the number of points in set P.
    For i = 1 to n − 1 do
        d = ∞
        For each pair of endpoints (s, t) from distinct vertex chains
            if dist(s, t) ≤ d then sm = s, tm = t, and d = dist(s, t)
        Connect (sm, tm) by an edge
    Connect the two endpoints by an edge
```

The closest-pair heuristic is somewhat more complicated and less efficient
than the previous one, but at least it gives the right answer in some examples.
But this is not true in all examples.

Thus this second algorithm is also wrong. Which one of these algorithms performs
better? You can’t tell just by looking at them. Clearly, both heuristics can
end up with very bad tours on very innocent-looking input.
At this point, you might wonder what a correct algorithm for our problem looks
like. Well, we could try enumerating all possible orderings of the set of points, and
then select the ordering that minimizes the total length:

```
OptimalTSP(P)
    d = ∞
    For each of the n! permutations Pi of point set P
        If (cost(Pi) ≤ d) then d = cost(Pi) and Pmin = Pi
    Return Pmin
```

Since all possible orderings are considered, we are guaranteed to end up with
the shortest possible tour. This algorithm is correct, since we pick the best of all
the possibilities. But it is also extremely slow. The fastest computer in the world
couldn’t hope to enumerate all the 20! =2,432,902,008,176,640,000 orderings of 20
points within a day. For real circuit boards, where n ≈ 1, 000, forget about it.

The quest for an efficient algorithm to solve this problem, called the traveling
salesman problem (TSP), will take us through much of this book.

> Take-Home Lesson: There is a fundamental difference between algorithms,
which always produce a correct result, and heuristics, which may usually do a
good job but without providing any guarantee.
