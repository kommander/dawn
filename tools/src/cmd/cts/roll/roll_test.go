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

package roll

import (
	"context"
	"fmt"
	"testing"

	"dawn.googlesource.com/dawn/tools/src/buildbucket"
	"dawn.googlesource.com/dawn/tools/src/cmd/cts/common"
	"dawn.googlesource.com/dawn/tools/src/git"
	"dawn.googlesource.com/dawn/tools/src/oswrapper"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

// The following functions and those that call it cannot currently be tested due
// to using exec.
//   - generateFiles
//   - gnTSDepList
//   - genTestList

func MustParseHash(s string) git.Hash {
	hash, err := git.ParseHash("d5e605a556408eaeeda64fb9d33c3f596fd90b70")
	if err != nil {
		panic(err)
	}
	return hash
}

func rollCommitMessageFor(ctsGitURL string) string {
	r := roller{
		cfg: common.Config{
			Builders: map[string]buildbucket.Builder{
				"Win":   {Project: "chromium", Bucket: "try", Builder: "win-dawn-rel"},
				"Mac":   {Project: "dawn", Bucket: "try", Builder: "mac-dbg"},
				"Linux": {Project: "chromium", Bucket: "try", Builder: "linux-dawn-rel"},
			},
		},
		flags: rollerFlags{ctsGitURL: ctsGitURL},
	}

	r.cfg.Git.CTS.Host = "chromium.googlesource.com"
	r.cfg.Git.CTS.Project = "external/github.com/gpuweb/cts"

	msg := r.rollCommitMessage(
		"d5e605a556408eaeeda64fb9d33c3f596fd90b70",
		"29275672eefe76986bd4baa7c29ed17b66616b1b",
		[]git.CommitInfo{
			{
				Hash:    MustParseHash("d5e605a556408eaeeda64fb9d33c3f596fd90b70"),
				Subject: "Added thing A",
			},
			{
				Hash:    MustParseHash("29275672eefe76986bd4baa7c29ed17b66616b1b"),
				Subject: "Tweaked thing B",
			},
		},
		"I4aa059c6c183e622975b74dbdfdfe0b12341ae15",
	)

	return msg
}

func TestRollCommitMessageFromInternal(t *testing.T) {
	msg := rollCommitMessageFor("https://chromium.googlesource.com/external/github.com/gpuweb/cts")
	expect := `Roll third_party/webgpu-cts/ d5e605a55..29275672e (2 commits)

Regenerated:
 - expectations.txt
 - compat-expectations.txt
 - ts_sources.txt
 - test_list.txt
 - resource_files.txt
 - webtest .html files


https://chromium.googlesource.com/external/github.com/gpuweb/cts/+log/d5e605a55640..29275672eefe
 - d5e605 Added thing A
 - d5e605 Tweaked thing B

Created with './tools/run cts roll'

Cq-Include-Trybots: luci.chromium.try:linux-dawn-rel,win-dawn-rel;luci.dawn.try:mac-dbg
Include-Ci-Only-Tests: true
Change-Id: I4aa059c6c183e622975b74dbdfdfe0b12341ae15
`
	if diff := cmp.Diff(msg, expect); diff != "" {
		t.Errorf("rollCommitMessage: %v", diff)
	}
}

func TestRollCommitMessageFromExternal(t *testing.T) {
	msg := rollCommitMessageFor("https://www.github.com/a_cts_contributor/cts.git")

	expect := `[DO NOT` + ` SUBMIT] Roll third_party/webgpu-cts/ d5e605a55..29275672e (2 commits)

Rolled from external repo: https://www.github.com/a_cts_contributor/cts.git

Regenerated:
 - expectations.txt
 - compat-expectations.txt
 - ts_sources.txt
 - test_list.txt
 - resource_files.txt
 - webtest .html files


https://chromium.googlesource.com/external/github.com/gpuweb/cts/+log/d5e605a55640..29275672eefe
 - d5e605 Added thing A
 - d5e605 Tweaked thing B

Created with './tools/run cts roll'

Cq-Include-Trybots: luci.chromium.try:linux-dawn-rel,win-dawn-rel;luci.dawn.try:mac-dbg
Include-Ci-Only-Tests: true
Commit: false
Change-Id: I4aa059c6c183e622975b74dbdfdfe0b12341ae15
`
	if diff := cmp.Diff(msg, expect); diff != "" {
		t.Errorf("rollCommitMessage: %v", diff)
	}
}

func TestRoller_GenResourceFilesList_NonExistentDirectory(t *testing.T) {
	wrapper := oswrapper.CreateMemMapOSWrapper()
	r := roller{
		ctsDir: "/non_existent",
	}

	_, err := r.genResourceFilesList(context.Background(), wrapper)
	require.ErrorContains(t, err, "open /non_existent/src/resources: file does not exist")
}

func TestRoller_GenResourceFilesList_Success(t *testing.T) {
	wrapper := oswrapper.CreateMemMapOSWrapper()
	err := wrapper.MkdirAll("/root/src/resources/subdir_1", 0o700)
	require.NoErrorf(t, err, "Error creating directory: %v", err)
	err = wrapper.MkdirAll("/root/src/resources/subdir_2", 0o700)
	require.NoErrorf(t, err, "Error creating cirectory: %v", err)

	f, err := wrapper.Create("/root/src/resources/subdir_1/file_1.txt")
	require.NoErrorf(t, err, "Error creating file: %v", err)
	defer f.Close()
	f, err = wrapper.Create("/root/src/resources/subdir_1/file_2.txt")
	require.NoErrorf(t, err, "Error creating file: %v", err)
	defer f.Close()
	f, err = wrapper.Create("/root/src/resources/subdir_2/file_1.txt")
	require.NoErrorf(t, err, "Error creating file: %v", err)
	defer f.Close()

	r := roller{
		ctsDir: "/root",
	}
	fileList, err := r.genResourceFilesList(context.Background(), wrapper)
	require.NoErrorf(t, err, "Error generating resource file list: %v", err)
	expectedList := `subdir_1/file_1.txt
subdir_1/file_2.txt
subdir_2/file_1.txt
`
	require.Equal(t, expectedList, fileList)
}

func TestRoller_GenWebTestSources_NonExistentDirectory(t *testing.T) {
	wrapper := oswrapper.CreateMemMapOSWrapper()
	r := roller{
		ctsDir: "/non_existent",
	}

	_, err := r.genWebTestSources(context.Background(), wrapper)
	require.ErrorContains(t, err, "open /non_existent/src/webgpu: file does not exist")
}

func TestRoller_GenWebTestSources_Success(t *testing.T) {
	htmlContent := `<!-- Comment -->
<html name=foo>
</html>`

	wrapper := oswrapper.CreateMemMapOSWrapper()
	directories := []string{"subdir_1", "subdir_2"}
	files := []string{"file_1.html", "file_2.txt", "file_3.html"}
	for _, directory := range directories {
		directory_path := fmt.Sprintf("/root/src/webgpu/%s", directory)
		err := wrapper.MkdirAll(directory_path, 0o700)
		require.NoErrorf(t, err, "Error creating directory: %v", err)
		for _, filename := range files {
			f, err := wrapper.Create(fmt.Sprintf("%s/%s", directory_path, filename))
			require.NoErrorf(t, err, "Error creating file: %v", err)
			defer f.Close()
			f.Write([]byte(htmlContent))
		}
	}

	r := roller{
		ctsDir: "/root",
	}
	fileMap, err := r.genWebTestSources(context.Background(), wrapper)
	require.NoErrorf(t, err, "Error generating web test sources: %v", err)

	expectedHtmlContent := `<!-- Comment -->
<html name=foo>
  <base ref="/gen/third_party/dawn/webgpu-cts/src/webgpu" />
</html>`
	expectedMapContent := map[string]string{
		"webgpu-cts/webtests/subdir_1/file_1.html": expectedHtmlContent,
		"webgpu-cts/webtests/subdir_1/file_3.html": expectedHtmlContent,
		"webgpu-cts/webtests/subdir_2/file_1.html": expectedHtmlContent,
		"webgpu-cts/webtests/subdir_2/file_3.html": expectedHtmlContent,
	}
	require.Equal(t, expectedMapContent, fileMap)
}
