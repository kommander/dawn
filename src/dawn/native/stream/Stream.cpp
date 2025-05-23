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

#include "dawn/native/stream/Stream.h"

#include <string>

#include "dawn/native/Limits.h"

namespace dawn::native::stream {

constexpr void StreamIn(Sink* s) {}

MaybeError StreamOut(Source* s) {
    return {};
}

template <>
void Stream<std::string>::Write(Sink* s, const std::string& t) {
    StreamIn(s, t.length());
    size_t size = t.length() * sizeof(char);
    if (size > 0) {
        void* ptr = s->GetSpace(size);
        memcpy(ptr, t.data(), size);
    }
}

template <>
MaybeError Stream<std::string>::Read(Source* s, std::string* t) {
    size_t length;
    DAWN_TRY(StreamOut(s, &length));
    const void* ptr;
    DAWN_TRY(s->Read(&ptr, length));
    *t = std::string(static_cast<const char*>(ptr), length);
    return {};
}

template <>
void Stream<std::wstring>::Write(Sink* s, const std::wstring& t) {
    StreamIn(s, t.length());
    size_t size = t.length() * sizeof(wchar_t);
    if (size > 0) {
        void* ptr = s->GetSpace(size);
        memcpy(ptr, t.data(), size);
    }
}

template <>
MaybeError Stream<std::wstring>::Read(Source* s, std::wstring* t) {
    size_t length;
    DAWN_TRY(StreamOut(s, &length));
    size_t size = length * sizeof(wchar_t);
    const void* ptr;
    DAWN_TRY(s->Read(&ptr, size));
    *t = std::wstring(static_cast<const wchar_t*>(ptr), length);
    return {};
}

template <>
void Stream<std::string_view>::Write(Sink* s, const std::string_view& t) {
    StreamIn(s, t.length());
    size_t size = t.length() * sizeof(char);
    if (size > 0) {
        void* ptr = s->GetSpace(size);
        memcpy(ptr, t.data(), size);
    }
}

template <>
void Stream<std::wstring_view>::Write(Sink* s, const std::wstring_view& t) {
    StreamIn(s, t.length());
    size_t size = t.length() * sizeof(wchar_t);
    if (size > 0) {
        void* ptr = s->GetSpace(size);
        memcpy(ptr, t.data(), size);
    }
}

}  // namespace dawn::native::stream
