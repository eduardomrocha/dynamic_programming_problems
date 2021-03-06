"""Can construct.

This script implements a solution to the can construct problem using memoization.

Given a string 'target' and a list of strings 'word_bank', return a boolean indicating
whether or not the 'target' can be constructed by concatenating elements of 'word_bank'
list.

You may reuse elements of 'word_bank' as many times as needed.
"""
from typing import List


def can_construct(target: str, word_bank: List[str], memo: dict = None) -> bool:
    """Check if a given target can be constructed with a given word bank.

    Args:
        target (str): String that must be constructed.
        word_bank (List[str]): List of strings that can possibly construct the target.
        memo (dict): A hashmap to cache combinations already calculated.

    Returns:
        bool: Returns True if it can be constructed or False if it cannot.
    """
    if memo is None:
        memo = {}

    if target in memo:
        return memo[target]

    if not target:
        return True

    for word in word_bank:
        if target.startswith(word):
            if can_construct(target[len(word) :], word_bank, memo):
                memo[target] = True
                return True

    memo[target] = False
    return False


if __name__ == "__main__":
    print(can_construct("abcdef", ["ab", "abc", "cd", "def", "abcd"]))  # True
    print(
        can_construct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"])
    )  # False
    print(
        can_construct("enterapotentpot", ["a", "p", "ent", "enter", "ot", "o", "t"])
    )  # True
    print(
        can_construct(
            "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
            ["e", "ee", "eee", "eeee", "eeeee", "eeeeee"],
        )
    )  # False
