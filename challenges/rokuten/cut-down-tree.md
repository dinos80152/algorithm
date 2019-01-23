# Cut down tree

## Problem

There are N trees growing along a street. For safety reasons, one of the trees needs to be cut down. The local council wants the street to look ordered, so all the remaining trees should be sorted in non-decreasing order of height. Your goal is to count the number of ways this can be done.

Write a function:

```java
class Solution
{
    public int solution(int[] A);
}
```

that, given an array A consisting of N integers, where A[K] denotes the height of the K-th tree, returns the number of trees you can cut to satisfy the above condition. If the condition cannot be satisfied, your function should return 0.

Examples:

1. Given A = [3, 4, 5, 4] your function should return 2. You can cut down the tree of height 5 or second tree of height 4:

    ![ordered_trees01](ordered_trees01.png)
    ![ordered_trees02](ordered_trees02.png)

1. Given A = [4, 5, 2, 3, 4], your function should return 0. After cutting down any of these trees, the rest of the trees will not be ordered by height.

1. Given A = [1, 2, ..., 100,000] (consecutive numbers from 1 to 100,000), your function should return 100,000. You can cut down any of these trees.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000].

## Reference

* Rokuten Japan Codility Exam