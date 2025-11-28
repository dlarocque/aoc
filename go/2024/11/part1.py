from __future__ import annotations

import argparse
import itertools
import os.path
import math

import pytest

from functools import cache

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')

@cache
def cnt(x: int, d=75) -> int:
    if d == 0: return 1
    if x == 0: return cnt(1, d-1)
    l = math.floor(math.log10(x)) + 1
    if l % 2 == 0:
        return cnt((x - (x % 10**(l//2))) / (10**(l//2)), d-1) + cnt(x % 10**(l//2), d-1)

    return cnt(x * 2024, d-1)

def compute(s: str) -> int:
    s = s.strip()
    arr = [int(x) for x in s.split(' ')]
    return sum(cnt(x) for x in arr)


INPUT_S = '''\
125 17
'''
EXPECTED = 55312


@pytest.mark.parametrize(
    ('input_s', 'expected'),
    (
        (INPUT_S, EXPECTED),
    ),
)
def test(input_s: str, expected: int) -> None:
    assert compute(input_s) == expected


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument('data_file', nargs='?', default=INPUT_TXT)
    args = parser.parse_args()

    with open(args.data_file) as f:
        print(compute(f.read()))

    return 0


if __name__ == '__main__':
    raise SystemExit(main())