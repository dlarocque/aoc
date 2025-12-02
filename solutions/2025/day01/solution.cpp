#include "solution.h"

#include <cassert>
#include <cstring>
#include <iostream>
#include <string>
#include <vector>

#include "aoc_utils.hpp"
#include "string.hpp"

using namespace std;
using namespace aoc::utils;

namespace day01 {

long long solve_part_1(const std::string& input) {
  // R is addition
  // L is subtraction
  int pos = 50;
  int res = 0;

  const auto instructions = split(input, '\n');
  assert(instructions.size() > 0);

  for (const auto& instruction : instructions) {
    assert(instruction.length() > 1);
    int dist = stoi(instruction.substr(1, instruction.length() - 1).data()) % 100;
    assert(dist > 0 & dist < 100);
    if (instruction[0] == 'L') {  // subtract
      if (dist > pos) {
        pos = 100 - (dist - pos);
      } else {
        pos -= dist;
      }
    } else {  // add
      if ((pos + dist) > 99) {
        pos = 0 + ((dist + pos) - 100);
      } else {
        pos += dist;
      }
    }

    if (pos == 0) {
      res++;
    }
  }

  return res;
}

long long solve_part_2(const std::string& input) {
  // R is addition
  // L is subtraction
  int pos = 50;
  int res = 0;

  const auto instructions = split(input, '\n');

  for (const auto& instruction : instructions) {
    int dist = stoi(instruction.substr(1, instruction.length() - 1).data());
    res += dist / 100;
    dist %= 100;
    if (instruction[0] == 'L') { // subtract
      // We roll past 0 if the distance is greater than pos.
      if (dist > pos) {
        if (pos != 0) {
          res++;
        }
        pos = 100 - (dist - pos);
      } else {
        pos -= dist;
      }
    } else {  // add
      // We roll past 0 if the distance is greater than the distance between pos and 100. 
      if ((pos + dist) > 99) {
        pos = 0 + ((dist + pos) - 100);
        if (pos != 0) {
          res++;
        }
      } else {
        pos += dist;
      }
    }

    if (pos == 0) {
      res++;
    }
    cout << "Instruction: " << instruction << " -> (" << pos << ", " << res << ")" << endl;
  }

  return res;
}
} // namespace day01
