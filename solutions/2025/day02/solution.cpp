#include "solution.h"

#include <cassert>
#include <iostream>
#include <numeric>
#include <set>
#include <string>
#include <utility>
#include <vector>

#include "debug_log.hpp"
#include "string.hpp"

using namespace std;
using namespace aoc::utils;

using ll = long long;

namespace day02 {

ll expand_and_mirror(ll x) { return x + (x * powl(10, (to_string(x).length()))); }

ll solve_part_1(const string& input) {
  ll res = 0;
  set<ll> ids{};
  vector<string> range_strs = split(input, ',');
  for (string& range_str : range_strs) {
    cout << range_str << endl;
    vector<string> temp = split(range_str, '-');
    pair<ll, ll> range = {stoll(temp.at(0)), stoll(temp.at(1))};
    string s1 = to_string(range.first);
    string s2 = to_string(range.second);

    if (s1.length() % 2 == 1 && s2.length() % 2 == 1) {
      continue;
    }

    for (ll i = range.first; i <= range.second; i++) {
      string i_str = to_string(i);
      ll half = stoll(i_str.substr(0, (i_str.length() + 1) / 2));
      ll expanded = expand_and_mirror(half);
      // cout << i << " " << expanded << endl;
      if (expanded >= range.first && expanded <= range.second) {
        // cout << "found " << expanded << endl;
        ids.insert(expanded);
        // res += expanded;
      }
    }

    // brute force
    // start = first two digits of start
    // end =

    /*
        ll end = expand_and_mirror(stoll(s2.substr(0, (s2.length()) / 2)));
        ll curr = stoll(s1.substr(0, (s1.length() + 1) / 2));
        ll full = expand_and_mirror(curr);
        cout << "curr: " << curr << ", full: " << full << ", end: " << end << endl;
        while (full <= end) {
          cout << "[" << curr << " -> " << full << "]" << endl;
          if (full >= range.first && full <= range.second) {
            cout << "  INVALID ID: " << full << endl;
            res += full;
          }
          curr++;
          full = expand_and_mirror(curr);
        }
      }
    */
  }
  return accumulate(ids.begin(), ids.end(), 0LL);
}

long long solve_part_2(const std::string& input) { return 0; }
}  // namespace day02
