"""How sum.

This script implements a solution to the how sum problem using memoization.

Given an integer target_sum and a list of numbers as arguments, return a list containing
any combination of elements that add up to exactly the target_sum. If there is no
combination that adds up to the target_sum, then return None.
"""
from typing import List, Union


def how_sum(
    target_sum: int, numbers: List[int], memo: dict = None
) -> Union[List[int], None]:
    """Find a combination of elements from numbers that add up to exactly the target_sum.

    Args:
        target_sum (int): Value to be generated.
        numbers (List[int]): List of numbers to be used to generate the target sum.
        memo (dict): A hashmap to cache combinations already calculated.

    Returns:
        Union[List[int], None]: A list containing any combination of elements that add
        up to exactly the target_sum, or None in case there is no combination that adds
        up.
    """
    if memo is None:
        memo = {}

    if target_sum in memo:
        return memo[target_sum]

    if target_sum == 0:
        return []

    if target_sum < 0:
        return None

    for number in numbers:
        remainder = target_sum - number
        result = how_sum(remainder, numbers, memo)
        if result is not None:
            result.append(number)
            memo[target_sum] = result
            return result

    memo[target_sum] = None
    return None


if __name__ == "__main__":
    print(how_sum(7, [2, 3]))  # [3,2,2]
    print(how_sum(7, [5, 3, 4, 7]))  # [4,3]
    print(how_sum(7, [2, 4]))  # None
    print(how_sum(8, [2, 3, 5]))  # [2,2,2,2]
    print(how_sum(300, [7, 14]))  # None
