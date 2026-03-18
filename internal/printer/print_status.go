package printer

import (
	"fmt"

	"github.com/sbrown3212/orcabak/internal/git"
)

func (p *Printer) PrintStatus(status git.GitStatus) error {
	if status.Branch.Name != "" {
		if _, err := fmt.Fprintln(p.Out, "On branch:", status.Branch.Name); err != nil {
			return err
		}
	}

	if err := p.printFileSection("Staged", status.Staged); err != nil {
		return err
	}

	if err := p.printFileSection("Unstaged", status.Unstaged); err != nil {
		return err
	}

	if err := p.printStringSection("Untracked", status.Untracked); err != nil {
		return err
	}

	if err := p.printStringSection("Conflicts", status.Conflicts); err != nil {
		return err
	}

	return nil
}

func (p *Printer) printFileStatus(fs git.FileStatus) error {
	if fs.OrigPath != "" {
		_, err := fmt.Fprintf(p.Out, "    %s: %s -> %s\n", fs.Type, fs.OrigPath, fs.Path)
		if err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprintf(p.Out, "    %s: %s\n", fs.Type, fs.Path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Printer) printNewLine() error {
	_, err := fmt.Fprintln(p.Out)
	if err != nil {
		return err
	}

	return nil
}

func (p *Printer) printFileSection(title string, files []git.FileStatus) error {
	if len(files) == 0 {
		return nil
	}

	if _, err := fmt.Fprintln(p.Out, title, ":"); err != nil {
		return err
	}

	for _, fs := range files {
		if err := p.printFileStatus(fs); err != nil {
			return err
		}
	}

	return p.printNewLine()
}

func (p *Printer) printStringSection(title string, strings []string) error {
	if len(strings) == 0 {
		return nil
	}

	if _, err := fmt.Fprintln(p.Out, title, ":"); err != nil {
		return err
	}

	for _, str := range strings {
		if _, err := fmt.Fprintln(p.Out, "    ", str); err != nil {
			return err
		}
	}

	return p.printNewLine()
}
