#include <gtest/gtest.h>

#include "string.hpp"

using namespace std;
using namespace aoc::utils;

TEST(SplitTest, SplitsWithSpaceDelim) {
  const vector<string> vec{"one", "two"};
  EXPECT_EQ(split("one two", ' '), vec);
}

TEST(SplitTest, SplitsWithCommaDelim) {
  const vector<string> vec{"one", "two", "three", "four"};
  EXPECT_EQ(split("one,two,three,four", ','), vec);
}

TEST(SplitTest, Splits) {
  vector<string> tokens{};
  {
    string s = "hello world";
    tokens = split(s, ' ');
  }

  vector<string> expected_tokens{"hello", "world"};
  EXPECT_EQ(tokens, expected_tokens);
}
