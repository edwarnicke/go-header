package messages

import (
	"errors"
	"fmt"
	"strings"

	"github.com/denis-tingajkin/go-header/text"
	"github.com/denis-tingajkin/go-header/utils"
)

//ErrorMsg returns const string if err is nil otherwise returns formated error
func ErrorMsg(err error) string {
	if err == nil {
		return "<error is nil>"
	}
	return fmt.Sprintf("Error: %v", utils.AddTabPerLine(err.Error()))
}

func CatNotParseAsYear() error {
	return errors.New("can not parse as year")
}

func UnknownField(f string) error {
	return fmt.Errorf("unknown field: %v", f)
}

func MasterBranchCanNotBeEmpty() error {
	return errors.New("master branch can not be empty if scope.policy not equals none")
}

func DetectedInfiniteRecursiveEntry(entries ...string) error {
	return fmt.Errorf("detected infinite recursive entry: \"%v\"", strings.Join(entries, " -> "))
}

func Ambiguous(first ErrorList, second ErrorList) error {
	firstText := first.String()
	secondText := second.String()
	return fmt.Errorf("ambiguous parser error:\nCase 1:\n %v\nCase 2:\n%v", utils.AddTabPerLine(firstText), utils.AddTabPerLine(secondText))
}

func UnknownCopyrightHolder(position int, holder string, expectedHolders ...string) error {
	expected := strings.Join(expectedHolders, ", ")
	if expected == "" {
		expected = "any not empty string"
	}
	return fmt.Errorf("unknown copyright holder: \"%v\" at position %v. Expected: \"%v\"", holder, position, expected)
}

func CopyrightHolderAlreadyInUse(holder string) error {
	return fmt.Errorf("copyright holder \"%v\" already in use", holder)
}

func CanNotLoadTemplateFromFile(reason error) error {
	return fmt.Errorf("can not load template file. %v", ErrorMsg(reason))
}

func NoRules() error {
	return errors.New("no rules defined")
}

func IncorrectGoroutineCount(actual int) error {
	return fmt.Errorf("incorrect goroutine count. Actual: \"%v\". Expected: value should be more than zero", actual)
}

func CantProcessField(name string, err error) error {
	return fmt.Errorf("can not process field: \"%v\". %v", name, ErrorMsg(err))
}

func TemplateNotProvided() error {
	return errors.New("template not provided")
}

func UnknownPattern(patternName string) error {
	return fmt.Errorf("template: unknown pattern %v", patternName)
}

func VerifyFuncNotProvided() error {
	return errors.New("verify func not provided")
}

func WrongYear() error {
	return errors.New("expects the interval from the year the file was created to the current year")
}
func AnalysisError(location text.Location, reason error) error {
	return fmt.Errorf("position: %v, %v", location, ErrorMsg(reason))
}

func Diff(actual, expected interface{}) error {
	actualText := utils.AddTabPerLine(fmt.Sprint(actual))
	expectedText := utils.AddTabPerLine(fmt.Sprint(expected))
	return fmt.Errorf("\nexpected:\n%v\nactual:\n%v", expectedText, actualText)
}

func Missed(what string) error {
	return fmt.Errorf("missed: %v", utils.AddTabPerLine(what))
}

func NotExpected(what string) error {
	return fmt.Errorf("not expected text: %v", utils.AddTabPerLine(what))
}

func NoHeader() error {
	return errors.New("file has not license header")
}
