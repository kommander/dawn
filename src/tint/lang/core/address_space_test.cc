// Copyright 2022 The Dawn & Tint Authors
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

////////////////////////////////////////////////////////////////////////////////
// File generated by 'tools/src/cmd/gen' using the template:
//   src/tint/lang/core/address_space_test.cc.tmpl
//
// To regenerate run: './tools/run gen'
//
//                       Do not modify this file directly
////////////////////////////////////////////////////////////////////////////////

#include "src/tint/lang/core/address_space.h"

#include <gtest/gtest.h>

#include <string>

#include "src/tint/utils/text/string.h"

namespace tint::core {
namespace {

namespace parse_print_tests {

struct Case {
    const char* string;
    AddressSpace value;
};

inline std::ostream& operator<<(std::ostream& out, Case c) {
    return out << "'" << std::string(c.string) << "'";
}

static constexpr Case kValidCases[] = {
    {"function", AddressSpace::kFunction},      {"immediate", AddressSpace::kImmediate},
    {"pixel_local", AddressSpace::kPixelLocal}, {"private", AddressSpace::kPrivate},
    {"storage", AddressSpace::kStorage},        {"uniform", AddressSpace::kUniform},
    {"workgroup", AddressSpace::kWorkgroup},
};

static constexpr Case kInvalidCases[] = {
    {"fccnctin", AddressSpace::kUndefined},      {"ucti3", AddressSpace::kUndefined},
    {"functVon", AddressSpace::kUndefined},      {"imm1diate", AddressSpace::kUndefined},
    {"qqmmeJite", AddressSpace::kUndefined},     {"illmediat77", AddressSpace::kUndefined},
    {"pixppHq_local", AddressSpace::kUndefined}, {"ciellovl", AddressSpace::kUndefined},
    {"ibGl_local", AddressSpace::kUndefined},    {"priviive", AddressSpace::kUndefined},
    {"8WWivate", AddressSpace::kUndefined},      {"pxxvate", AddressSpace::kUndefined},
    {"sXraggg", AddressSpace::kUndefined},       {"traXe", AddressSpace::kUndefined},
    {"stor3ge", AddressSpace::kUndefined},       {"Eniform", AddressSpace::kUndefined},
    {"uPTTform", AddressSpace::kUndefined},      {"unifodxx", AddressSpace::kUndefined},
    {"44orkgroup", AddressSpace::kUndefined},    {"VVorkSSroup", AddressSpace::kUndefined},
    {"wrkgro2Rp", AddressSpace::kUndefined},
};

using AddressSpaceParseTest = testing::TestWithParam<Case>;

TEST_P(AddressSpaceParseTest, Parse) {
    const char* string = GetParam().string;
    AddressSpace expect = GetParam().value;
    EXPECT_EQ(expect, ParseAddressSpace(string));
}

INSTANTIATE_TEST_SUITE_P(ValidCases, AddressSpaceParseTest, testing::ValuesIn(kValidCases));
INSTANTIATE_TEST_SUITE_P(InvalidCases, AddressSpaceParseTest, testing::ValuesIn(kInvalidCases));

using AddressSpacePrintTest = testing::TestWithParam<Case>;

TEST_P(AddressSpacePrintTest, Print) {
    AddressSpace value = GetParam().value;
    const char* expect = GetParam().string;
    EXPECT_EQ(expect, ToString(value));
}

INSTANTIATE_TEST_SUITE_P(ValidCases, AddressSpacePrintTest, testing::ValuesIn(kValidCases));

}  // namespace parse_print_tests

}  // namespace
}  // namespace tint::core
