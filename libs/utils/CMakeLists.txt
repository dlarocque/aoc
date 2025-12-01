# --- Shared Utility Library ---
add_library(aoc_utils STATIC
    src/string.cpp
)

# Expose the include directory to any target linking this library
target_include_directories(aoc_utils PUBLIC 
    ${CMAKE_CURRENT_SOURCE_DIR}/include
)

# Link GTest just in case utils need internal assertions
target_link_libraries(aoc_utils PUBLIC GTest::gtest)

# --- Utility Unit Tests ---
add_executable(utils_tests
    tests/test_string.cpp
    # Add other utility test files here
)

target_link_libraries(utils_tests PRIVATE
    aoc_utils
    GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(utils_tests)