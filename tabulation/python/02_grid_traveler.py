"""Grid traveler.

This script implements a solution to the grid traveler problem using tabulation.

Say that you are a traveler on a 2D grid. You begin in the top-left corner and your goal
is to travel to the bottom-right corner. You may only move down or right.

In how many ways can you travel to the goal on a grid with dimensions m * n?
"""


def grid_traveler(m: int, n: int) -> int:
    """Grid traveler.

    Given the number of lines and columns, return the number of ways the traveler can
    travel to the goal.

    Args:
        m (int): Number of lines.
        n (int): Number of columns.

    Returns:
        int: The number of ways the traveler can travel to the goal.
    """
    table = [[0 for _ in range(n + 1)] for _ in range(m + 1)]
    table[1][1] = 1

    for i in range(m + 1):
        for j in range(n + 1):
            if j + 1 <= n:
                table[i][j + 1] += table[i][j]
            if i + 1 <= m:
                table[i + 1][j] += table[i][j]

    return table[m][n]


if __name__ == "__main__":
    print(grid_traveler(1, 1))  # 1
    print(grid_traveler(2, 3))  # 3
    print(grid_traveler(3, 2))  # 3
    print(grid_traveler(3, 3))  # 6
    print(grid_traveler(18, 18))  # 2333606220
