"""Can sum.

This script implements a solution to the can sum problem using tabulation.

Given an integer target_sum and a list of numbers as arguments, return a boolean
indicating whether or not it is possible to generate the target_sum using numbers from
the list, using an element of the list as many times as needed and assuming that all
input numbers are nonnegative.
"""
from typing import List


def can_sum(target_sum: int, numbers: List[int]) -> bool:
    """Check if the target_sum can be gereteded by adding up elements from numbers.

    Args:
        target_sum (int): Value to be generated.
        numbers (List[int]): List of numbers to be used to generate the target sum.

    Returns:
        bool: A boolean indicating whether or not it is possible to geerate the
        target_sum using numbers from the list.
    """
    table = [False] * (target_sum + 1)
    table[0] = True

    for i in range(target_sum):
        if table[i]:
            for number in numbers:
                if (i + number) <= target_sum:
                    table[i + number] = True

    return table[target_sum]


if __name__ == "__main__":
    print(can_sum(7, [2, 3]))  # True
    print(can_sum(7, [5, 3, 4, 7]))  # True
    print(can_sum(7, [2, 4]))  # False
    print(can_sum(8, [2, 3, 5]))  # True
    print(can_sum(300, [7, 14]))  # False
