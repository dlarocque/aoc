#include <gtest/gtest.h>

#include <iostream>

#include "aoc_utils.hpp"
#include "solution.h"

namespace {
const std::string EXAMPLE_INPUT = R"(987654321111111
811111111111119
234234234234278
818181911112111
)";
}

TEST(Day03, Part1_Example) { EXPECT_EQ(day03::solve_part_1(EXAMPLE_INPUT), 357); }

TEST(Day03, Part1_Solution) {
  const std::filesystem::path dirPath = AOC_INPUT_DIR;
  const std::filesystem::path fullPath = dirPath / "2025/day03.txt";
  auto input = aoc::utils::read_file(fullPath);
  std::cout << "[SOLUTION] Part 1: " << day03::solve_part_1(input) << std::endl;
}

/*
TEST(Day03, Part2_Example) {
    // EXPECT_EQ(day03::solve_part_2(EXAMPLE_INPUT), 0);
}

TEST(Day03, Part2_Solution) {
    // auto input = aoc::utils::read_file(AOC_INPUT_DIR "/day03.txt");
    // std::cout << "[SOLUTION] Part 2: " << day03::solve_part_2(input) << std::endl;
}
*/
