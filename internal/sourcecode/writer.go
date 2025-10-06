package sourcecode

import (
	"fmt"
	"io"
)

type Writer struct{}

func (Writer) Write(w io.Writer, file File) error {
	const className = "some-generated-class-name"

	if _, err := fmt.Fprintf(w, "<style>\n.%s {\n%s\n}\n</style>", className, file.Template.CSS); err != nil {
		return fmt.Errorf("failed to write CSS: %w", err)
	}

	if _, err := fmt.Fprintf(w, "\n<div class=%q>\n%s\n</div>", className, file.Template.Markup); err != nil {
		return fmt.Errorf("failed to write markup: %w", err)
	}

	return nil
}
