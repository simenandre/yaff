package main_test

import (
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/simenandre/yaff/internal/sourcecode"
)

//go:embed README.md
var readme string

func getExampleSource(t *testing.T) string {
	const (
		prefix = "```html"
		suffix = "```"
	)

	withoutPrefix := strings.Split(readme, prefix)[1]
	source := strings.Split(withoutPrefix, suffix)[0]
	require.NotEmpty(t, source, "example source code extracted from readme")
	return source
}

func TestReadmeExample(t *testing.T) {
	source := getExampleSource(t)
	const sourceFilePath = "/some/path/to/the/file.jaff"

	var parser sourcecode.Parser
	file, err := parser.ParseFile(&fakeFile{
		Reader: strings.NewReader(source),
		name:   sourceFilePath,
	})
	require.NoError(t, err, "failed to parse file")

	t.Run("parse", func(t *testing.T) {
		expected := sourcecode.File{
			Path: sourceFilePath,
			Preamble: strings.TrimSpace(`
package main

import (
	"github.com/simenandre/yaff/components"
)

colors := [2]string{"black", "white"}
		`),
			Template: sourcecode.Template{
				Markup: strings.TrimSpace(`
<h1>Hello world</h1>


{#each colors as color}
	<components.HelloWorld color={color} />
{/each}
			`),
				CSS: strings.TrimSpace(`
h1 {
	padding: 1px:
}
			`),
			},
		}
		require.Empty(t, cmp.Diff(expected, file))
	})
	if t.Failed() {
		t.FailNow()
	}

	var (
		writer sourcecode.Writer
		output strings.Builder
	)
	require.NoError(t, writer.Write(&output, file))

	t.Run("write", func(t *testing.T) {
		expected := strings.TrimSpace(`
<style>
.some-generated-class-name {
h1 {
	padding: 1px:
}
}
</style>
<div class="some-generated-class-name">
<h1>Hello world</h1>


{#each colors as color}
	<components.HelloWorld color={color} />
{/each}
</div>
		`)
		require.Empty(t, cmp.Diff(
			strings.Split(expected, "\n"),
			strings.Split(output.String(), "\n")),
		)
	})
}

type fakeFile struct {
	io.Reader
	name string
}

func (f *fakeFile) Name() string {
	return f.name
}
