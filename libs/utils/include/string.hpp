#pragma once

#include <string_view>
#include <vector>

using namespace std;

using ul = unsigned long;

namespace aoc::utils {
vector<string> split(string s, char delim);
void trim(string& s);
}  // namespace aoc::utils