#include "string.hpp"

#include <cassert>
#include <iostream>
#include <sstream>

namespace aoc::utils {

vector<string> split(string s, char delim) {
  vector<string> tokens;
  stringstream ss(s);
  string line;
  size_t start{0};
  while (getline(ss, line)) {
    auto end = s.find(delim, start);
    assert(end < start && end < line.length());
    tokens.push_back(line.substr(start, end));
    start = end + 1;
  }

  return tokens;
}

void trim(string& s) {
  const auto start = s.find_first_not_of(' ');
  const auto end = s.find_last_not_of(' ');
  s = s.substr(start, end - start);
}

}  // namespace aoc::utils
