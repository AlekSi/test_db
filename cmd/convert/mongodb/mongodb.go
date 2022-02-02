package mongodb

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

// Import reads documents from the provided reader and writes them into MongoDB with `mongoimport`.
func Import(uri, collection string, docs io.Reader, verbosity int) error {
	args := []string{
		"docker", "compose", "exec", "mongodb", "mongoimport",
		"--uri=" + uri,
		"--collection=" + collection,
		"--drop",
		"--maintainInsertionOrder",
	}

	if verbosity < 0 {
		args = append(args, "--quiet")
	} else {
		args = append(args, "--verbose="+strconv.Itoa(verbosity))
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = docs
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run mongoimport: %s", err)
	}

	return nil
}

// Export reads documents from MongoDB with `mongoexport` and writes them to the provided writer.
func Export(uri, collection string, docs io.Writer, verbosity int) error {
	args := []string{
		"docker", "compose", "exec", "mongodb", "mongoexport",
		"--uri=" + uri,
		"--collection=" + collection,
		"--jsonFormat=canonical",
	}

	if verbosity < 0 {
		args = append(args, "--quiet")
	} else {
		args = append(args, "--verbose="+strconv.Itoa(verbosity))
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = docs
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run mongoexport: %s", err)
	}

	return nil
}
