#include "string.hpp"

#include <cassert>
#include <sstream>
#include <string>

namespace aoc::utils {

vector<long long> tokenizeToDigits(string s) {
  vector<long long> tokens;
  stringstream ss(s);
  string line;
  while (getline(ss, line)) {
    for (size_t i = 0; i < line.length(); i++) {
      tokens.push_back(static_cast<long long>(line.at(i) - '0'));
    }
  }

  return tokens;
}

vector<string> split(string s, char delim) {
  vector<string> tokens;
  stringstream ss(s);
  string line;
  while (getline(ss, line)) {
    // cout << "line: " << line << endl;
    size_t start{0};
    size_t end{0};
    while (end != string::npos) {
      end = line.find(delim, start);
      // cout << "token: (" << start << ", " << end << "), " << line.substr(start, (end-start)) <<
      // endl;
      tokens.push_back(line.substr(start, end - start));
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
