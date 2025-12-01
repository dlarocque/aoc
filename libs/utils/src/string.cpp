#include "string.hpp"

#include <cassert>
#include <iostream>
#include <sstream>

namespace aoc::utils {

vector<string> split(string s, char delim) {
  vector<string> tokens;
  stringstream ss(s);
  string line;
  while (getline(ss, line)) {
    size_t start{0};
    size_t end {0};
    while (end != string::npos) {
      end = line.find(delim, start);
      tokens.push_back(line.substr(start, end));
      start = end + 1;
    }
  }

  return tokens;
}

void trim(string& s) {
  const auto start = s.find_first_not_of(' ');
  const auto end = s.find_last_not_of(' ');
  s = s.substr(start, end - start);
}

}  // namespace aoc::utils
