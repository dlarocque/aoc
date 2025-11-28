from __future__ import annotations

import argparse
import itertools
import os.path

import pytest

INPUT_TXT = os.path.join(os.path.dirname(__file__), 'input.txt')


def compute(s: str) -> int:
    s = s.strip()
    print(s)
    lines = s.split('\n')
    grid = [list(line) for line in lines]
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] == '0':
                grid[i][j] = 0
            else:
                grid[i][j] = int(grid[i][j])

    grid = [[-1] * len(grid[0])] + grid + [[-1] * len(grid[0])]
    for i in range(len(grid)):
        grid[i] = [-1] + grid[i] + [-1]
    
    cnt = 0
    def hike(i: int, j: int) -> int:
        for di, dj in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            if grid[i + di][j + dj] == grid[i][j] + 1:
                if grid[i + di][j + dj] == 9:
                    nonlocal cnt
                    cnt += 1
                else:
                    hike(i + di, j + dj)
    
    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[i]) - 1):
            if grid[i][j] == 0:
                print('hike', i, j)
                hike(i, j)

    return cnt


INPUT_S = '''\
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
'''
EXPECTED = 81


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