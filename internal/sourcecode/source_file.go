package sourcecode

type File struct {
	Path     string
	Preamble string
	Template Template
}

type Template struct {
	Markup string
	CSS    string
}
