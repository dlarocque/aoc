#!/bin/bash

# Usage: ./scripts/create_day.sh <day_number> [year]
# Example: ./scripts/create_day.sh 05 2024

# 1. Parse Arguments
DAY_NUM=$1
YEAR=${2:-2024} # Default to 2024 if not provided

if [ -z "$DAY_NUM" ]; then
    echo "Error: Please provide a day number."
    echo "Usage: $0 <day_number> [year]"
    exit 1
fi

# Ensure day is padded (e.g., 5 -> 05)
PADDED_DAY=$(printf "%02d" $DAY_NUM)
DAY_NAME="day${PADDED_DAY}"

# Define Paths
SOLUTIONS_DIR="solutions/${YEAR}/${DAY_NAME}"
INPUTS_DIR="inputs/${YEAR}"
CMAKE_FILE="solutions/${YEAR}/CMakeLists.txt"

echo "Creating setup for Year: $YEAR, Day: $PADDED_DAY..."

# 2. Create Directories
mkdir -p "$SOLUTIONS_DIR"
mkdir -p "$INPUTS_DIR"

# 3. Create Input File
INPUT_FILE="${INPUTS_DIR}/${DAY_NAME}.txt"
if [ ! -f "$INPUT_FILE" ]; then
    touch "$INPUT_FILE"
    echo "Created input file: $INPUT_FILE"
else
    echo "Input file already exists: $INPUT_FILE"
fi

# 4. Generate C++ Files

# --- solution.h ---
HEADER_FILE="${SOLUTIONS_DIR}/solution.h"
if [ ! -f "$HEADER_FILE" ]; then
    cat <<EOF > "$HEADER_FILE"
#pragma once
#include <string>

namespace ${DAY_NAME} {
    long long solve_part_1(const std::string& input);
    long long solve_part_2(const std::string& input);
}
EOF
    echo "Created: $HEADER_FILE"
fi

# --- solution.cpp ---
CPP_FILE="${SOLUTIONS_DIR}/solution.cpp"
if [ ! -f "$CPP_FILE" ]; then
    cat <<EOF > "$CPP_FILE"
#include "solution.h"
#include "aoc_utils.hpp"
#include <string>
#include <vector>

namespace ${DAY_NAME} {

    long long solve_part_1(const std::string& input) {
        return 0;
    }

    long long solve_part_2(const std::string& input) {
        return 0;
    }
}
EOF
    echo "Created: $CPP_FILE"
fi

# --- tests.cpp ---
TEST_FILE="${SOLUTIONS_DIR}/tests.cpp"
if [ ! -f "$TEST_FILE" ]; then
    cat <<EOF > "$TEST_FILE"
#include <gtest/gtest.h>
#include <iostream>
#include "solution.h"
#include "aoc_utils.hpp"

namespace {
    const std::string EXAMPLE_INPUT = R"(
)";
}

TEST(Day${PADDED_DAY}, Part1_Example) {
    // EXPECT_EQ(${DAY_NAME}::solve_part_1(EXAMPLE_INPUT), 0);
}

TEST(Day${PADDED_DAY}, Part1_Solution) {
    // auto input = aoc::utils::read_file(AOC_INPUT_DIR "/${DAY_NAME}.txt");
    // std::cout << "[SOLUTION] Part 1: " << ${DAY_NAME}::solve_part_1(input) << std::endl;
}

/*
TEST(Day${PADDED_DAY}, Part2_Example) {
    // EXPECT_EQ(${DAY_NAME}::solve_part_2(EXAMPLE_INPUT), 0);
}

TEST(Day${PADDED_DAY}, Part2_Solution) {
    // auto input = aoc::utils::read_file(AOC_INPUT_DIR "/${DAY_NAME}.txt");
    // std::cout << "[SOLUTION] Part 2: " << ${DAY_NAME}::solve_part_2(input) << std::endl;
}
*/
EOF
    echo "Created: $TEST_FILE"
fi

# 5. Update CMakeLists.txt
if [ -f "$CMAKE_FILE" ]; then
    if grep -q "create_day(${DAY_NAME})" "$CMAKE_FILE"; then
        echo "CMakeLists.txt already contains ${DAY_NAME}."
    else
        echo "create_day(${DAY_NAME})" >> "$CMAKE_FILE"
        echo "Updated $CMAKE_FILE with create_day(${DAY_NAME})"
    fi
else
    echo "Warning: $CMAKE_FILE does not exist. Please check your year directory."
fi

echo "Done! Ready to solve Day $PADDED_DAY."