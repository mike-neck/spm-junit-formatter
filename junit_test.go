package main

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestCaseSuccess(t *testing.T) {
	testCase := TestCase{ClassName: "MyApp2Tests.MyApp2Tests", Name: "testExample2", Time: 0.001}
	bytes, err := xml.Marshal(&testCase)
	if err != nil {
		t.Fail()
	}
	xmlString := string(bytes)
	assert.Equal(t, `<TestCase classname="MyApp2Tests.MyApp2Tests" name="testExample2" time="0.001"></TestCase>`, xmlString)
}

func TestTestCaseFailure(t *testing.T) {
	failure := &FailureResult{Type: "XCTAssertFalse failed", Message: "test2", Description: "/Users/mike/tmp/swift-package-test/MyApp/Tests/MyAppTests/MyAppTests.swift:13"}
	testCase := &TestCase{ClassName: "MyApp2Tests.MyApp2Tests", Name: "testExample2", Time: 0.001, Failure: failure}
	bytes, err := xml.Marshal(&testCase)
	if err != nil {
		t.Fail()
	}
	xmlString := string(bytes)
	assert.Equal(t, `<TestCase classname="MyApp2Tests.MyApp2Tests" name="testExample2" time="0.001"><failure type="XCTAssertFalse failed" message="test2">/Users/mike/tmp/swift-package-test/MyApp/Tests/MyAppTests/MyAppTests.swift:13</failure></TestCase>`, xmlString)
}

func TestTestCaseError(t *testing.T) {
	errResult := &ErrorResult{Type: "XCTAssertFalse failed", Message: "test2", Description: "<unknown>:0"}
	testCase := &TestCase{ClassName: "MyApp2Tests.MyApp2Tests", Name: "testExample2", Time: 0.001, Error: errResult}
	bytes, err := xml.Marshal(&testCase)
	if err != nil {
		t.Fail()
	}
	xmlString := string(bytes)
	assert.Equal(t, `<TestCase classname="MyApp2Tests.MyApp2Tests" name="testExample2" time="0.001"><error type="XCTAssertFalse failed" message="test2">&lt;unknown&gt;:0</error></TestCase>`, xmlString)
}

func TestTestCaseSkipped(t *testing.T) {
	testCase := &TestCase{ClassName: "MyApp2Tests.MyApp2Tests", Name: "testExample2", Time: 0.001, Skipped: Skipped()}
	bytes, err := xml.Marshal(&testCase)
	if err != nil {
		t.Fail()
	}
	xmlString := string(bytes)
	assert.Equal(t, `<TestCase classname="MyApp2Tests.MyApp2Tests" name="testExample2" time="0.001"><skipped></skipped></TestCase>`, xmlString)
}

func TestTestSuite(t *testing.T) {
	var testCases []TestCase
	testCases = append(testCases, TestCase{ClassName: "MyApp2Tests.MyApp2Tests", Name: "testExample2", Time: 0.001})
	failure := &FailureResult{Type: "XCTAssertFalse failed", Message: "test2", Description: "/Users/mike/tmp/swift-package-test/MyApp/Tests/MyAppTests/MyAppTests.swift:13"}
	testCases = append(testCases, TestCase{ClassName: "MyApp2Tests.MyApp2Tests", Name: "testExample2", Time: 0.001, Failure: failure})
	stdOut := StdOut{"test stdout test stdout"}
	suite := TestSuite{TestCases: testCases, Name: "MyAppPackageTests.xctest.MyApp2Tests", TestCount: 2, ErrorCount: 0, FailureCount: 1, Seconds: 0.214, StdOut: &stdOut}
	bytes, err := xml.Marshal(&suite)
	if err != nil {
		t.Fail()
	}
	xmlString := string(bytes)
	assert.Equal(
		t,
		`<TestSuite name="MyAppPackageTests.xctest.MyApp2Tests" tests="2" errors="0" failures="1" time="0.214"><testcase classname="MyApp2Tests.MyApp2Tests" name="testExample2" time="0.001"></testcase><testcase classname="MyApp2Tests.MyApp2Tests" name="testExample2" time="0.001"><failure type="XCTAssertFalse failed" message="test2">/Users/mike/tmp/swift-package-test/MyApp/Tests/MyAppTests/MyAppTests.swift:13</failure></testcase><system-out><![CDATA[test stdout test stdout]]></system-out></TestSuite>`,
		xmlString)
}
