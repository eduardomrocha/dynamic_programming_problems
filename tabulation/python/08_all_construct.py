"""All construct.

This script implements a solution to the all construct problem using tabulation.

Given a string 'target' and a list of strings 'word_bank', return a 2D list containing
all of the ways that 'target' can be constructed by concatenating elements of the
'word_bank' list. Each element of the 2D list should represent one combination that
constructs the 'target'.

You may reuse elements of 'word_bank' as many times as needed.
"""
from typing import List, Union


def all_construct(target: str, word_bank: List[str]) -> Union[List[List], List]:
    """Count in how many ways the target can be constructed with a given word bank.

    Args:
        target (str): String that must be constructed.
        word_bank (List[str]): List of strings that can possibly construct the target.

    Returns:
        List[List[str]]:
        int: Returns the number of ways that 'target' can be constructed by
        concatenating elements of the 'word_bank' list.
    """
    table: List = [[] for _ in range(len(target) + 1)]
    table[0].append([])

    for i in range(len(target)):
        if len(table[i]) > 0:
            suffix = target[i:]
            for word in word_bank:
                if suffix.startswith(word):
                    for combination in table[i]:
                        new_combination = combination + [word]
                        table[i + len(word)].append(new_combination)

    return table[len(target)]


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
        all_construct("aaaaaaaaaaz", ["a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa"],)
    )  # []
