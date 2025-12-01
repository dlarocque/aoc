#include "file.hpp"
#include <string>
#include <fstream>
#include <sstream>
#include <filesystem>

namespace aoc::utils {
  std::string read_file(const std::filesystem::path& path) {
    std::ifstream inf {path};
    if (!inf.is_open()) {
      throw std::runtime_error("input file could not be read: " + path.string());
    }

    std::stringstream buffer;
    buffer << inf.rdbuf();

    return buffer.str();
  }
}
