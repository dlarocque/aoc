#ifdef NDEBUG
#include <iostream>

void aoc_assert_handler(const std::string& expression, const std::string& file, int line) {
  std::cout << "Assertion failed: " << expression << std::endl;
  std::cout << "" << file << ":" << line << std::endl;
}

#define AOC_ASSERT(condition)                             \
  do {                                                    \
    if (!(condition)) {                                   \
      aoc_assert_handler(#condition, __FILE__, __LINE__); \
    }                                                     \
  } while (0)
#else
#define AOC_ASSERT(condition) (void(0))
#endif
