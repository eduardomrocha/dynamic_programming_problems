"""Best sum.

This script implements a solution to the best sum problem using memoization.

Given an integer target_sum and a list of numbers as arguments, return a list containing
the shortest combination of elements that add up to exactly the target_sum. If there is
a tie for the shortest combination, you may return any one of the shortest.
"""
from copy import deepcopy
from typing import List, Union


def best_sum(
    target_sum: int, numbers: List[int], memo: dict = None
) -> Union[List[int], None]:
    """Find the best combination of elements that add up to exactly the target_sum.

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

    shortest_combination = None

    for number in numbers:
        remainder = target_sum - number
        remainder_combination = best_sum(remainder, numbers, memo)
        if remainder_combination is not None:
            [*combination] = remainder_combination
            combination.append(number)

            if shortest_combination is None or len(combination) < len(
                shortest_combination
            ):
                shortest_combination = combination

    memo[target_sum] = deepcopy(shortest_combination)
    return shortest_combination


if __name__ == "__main__":
    print(best_sum(7, [5, 3, 4, 7]))  # [7]
    print(best_sum(8, [2, 3, 5]))  # [3, 5]
    print(best_sum(8, [1, 4, 5]))  # [4, 4]
    print(best_sum(100, [1, 3, 5, 25]))  # [25, 25, 25, 25]
