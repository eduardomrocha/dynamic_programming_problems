"""Fibonacci.

This script implements a fibonacci function using tabulation.

The fibonacci function takes in a number as an argument. The function should return the
n-th number of the Fibonacci sequence.

The 0th number of the sequence is 0 and the 1st number of the sequence is 1.

To generate the next number of the sequence, we sum the previous two.
"""


def fibonacci(n: int) -> int:
    """Given a integer n, return the n-th number of the fibonacci sequence.

    Args:
        n (int): The index of a number in the fibonacci sequence.

    Returns:
        int: The n-th number of the fibonacci sequence.
    """
    table = [0] * (n + 1)
    table[1] = 1

    for i in range(len(table)):
        if i + 1 < len(table):
            table[i + 1] += table[i]
            if i + 2 < len(table):
                table[i + 2] += table[i]

    return table[n]


if __name__ == "__main__":
    print(fibonacci(6))  # 8
    print(fibonacci(7))  # 13
    print(fibonacci(8))  # 21
    print(fibonacci(50))  # 12586269025
