"""Grid traveler.

This script implements a solution to the grid traveler problem using memoization.

Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal
is to travel to the bottom-right corner. You may only move down or right.

In how many ways can you travel to the goal on a grid with dimensions m * n?
"""


def grid_traveler(m: int, n: int, memo: dict = None) -> int:
    """Grid traveler.

    Given the number of lines and columns, return the number of ways the traveler can
    travel to the goal.

    Args:
        m (int): Number of lines.
        n (int): Number of columns.
        memo (dict): A hashmap to cache combinations already calculated.

    Returns:
        int: The number of ways the traveler can travel to the goal.
    """
    if memo is None:
        memo = {}

    key = str(m) + "," + str(n)

    if key in memo:
        return memo[key]

    if m == 1 and n == 1:
        return 1

    if m == 0 or n == 0:
        return 0

    memo[key] = grid_traveler(m - 1, n, memo) + grid_traveler(m, n - 1, memo)
    return memo[key]


if __name__ == "__main__":
    print(grid_traveler(1, 1))  # 1
    print(grid_traveler(2, 3))  # 3
    print(grid_traveler(3, 2))  # 3
    print(grid_traveler(3, 3))  # 6
    print(grid_traveler(18, 18))  # 2333606220
