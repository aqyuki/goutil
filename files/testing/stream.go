package testing

type (
	FileStreamChecker struct {
		StringContent string
	}
)

// Write writes to the stream. The data written to the stream is stored in StringContent.
func (s *FileStreamChecker) Write(p []byte) (n int, err error) {
	s.StringContent = string(p)
	return 0, nil
}

// EqualResult compares the data written to the stream with the data being sought.
func (s *FileStreamChecker) EqualResult(content string) bool {
	return s.StringContent == content
}
