package postgresql

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Export dumps PostgreSQL database with pg_dump into provided writer.
func Export(uri string, out io.Writer, verbosity int) error {
	args := []string{
		"docker", "compose", "exec", "postgresql", "pg_dump",
		"--clean", "--create", "--inserts",
	}

	if verbosity > 0 {
		args = append(args, "--verbose")
	}

	args = append(args, uri)

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run pg_dump: %s", err)
	}

	return nil
}
