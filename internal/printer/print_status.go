package printer

import "github.com/sbrown3212/orcabak/internal/git"

func (p *Printer) PrintStatus(status git.GitStatus) error {
	if status.Branch.Name != "" {
		p.Println("On branch:", status.Branch.Name)
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
		p.Printf("    %s: %s -> %s\n", fs.Type, fs.OrigPath, fs.Path)
	} else {
		p.Printf("    %s: %s\n", fs.Type, fs.Path)
	}

	return nil
}

func (p *Printer) printNewLine() error {
	p.Println()

	return nil
}

func (p *Printer) printFileSection(title string, files []git.FileStatus) error {
	if len(files) == 0 {
		return nil
	}

	p.Println(title + ":")

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

	p.Println(title + ":")

	for _, str := range strings {
		p.Println("    ", str)
	}

	return p.printNewLine()
}
