"""Fibonacci.

This script implements a fibonacci function using memoization.

The fibonacci function takes in a number as an argument. The function should return the
n-th number of the Fibonacci sequence. The first and second number of the sequence is 1.
"""


def fibonacci(n: int, memo: dict = None) -> int:
    """Given a integer n, return the n-th number of the fibonacci sequence.

    Args:
        n (int): The index of a number in the fibonacci sequence.
        memo (dict): A hashmap to cache combinations already calculated.

    Returns:
        int: The n-th number of the fibonacci sequence.
    """
    if memo is None:
        memo = {}

    if n in memo:
        return memo[n]

    if n < 2:
        return n

    memo[n] = fibonacci(n - 1, memo) + fibonacci(n - 2, memo)

    return memo[n]


if __name__ == "__main__":
    print(fibonacci(6))  # 8
    print(fibonacci(7))  # 13
    print(fibonacci(8))  # 21
    print(fibonacci(50))  # 12586269025
