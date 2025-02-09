package mock

import "errors"

type MockImageProcessor struct {
	CalledInputPath  string
	CalledOutputPath string
	ShouldFail       bool
}

func (m *MockImageProcessor) ConvertToGray(inputPath, outputPath string) error {
	m.CalledInputPath = inputPath
	m.CalledOutputPath = outputPath
	if m.ShouldFail {
		return errors.New("mock conversion failed")
	}
	return nil
}
