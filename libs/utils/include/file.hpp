#pragma once

#include <string_view>
#include <vector>
#include <filesystem>

using ul = unsigned long;

namespace aoc::utils {
  std::string read_file(const std::filesystem::path& path);
}  // namespace aoc::utils
