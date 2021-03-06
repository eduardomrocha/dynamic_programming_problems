"""Count construct.

This script implements a solution to the count construct problem using tabulation.

Given a string 'target' and a list of strings 'word_bank', return the number of ways
that 'target' can be constructed by concatenating elements of the 'word_bank' list.

You may reuse elements of 'word_bank' as many times as needed.
"""
from typing import List


def count_construct(target: str, word_bank: List[str]) -> int:
    """Count in how many ways the target can be constructed with a given word bank.

    Args:
        target (str): String that must be constructed.
        word_bank (List[str]): List of strings that can possibly construct the target.

    Returns:
        int: Returns the number of ways that 'target' can be constructed by
        concatenating elements of the 'word_bank' list.
    """
    table: List = [0] * (len(target) + 1)
    table[0] = 1

    for i in range(len(target)):
        if table[i] > 0:
            suffix = target[i:]
            for word in word_bank:
                if suffix.startswith(word):
                    table[i + len(word)] += table[i]

    return table[len(target)]


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
