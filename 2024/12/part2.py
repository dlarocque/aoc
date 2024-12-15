from __future__ import annotations

import argparse
import itertools
import os.path
import math

import pytest

from functools import cache

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')

def compute(s: str) -> int:
    s = s.strip()
    grid = [list(line) for line in s.splitlines()]
    grid = [['.'] + row + ['.'] for row in grid]
    grid = [['.'] * len(grid[0])] + grid + [['.'] * len(grid[0])]
    visited = [[False] * len(grid[0]) for _ in range(len(grid))]

    def print_grid():
        # print something else to indicate its visited
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if visited[i][j] == True:
                    print('.', end='')
                else:
                    print(grid[i][j], end='')
            print()
            print()

    def dfs(y: int, x: int, yfence: bool, xfence: bool) -> tuple[int, int]:
        if grid[y][x] == '.':
            return 0, 0
        visited[y][x] = True
        a, p, = 1, 0
        next_yfence, next_xfence = False, False
        for dy, dx in ((0, 1), (1, 0), (0, -1), (-1, 0)):
            if grid[y+dy][x+dx] != grid[y][x]:
                if dy != 0 and y+dy != 0 and y+dy != len(grid)-1:
                    if not yfence:
                        p += 1
                        next_yfence = True
                if dx != 0 and x+dx != 0 and x+dx != len(grid[0])-1:
                    if not xfence:
                        p += 1
                        next_xfence = True
        
        print_grid()
        for dy, dx in ((0, 1), (1, 0), (0, -1), (-1, 0)):
            if grid[y+dy][x+dx] == grid[y][x] and visited[y+dy][x+dx] == False:
                a2, p2 = dfs(y+dy, x+dx, next_yfence, next_xfence)
                a += a2
                p += p2

        print(f'y: {y}, x: {x}, a: {a}, p: {p}, yfence: {yfence}, xfence: {xfence}, next_yfence: {next_yfence}, next_xfence: {next_xfence}')
        return a, p

    res = 0
    for i in range(1, len(grid)-1):
        for j in range(1, len(grid[0])-1):
            if visited[i][j] == False:
                a, p = dfs(i, j, False, False)
                res += a * p

    return res

INPUT_1 = '''\
AAAA
BBCD
BBCC
EEEC
'''
EXPECTED_1 = 80
INPUT_2 = '''\
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
'''
EXPECTED_2 = 436
INPUT_3 = '''\
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
'''
EXPECTED_3 = 236
INPUT_4 = '''\
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
'''
EXPECTED_4 = 368
INPUT_5 = '''\
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
'''
EXPECTED_5 = 1206


@pytest.mark.parametrize(
    ('input_s', 'expected'),
    (
        (INPUT_1, EXPECTED_1),
        # (INPUT_2, EXPECTED_2),
        # (INPUT_3, EXPECTED_3),
        # (INPUT_4, EXPECTED_4),
        # (INPUT_5, EXPECTED_5),
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