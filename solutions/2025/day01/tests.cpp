#include <gtest/gtest.h>

#include <iostream>
#include <filesystem>

#include "aoc_utils.hpp"
#include "solution.h"

namespace {
const std::string EXAMPLE_INPUT = R"(
)";
}

TEST(Day01, Part1_Example) {
  const std::string input = R"(L68
L30
R48
L5
R60
L55
L1
L99
R14
L82)";
  EXPECT_EQ(day01::solve_part_1(input), 3);
}

TEST(Day01, Part1_Solution) {
  const std::filesystem::path dirPath = AOC_INPUT_DIR;
  const std::filesystem::path fullPath = dirPath / "2025/day01.txt";
  auto input = aoc::utils::read_file(fullPath);
  EXPECT_EQ(day01::solve_part_1(input), 962);
}

TEST(Day01, Part2_Example) {
  const std::string input = R"(L68
L30
R48
L5
R60
L55
L1
L99
R14
L82)";
  EXPECT_EQ(day01::solve_part_2(input), 6);
}

TEST(Day01, Part2_Example2) {
  const std::string input = R"(L68
L30
R48
L5
R60
L55
L1
L399
R14
L82)";
  EXPECT_EQ(day01::solve_part_2(input), 9);
}

TEST(Day01, Part2_Solution) {
  const std::filesystem::path dirPath = AOC_INPUT_DIR;
  const std::filesystem::path fullPath = dirPath / "2025/day01.txt";
  auto input = aoc::utils::read_file(fullPath);
  EXPECT_EQ(day01::solve_part_2(input), 5782);
}
