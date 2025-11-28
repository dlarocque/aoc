#include "string.hpp"

namespace aoc::utils {
constexpr vector<string_view> split(string_view s, string_view delim) {
  vector<string_view> res{};
  ul start = 0;
  ul n = 0;
  while (start < s.length()) {
    n = s.find(delim, start);
    res.push_back(s.substr(start, n));
    start = n + delim.length();
  }

  return res;
}
}  // namespace aoc::utils