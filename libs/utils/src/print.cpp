#include "print.hpp"

#include <sstream>

using namespace std;

namespace aoc::utils {
string to_string(const vector<long long>& v) {
  std::stringstream ss;
  ss << "[";
  for (const auto& x : v) {
    ss << x << ", ";
  }
  ss << "]";
  return ss.str();
}
}  // namespace aoc::utils
