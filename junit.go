package main

type TestSuite struct {
	TestCases    []TestCase `xml:"testcase"`
	Name         string     `xml:"name,attr"`
	TestCount    int        `xml:"tests,attr"`
	ErrorCount   int        `xml:"errors,attr"`
	FailureCount int        `xml:"failures,attr"`
	Seconds      float64    `xml:"time,attr"`
	StdOut       *StdOut    `xml:"system-out"`
}

type TestCase struct {
	ClassName string         `xml:"classname,attr"`
	Name      string         `xml:"name,attr"`
	Time      float64        `xml:"time,attr"`
	Failure   *FailureResult `xml:"failure"`
	Error     *ErrorResult   `xml:"error"`
	Skipped   *SkippedResult `xml:"skipped"`
}

type FailureResult struct {
	Type        string `xml:"type,attr"`
	Message     string `xml:"message,attr"`
	Description string `xml:",chardata"`
}

type ErrorResult struct {
	Type        string `xml:"type,attr"`
	Message     string `xml:"message,attr"`
	Description string `xml:",chardata"`
}

var skipped = SkippedResult{}

type SkippedResult struct{}

func Skipped() *SkippedResult {
	return &skipped
}

type StdOut struct {
	Contents string `xml:",cdata"`
}
