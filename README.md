# Advent of Code (C++)

My C++ solutions for Advent of Code.

## Quick Commands

### Build

```bash
cmake -B build -S .
cmake --build build
```

### Run a Specific Day

Each day is a standalone test executable.

```bash
# Run everything for Day 1
./build/solutions/2024/test_2024_day01

# Run ONLY the solution (skip example tests)
./build/solutions/2024/test_2024_day01 --gtest_filter=*Solution*
```

### Start a New Day

Generates files in solutions/ and updates CMake.

```bash
# ./scripts/create_day.sh <day> <year>
./scripts/create_day.sh 5 2024
```

## Dependencies

- C++23 Compiler
- CMake 3.27+
