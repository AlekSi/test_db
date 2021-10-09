package mongodb

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Import(uri, collection string, docs io.Reader) error {
	args := []string{
		"docker", "compose", "exec", "mongodb", "mongoimport",
		"--uri=" + uri,
		"--collection=" + collection,
		"--drop",
		"--maintainInsertionOrder",
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = docs
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run mongoimport: %s", err)
	}

	return nil
}

func Export(uri, collection string, docs io.Writer) error {
	args := []string{
		"docker", "compose", "exec", "mongodb", "mongoexport",
		"--uri=" + uri,
		"--collection=" + collection,
		"--jsonFormat=canonical",
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = docs
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run mongoexport: %s", err)
	}

	return nil
}
