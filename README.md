# Orcabak

A CLI tool to give version control and remote backups to Orca Slicer profiles,
powered by Git.

`orca` from Orca Slicer, `bak` from backup files.

## Table of Contents

- [Overview](#overview)
- [Project Status](#project-status)
- [Key Features](#key-features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Tab Completion](#installing-tab-completion-optional-but-recommended)
- [Quick Start](#quick-start)
- [Common Workflows](#common-workflows)
- [Usage](#usage)
- [Design and Architecture](#design-and-architecture)
- [Challenges and Learnings](#challenges-and-learnings)
- [Motivation](#motivation)
- [Limitations](#limitations)
- [Future Improvements](#future-improvements)
- [Contributing](#contributing)

## Overview

Orcabak is designed to give Orca Slicer users more control and protection over
their profile files. It provides Git-based version control and remote backups
without having to navigate through Orca Slicer's configuration directories.

Orca Slicer does not provide a built-in way to track changes to profiles.
Because profiles have such a large number of adjustable parameters, it can be
difficult to understand what changes were made over time or how they affect
print quality. In addition, profile files are stored deep within system specific
directories, which can make them difficult to locate and manage directly.

Orcabak solves this by treating the profile directory within the Orca Slicer
configuration as a Git repository and exposes a simplified CLI that mirrors
familiar Git workflows. This allows users to manage profiles from any directory
through a Git-like interface, while Orcabak orchestrates Git commands under the
hood, providing familiar features such as staging, committing, and syncing with
remote repositories.

## Project Status

Orcabak is under active development. It currently supports core Git workflows for
managing Orca Slicer profile files, with planned improvements in usability,
architecture, and feature completeness.

## Key Features

### Git Porcelain v2 Status Parsing

- Parses structured Git output instead of relying on raw CLI text
- Enables more controlled and predictable status handling

### Smart Push Behavior

- Automatically detects remote and current branch when possible
- Sets upstream on first push (-u) when appropriate
- Reduces required user input for common workflows

### Fast-Forward Safe Pulls

- Uses fast-forward-only strategy to avoid implicit merge commits
- Prevents unintended history changes
- Surfaces divergence clearly instead of masking it

### CLI Architecture with Dependency Injection

- Separation between CLI, Git client, and command execution
- Designed with testability and future extensibility in mind
- Dependencies are created through factory functions for dependency injection
- Shared dependencies are managed through a centralized state struct

## Technologies Used

- Go
- Cobra (CLI framework)
- Viper (configuration management)
- Git (via shell execution)

## Installation

### Prerequisites

- Go (for installation with `go install`)
- Git (must be installed and available on PATH for Orcabak to execute Git
commands)
- Orca Slicer (or access to an existing Orca Slicer config directory)

> Orcabak is currently tested against Orca Slicer v2.3.1. Compatibility with
> other versions is expected but not guaranteed, as Orcabak depends on Orca
> Slicer's configuration directory structure.

### Install via Go

Orcabak is currently distributed via `go install`:

```sh
go install github.com/sbrown3212/orcabak/cmd/orcabak@latest
```

Future releases will include precompiled binaries for common platforms.

## Installing Tab Completion (optional but recommended)

> Tab completion is especially helpful when working with long Orca Slicer
profile filenames.

Installing the completion script enables Orcabak tab completion functionality to
all command and subcommands, as well as file name completion to the `add`
command, along with config option names with all config sub commands. Tab
completion will be coming soon for all command arguments.

### macOS/Linux

> The default shell in macOS is `Zsh`. If unsure which shell you are using,
> enter `echo $SHELL` to find out.

#### Bash

1. Generate completions script

    ```bash
    orcabak completion bash > ~/.orcabak_completion.bash
    ```

2. Load script in shell

    Add this to your `~/.bashrc` (or `~/.bash_profile` on macOS)

    ```bash
    source ~/.orcabak_completion.bash
    ```

3. Reload shell

    ```bash
    source ~/.bashrc
    ```

#### Zsh

1. Create completions directory (if it doesn't exist)

    ```zsh
    mkdir -p ~/.zsh/completions/
    ```

2. Generate completion script

    ```zsh
    orcabak completion zsh > ~/.zsh/completions/_orcabak
    ```

3. Add the directory to your `fpath`

    Add this line to your `~/.zshrc`:

    ```zsh
    fpath=(~/.zsh/completions $fpath)
    ```

4. Initialize completion (add to end of `~/.zshrc` if not already present)

    ```zsh
    autoload -Uz compinit
    compinit
    ```

5. Reload shell

    ```zsh
    source ~/.zshrc
    ```

#### Fish

1. Generate completions script

    ```fish
    orcabak completion fish > ~/.config/fish/completions/orcabak.fish
    ```

2. Restart shell

    ```fish
    exec fish
    ```

### Windows

#### PowerShell

##### Quick Test (current session only)

1. Generate completion script

    ```PowerShell
    orcabak completion powershell | Out-String | Invoke-Expression
    ```

##### Persistent Setup

1. Open PowerShell profile in editor

    ```PowerShell
    notepad $PROFILE
    ```

2. Generate completion script

    ```PowerShell
    orcabak completion powershell | Out-String | Invoke-Expression
    ```

### Verify Installation

After installation, try:

```sh
orcabak [TAB]
```

If installed successfully, the available commands and flags should appear in the
command prompt.

## Quick Start

Orcabak closely mimics Git. If unfamiliar with basic Git concepts
(staging, commits, remote repositories, etc.), have a look at
[this](https://www.freecodecamp.org/news/learn-the-basics-of-git-in-under-10-minutes-da548267cc91/)
article from FreeCodeCamp.

In order to push, or back up, Orca Slicer profiles (optional but recommended) to
GitHub (or another Git remote), create a new repository and
copy the URL to use in the `remote add` command below.

```sh
# Initialize Orca Slicer profile directory as a git repository
orcabak init

# Stage files
orcabak add .

# Commit changes
orcabak commit "Initial backup"

# Add remote repository
orcabak remote add origin <repo-url>

# Push to remote
orcabak push
```

## Common Workflows

### Initial Setup

```sh
# Initialize repository
orcabak init

# (Optional) Define location of Orca Slicer config directory
orcabak config set --orca-cfg-path "path/to/OrcaSlicer"
```

> Setting the Orca Slicer config path is not required for most. If you have not
> customized the location of Orca Slicer's config directory, Orcabak will
> default to the `OrcaSlicer/` directory in your operating system's default
> location for application config directories.
>
> When specifying the Orca Slicer config directory, be sure to give the path to
> the `OrcaSlicer/` directory, and NOT to `OrcaSlicer/user/default` (where the
> profile files are stored and where Orcabak will create the Git repository)

### Saving Changes

```sh
# Stage files
orcabak add <file_names>

# Commit changes
orcabak commit "Update profiles"
```

### Pushing to Remote

```sh
# Add a remote (if not already set)
orcabak remote add origin <repo-url>

# Push changes
orcabak push
```

```sh
# if multiple remotes exist
orcabak push <remote_name>
```

### Restoring on a New Machine

```sh
# Add a remote (if not already set)
orcabak remote add origin <repo-url>

# Pull profiles from remote
orcabak pull <remote_name> <remote_branch>
```

```sh
# If upstream is already set
orcabak pull
```

## Usage

### Global Flags

- `-v`, `--verbose` - enable verbose output

### Core Commands

- `init` - initialize a Git repository in the Orca Slicer config directory
(`OrcaSlicer/user/default/`)
- `status` - show repository status
- `add` - stage changes
- `commit` - commit staged changes
- `push` - push commits to a remote
- `pull` - pull changes from a remote (fast-forward only)

### Config

- `config get` - retrieve a config value
- `config set` - set a config value
- `config unset` - remove a config value
- `config list` - list all config values

### Remote Management

- `remote add` - add a remote repository
- `remote remove` - remove a remote repository

## Design and Architecture

### High-Level Structure

```markdown
CLI (cobra)
  ↓
Git Client
  ↓
Command Runner (executes git commands)
```

### Design Notes

- CLI layer currently contains business logic (planned for refactor)
- Git interactions implemented via shelling out
- Viper used for configuration management
- Dependency injection used to construct commands and dependencies

### Why This Matters

- Improves testability
- Enables future separation into application layer
- Keeps Git concerns isolated

## Challenges and Learnings

### Git Status Parsing

Parsing Git status output was by far the most difficult but crucial part of the
project. Many Orcabak features rely on the output that comes from the `status`
command. When diving into Git's documentation, I came across the "short" status
output, and then the porcelain standards. Porcelain v2 was chosen because it
was more detailed and robust than v1 and "short" output. This made parsing the
status output within Orcabak significantly easier.

### Abstracting Git vs Mimicking Git

Initially I planned on abstracting Git to make the tool easier to use for those
that are less familiar. This turned out to be quite challenging, and I pivoted
to mimic Git instead. Mimicking Git's core commands allowed me to create the
backbone of Orcabak before trying to implement more advanced features and
deciding how things should be abstracted.

### Project Complexity

I underestimated how difficult it would be to create a simple Git wrapper.
Because of this, there were a few features left out. Instead of building out the
Git-like commands to be as feature full as Git's actual commands, I had to focus
on prioritizing what Orcabak users would need most.

## Motivation

I had installed Klipper on my old Creality Ender 3 a few years back, and I had
switched to using Orca Slicer around the same time. The print quality was
initially pretty good, but slowly went down hill. But, I didn't know why. So I
wanted to redo everything software related to my 3D printer. But this time, I
wanted version control.

Using Git to track my printer's Klipper config files was not too difficult.
Klipper is just a firmware that runs on Linux, so I initialized the
`printer-config/` directory as a git repo and pushed it to a GitHub repo.

The same could definitely be done with Orca Slicer profile files, but these are
edited a lot more frequently, at least in my experience. As these particular
files seemed to be nested deeply within the directories of my computer, I
was never going to remember the path in order to use Git.

I was also interested in learning more about how tooling is built, and wondered
how difficult it would be to make a tool that implemented the basics of Git, but
had Orca Slicer specific output. So, in order to avoid having an extra sticky
note on my monitor, and to save a few minutes every time I wanted to commit a
change, I spent several times more time and effort to make a tool that would
make this easier.

It turns out that the deterioration in print quality may have been due more to
the aging hardware rather than poor slicer profiles. But I am glad to have taken
this deep dive into better understanding how Git works, writing tooling with Go,
and learning more about software architecture. And I am not finished with
Orcabak, there are many features and improvements yet to come.

## Limitations

- `pull` only supports fast-forward (no merge/rebase handling)
- `commit` requires message via CLI (no editor support)
- Limited flag support compared to Git
- Minimal automated testing outside of status parsing
- Some error messages still surface raw Git output

## Future Improvements

- Introduce application layer
- Improve error handling with domain specific errors
- Expand `pull` to support merge/rebase
- Add branch management commands
- Implement `diff` and `log` commands
- Add editor based commit messages
- Improve tab completion across commands
- Improve usability for users unfamiliar with Git
- Official binary releases

## Contributing

### Fork and Clone the repo

Fork the repository on GitHub, then clone your fork locally:

```sh
git clone https://github.com/sbrown3212/orcabak
cd orcabak
```

### Build the compiled binary

```sh
go build ./cmd/orcabak
```

### Run the test suite

```sh
go test ./...
```

### Submit a pull request

Open a pull request from your fork to the `main` branch of this repository.
