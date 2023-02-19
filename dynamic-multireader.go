package dynamicmultireader

import "io"

type dynamicMultiReader struct {
	currentReader    io.Reader
	newReaderFactory NewReaderFactory
	n                int
}

func (mr *dynamicMultiReader) Read(p []byte) (n int, err error) {
	for mr.currentReader != nil {
		n, err = mr.currentReader.Read(p)
		if err == io.EOF {
			mr.currentReader = mr.newReaderFactory(mr.n)
			mr.n++
		}
		if n > 0 || err != io.EOF {
			if err == io.EOF && mr.currentReader != nil {
				// Don't return io.EOF yet. More readers remain.
				err = nil
			}
			return
		}
	}
	return 0, io.EOF
}

type NewReaderFactory func(
	// 0-based invocation counter
	n int,
) io.Reader

// Works as standard io.MultiReader, but instead of accapting an array of readers
// DynamicMultiReader accepts NewReaderFactory, which should return new readers, or nil, if no readers are remaining.
// This kind of DynamicMultiReader is needed when you get your source readers not at once, but sequentially.
func DynamicMultiReader(newReaderFactory NewReaderFactory) io.Reader {
	initialReader := newReaderFactory(0)

	return &dynamicMultiReader{initialReader, newReaderFactory, 1}
}
