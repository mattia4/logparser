package cli

import (
	"fmt"
	"os"
)

func CheckArgs() (string, error) {

	cliErrorUseCase := fmt.Errorf("use case: %s <file_path>", os.Args[0])

	if len(os.Args) < 2 {
		return "", cliErrorUseCase
	}

	return os.Args[1], nil
}

func GetInputFilePathOrError(errorHandler func(err error)) string {
	err := fmt.Errorf("use case: %s <file_path>", os.Args[0])

	if len(os.Args) < 2 {
		errorHandler(err)
	}

	return os.Args[1]
}
