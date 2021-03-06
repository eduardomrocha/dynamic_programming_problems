"""How sum.

This script implements a solution to the how sum problem using tabulation.

Given an integer target_sum and a list of numbers as arguments, return a list containing
any combination of elements that add up to exactly the target_sum. If there is no
combination that adds up to the target_sum, then return None.
"""
from typing import List, Union


def how_sum(target_sum: int, numbers: List[int]) -> Union[List[int], None]:
    """Find a combination of elements from numbers that add up to exactly the target_sum.

    Args:
        target_sum (int): Value to be generated.
        numbers (List[int]): List of numbers to be used to generate the target sum.

    Returns:
        Union[List[int], None]: A list containing any combination of elements that add
        up to exactly the target_sum, or None in case there is no combination that adds
        up.
    """
    table: List = [None for _ in range(target_sum + 1)]
    table[0] = []

    for i in range(target_sum):
        if table[i] is not None:
            for number in numbers:
                if (i + number) <= target_sum:
                    table[i + number] = table[i] + [number]

    return table[target_sum]


if __name__ == "__main__":
    print(how_sum(7, [2, 3]))  # [3,2,2]
    print(how_sum(7, [5, 3, 4, 7]))  # [4,3]
    print(how_sum(7, [2, 4]))  # None
    print(how_sum(8, [2, 3, 5]))  # [2,2,2,2]
    print(how_sum(300, [7, 14]))  # None
