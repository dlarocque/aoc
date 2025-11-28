#pragma once

#include <string_view>
#include <vector>

using namespace std;

using ul = unsigned long;

namespace aoc::utils {
constexpr vector<string_view> split(string_view s, string_view delim);
}