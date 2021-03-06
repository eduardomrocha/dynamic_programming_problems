"""Best sum.

This script implements a solution to the best sum problem using tabulation.

Given an integer target_sum and a list of numbers as arguments, return a list containing
the shortest combination of elements that add up to exactly the target_sum. If there is
a tie for the shortest combination, you may return any one of the shortest.
"""
from typing import List, Union


def best_sum(target_sum: int, numbers: List[int]) -> Union[List[int], None]:
    """Find the best combination of elements that add up to exactly the target_sum.

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
                    combination = table[i] + [number]
                    if table[i + number] is None or len(combination) < len(
                        table[i + number]
                    ):
                        table[i + number] = combination

    return table[target_sum]


if __name__ == "__main__":
    print(best_sum(7, [5, 3, 4, 7]))  # [7]
    print(best_sum(8, [2, 3, 5]))  # [3, 5]
    print(best_sum(8, [1, 4, 5]))  # [4, 4]
    print(best_sum(100, [1, 3, 5, 25]))  # [25, 25, 25, 25]
