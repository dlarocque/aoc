cmake_minimum_required(VERSION 3.27)
project(aoc_cpp VERSION 1.0 LANGUAGES CXX)

# --- Toolchain & Standard ---
set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

# --- Compiler Options & Sanitizers ---
# Enable AddressSanitizer (ASan) and UndefinedBehaviorSanitizer (UBSan) for Debug builds
option(ENABLE_SANITIZERS "Enable ASan and UBSan" ON)

if(ENABLE_SANITIZERS AND CMAKE_CXX_COMPILER_ID MATCHES "Clang|GNU")
    add_compile_options(-fsanitize=address,undefined -fno-omit-frame-pointer)
    add_link_options(-fsanitize=address,undefined)
endif()

# Strict warnings for production-quality code
add_compile_options(
    -Wall 
    -Wextra 
    -Wpedantic 
    -Wconversion 
    -Wshadow
    -Wno-unused-parameter
)

# --- Global Definitions ---
# Injects the absolute path to the inputs directory into the C++ code
add_compile_definitions(AOC_INPUT_DIR="${CMAKE_SOURCE_DIR}/inputs")

# --- Dependencies ---
include(FetchContent)

# GoogleTest
FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG v1.14.0
)
# Prevent GTest from installing itself to system directories
set(INSTALL_GTEST OFF CACHE BOOL "" FORCE)
FetchContent_MakeAvailable(googletest)

# --- Macros ---

# Macro: create_day
# Usage: create_day(day01)
# Assumes the existence of a folder named 'day01' in the current directory
# containing 'solution.cpp' and 'tests.cpp'.
macro(create_day DAY_NAME)
    # Get the current year from the parent directory name (or set variable in subdir)
    if(NOT DEFINED CURRENT_YEAR)
        get_filename_component(CURRENT_YEAR ${CMAKE_CURRENT_SOURCE_DIR} NAME)
    endif()

    set(TARGET_LIB "${CURRENT_YEAR}_${DAY_NAME}_lib")
    set(TARGET_TEST "test_${CURRENT_YEAR}_${DAY_NAME}")

    # 1. Solution Library (The logic)
    add_library(${TARGET_LIB} STATIC
        ${DAY_NAME}/solution.cpp
    )
    
    # Link shared utilities and include current day's headers
    target_link_libraries(${TARGET_LIB} PUBLIC aoc_utils)
    target_include_directories(${TARGET_LIB} PUBLIC ${DAY_NAME})

    # 2. Test Executable (The runner)
    add_executable(${TARGET_TEST}
        ${DAY_NAME}/tests.cpp
    )

    # Link logic and GTest
    target_link_libraries(${TARGET_TEST} PRIVATE
        ${TARGET_LIB}
        GTest::gtest_main
    )

    # Register with CTest
    include(GoogleTest)
    gtest_discover_tests(${TARGET_TEST})
endmacro()

# --- Subdirectories ---
add_subdirectory(libs/utils)
add_subdirectory(solutions/2024)