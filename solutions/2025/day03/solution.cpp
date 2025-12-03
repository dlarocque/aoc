#include "solution.h"

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

#include "print.hpp"
#include "string.hpp"

using namespace std;
using namespace aoc::utils;

using ll = long long;

namespace day03 {

ll joined_digits(ll a, ll b) { return (a * 10) + b; }

long long solve_part_1(const std::string& input) {
  ll res = 0;
  const auto lines = aoc::utils::split(input, '\n');
  for (const auto& line : lines) {
    ll max = 0;
    const vector<ll> digits = tokenizeToDigits(line);
    vector<ll> maxToLeftIncl(digits.size(), 0);
    vector<ll> maxToRight(digits.size(), 0);

    maxToLeftIncl[0] = digits[0];
    for (size_t i = 1; i < digits.size(); i++) {
      maxToLeftIncl[i] = std::max(maxToLeftIncl[i - 1], digits[i]);
    }
    maxToRight[maxToRight.size() - 1] = -1000;
    for (int i = static_cast<int>(maxToRight.size()) - 2; i >= 0; i--) {
      maxToRight[i] = std::max(digits[i + 1], maxToRight[i + 1]);
    }

    cout << "   " << to_string(digits) << endl;
    cout << "l: " << to_string(maxToLeftIncl) << endl;
    cout << "r: " << to_string(maxToRight) << endl;

    for (size_t i = 0; i < digits.size(); i++) {
      max = std::max(joined_digits(maxToLeftIncl[i], maxToRight[i]), max);
    }

    res += max;
  }
  return res;
}

long long solve_part_2(const std::string& input) { return 0; }
}  // namespace day03
