#include <gtest/gtest.h>

#include <iostream>

#include "aoc_utils.hpp"
#include "solution.h"

namespace {
const std::string EXAMPLE_INPUT = R"(
)";
}

TEST(Day02, Part1_Example) {
  const std::string input =
      R"(11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124)";
  EXPECT_EQ(day02::solve_part_1(input), 1227775554);
}

TEST(Day02, Part1_Solution) {
  const std::filesystem::path dirPath = AOC_INPUT_DIR;
  const std::filesystem::path fullPath = dirPath / "2025/day02.txt";
  auto input = aoc::utils::read_file(fullPath);
  EXPECT_EQ(day02::solve_part_1(input), 0);
}

TEST(Day02, Part2_Example) {
  const std::string input =
      R"(11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124)";
  EXPECT_EQ(day02::solve_part_2(input), 4174379265);
}

TEST(Day02, Part2_Solution) {
  const std::filesystem::path dirPath = AOC_INPUT_DIR;
  const std::filesystem::path fullPath = dirPath / "2025/day02.txt";
  auto input = aoc::utils::read_file(fullPath);
  EXPECT_EQ(day02::solve_part_2(input), 0);
}
