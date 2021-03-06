"""All construct.

This script implements a solution to the all construct problem using memoization.

Given a string 'target' and a list of strings 'word_bank', return a 2D list containing
all of the ways that 'target' can be constructed by concatenating elements of the
'word_bank' list. Each element of the 2D list should represent one combination that
constructs the 'target'.

You may reuse elements of 'word_bank' as many times as needed.
"""
from typing import List, Union


def all_construct(
    target: str, word_bank: List[str], memo: dict = None
) -> Union[List[List], List]:
    """Construct all ways that target can be constructed using the words in word_bank.

    Args:
        target (str): String that must be constructed.
        word_bank (List[str]): List of strings that can possibly construct the target.
        memo (dict): A hashmap to cache combinations already calculated.

    Returns:
        List[List[str]]: A list containing lists of strings that together construct the
        target.
    """
    if memo is None:
        memo = {}

    if target in memo:
        return memo[target]

    if not target:
        return [[]]

    result = []

    for word in word_bank:
        if target.startswith(word):
            suffix_ways = all_construct(target[len(word) :], word_bank, memo)
            target_ways = [([word] + way) for way in suffix_ways]
            result += target_ways

    memo[target] = result
    return result


if __name__ == "__main__":
    print(all_construct("purple", ["purp", "p", "ur", "le", "purpl"]))
    """
    [
        ["purp", "le"],
        ["p", "ur", "p", "le"]
    ]
    """
    print(all_construct("abcdef", ["ab", "abc", "cd", "def", "abcd", "ef", "c"]))
    """
    [
        ["ab", "cd", "ef"],
        ["ab", "c", "def"],
        ["abc", "def"],
        ["abcd", "ef"]
    ]
    """
    print(
        all_construct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"])
    )  # []
    print(
        all_construct(
            "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
            ["e", "ee", "eee", "eeee", "eeeee", "eeeeee"],
        )
    )  # []
