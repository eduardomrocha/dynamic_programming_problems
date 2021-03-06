"""Count construct.

This script implements a solution to the count construct problem using memoization.

Given a string 'target' and a list of strings 'word_bank', return the number of ways
that 'target' can be constructed by concatenating elements of the 'word_bank' list.

You may reuse elements of 'word_bank' as many times as needed.
"""
from typing import List


def count_construct(target: str, word_bank: List[str], memo: dict = None) -> int:
    """Count in how many ways the target can be constructed with a given word bank.

    Args:
        target (str): String that must be constructed.
        word_bank (List[str]): List of strings that can possibly construct the target.
        memo (dict): A hashmap to cache combinations already calculated.

    Returns:
        int: Returns the number of ways that 'target' can be constructed by
        concatenating elements of the 'word_bank' list.
    """
    if memo is None:
        memo = {}

    if target in memo:
        return memo[target]

    if not target:
        return 1

    count = 0

    for word in word_bank:
        if target.startswith(word):
            suffix_count = count_construct(target[len(word) :], word_bank, memo)
            count += suffix_count

    memo[target] = count
    return count


if __name__ == "__main__":
    print(count_construct("purple", ["purp", "p", "ur", "le", "purpl"]))  # 2
    print(count_construct("abcdef", ["ab", "abc", "cd", "def", "abcd"]))  # 1
    print(
        count_construct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"])
    )  # 0
    print(
        count_construct("enterapotentpot", ["a", "p", "ent", "enter", "ot", "o", "t"])
    )  # 4
    print(
        count_construct(
            "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
            ["e", "ee", "eee", "eeee", "eeeee", "eeeeee"],
        )
    )  # 0
