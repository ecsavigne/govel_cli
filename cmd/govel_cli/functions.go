package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"

	"github.com/spf13/cobra"
)

var (
	repo             = "ecs_govel"
	token            = ""
	nameBranch       = "main"
	nameRemote       = "origin"
	repoSource       = ""
	userGit          = ""
	versionGovel_cli = "v1.1.12"
	versionEcs_govel = "v2.7.4"
)

func printVersion() {
	if versionFlag {
		fmt.Printf("version of govel_cli = \033[34m%s\033[0m and ecs_govel = \033[34m%s\033[0m\n", versionGovel_cli, versionEcs_govel)
		return
	}
}

func fValidateArgsOfRoot(cmd *cobra.Command, args []string) error {
	if len(args) == 1 && (args[0] == "execute" || slices.Contains(aliases, args[0])) {
		fmt.Println("entro govel argument 1")
		return nil
	} else if len(args) == 0 && (executeFlag == "execute" || slices.Contains(aliases, executeFlag)) {
		fmt.Println("entro govel flag execute argument 0")
		return nil
	} else if versionFlag {
		printVersion()
		return nil
	}

	return errors.New("govel_cli args not valid")
}

func validCantArgs(cmd *cobra.Command, args []string) error {
	switch {
	case cmd.Name() == "execute":
		if len(args) > 0 {
			return errors.New("command info not require args")
		}
	}

	return nil
}

func analizeFlags() string {
	find := slices.ContainsFunc(cmds, func(c *cobra.Command) bool {
		return slices.Contains(append([]string{c.Name()}, c.Aliases...), executeFlag)
	})
	if find {
		switch {
		case "execute" == executeFlag || slices.Contains(aliases, executeFlag):
			return executeFlag
		}
	} else {
		projectName = executeFlag
	}

	return ""
}

func fRunOfRoot(cmd *cobra.Command, args []string) {
	nameSubComand := ""
	var (
		subComand *cobra.Command
		e         error
	)

	if versionFlag {
		return
	}

	cantArg := len(args)
	if cantArg == 1 {
		subComand, _, e = cmd.Find([]string{args[0]})
		if e != nil {
			fmt.Println("Command ", args[0], " not found")
			return
		}
		nameSubComand = subComand.Name()
	} else {
		nameSubComand = executeFlag
	}

	if analizeFlags() != "execute" || !slices.Contains(aliases, analizeFlags()) {

	} else {
		nameSubComand = executeFlag
		if len(args) > 0 {
			cmd.Args(cmd, args)
		}
	}

	switch {
	case slices.Contains(aliases, nameSubComand), nameSubComand == "execute":
		executeCmd.Run(executeCmd, []string{projectName})
	default:
		fmt.Println("Command: ", args[0], " not found")
	}
}

func deleteFiles(nameFolder string) {
	fmt.Println("deleting files")
	path_name := path.Join(".", nameFolder, ".git")
	fmt.Printf("folder: %s\n", path_name)
	erros := os.RemoveAll(path_name)
	if erros != nil {
		fmt.Printf("Error deleting folder: %s, error: %s\n", path_name, erros.Error())
	}

	path_name = path.Join(".", nameFolder, "tmp")
	fmt.Printf("folder: %s\n", path_name)
	erros = os.RemoveAll(path_name)
	if erros != nil {
		fmt.Printf("Error deleting folder: %s, error: %s\n", path_name, erros.Error())
	}

	path_name = (path.Join(".", nameFolder, ".gitignore"))
	fmt.Printf("folder: %s\n", path_name)
	err := os.Remove(path_name)
	if err != nil {
		fmt.Printf("Error deleting file: %s, error: %s\n", path_name, erros.Error())
	}
}

func input(action string, val *string) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("\033[94mInsert %s: \033[0m\n", action)
	if reader.Scan() {
		*val = reader.Text()
	}
	fmt.Println("")
}

func inputToken() {
	input("Token", &token)
}

func inputOrigin() {
	input("Remote name", &nameRemote)
}

func inputBranch() {
	input("Branch name", &nameBranch)
}

func inputRepo() {
	fmt.Println("\033[31mCreate Repo in Git, only create\033[0m")
	input("Repo name", &repoSource)
}

func inputUser() {
	input("Git user name", &userGit)
}

func execCommand(cmds []string) error {
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()

}

func renameBranch() {
	gitRenameBranch := []string{
		"git",
		"branch",
		"-M",
		nameBranch,
	}

	execCommand(gitRenameBranch)
}

func firstCommit() {
	gitAdd := []string{
		"git",
		"add",
		".",
	}
	execCommand(gitAdd)

	gitCommit := []string{
		"git",
		"commit",
		"-m",
		"first commit",
	}
	execCommand(gitCommit)
}

func gitIgnoreCommit() {
	gitAdd := []string{
		"git",
		"add",
		".",
	}
	execCommand(gitAdd)

	gitCommit := []string{
		"git",
		"commit",
		"-m",
		"add .gitignore",
	}
	execCommand(gitCommit)
}

func pushInit() {
	gitPush := []string{
		"git",
		"push",
		"-u",
		nameRemote,
		nameBranch,
	}
	execCommand(gitPush)
}

func gitInit(nameFolder string) {
	err := os.Chdir(nameFolder)
	if err != nil {
		fmt.Println(err)
	}

	gitInit := []string{
		"git",
		"init",
	}
	execCommand(gitInit)
}

func addRemote() {
	gitAddRemote := []string{
		"git",
		"remote",
		"add",
		nameRemote,
		fmt.Sprintf("https://%s@github.com/%s/%s.git", token, userGit, repoSource),
	}
	execCommand(gitAddRemote)
}

func renameFolder(nameFolder string) {
	renameFolder := []string{
		"mv",
		"ecs_govel",
		nameFolder,
	}

	execCommand(renameFolder)
}

func setPermission(nameFolder string) {
	fmt.Println("set permission for ", nameFolder)
	_path := filepath.Join("..", nameFolder)
	setPermission := []string{
		"chmod",
		"-R",
		"777",
		_path,
	}

	execCommand(setPermission)
}

func createGitIgnore() {
	f, e := os.Create(".gitignore")
	if e != nil {
		fmt.Printf("error creating .gitignore: %s\n", e.Error())
	}
	defer f.Close()

	content := `
# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Code coverage profiles and other test artifacts
*.out
coverage.*
*.coverprofile
profile.cov

# Dependency directories (remove the comment below to include it)
vendor/

# Go workspace file
go.work
go.work.sum

# env file
.env

# Editor/IDE
# .idea/
# .vscode/
./test.json
./logs
./tmp/
./.vscode
.air.toml
./temp
`

	f.WriteString(content)
}

func createProject(nameFolder string) {
	gitClone := []string{
		"git",
		"clone",
		"https://github.com/ecsavigne/" + repo,
	}
	fmt.Println("Create project: ", nameFolder)

	if _, err := os.ReadDir(repo); err != nil {
		err := execCommand(gitClone)
		if err != nil {
			fmt.Println(err)
		}

		renameFolder(nameFolder)
		deleteFiles(nameFolder)
		// creatin repository
		inputRepo()
		inputUser()
		inputBranch()
		inputToken()
		inputOrigin()

		fmt.Printf("Repository: %s, User: %s, Branch: %s, Token: %s, Remote: %s\n", repoSource, userGit, nameBranch, token, nameRemote)

		gitInit(nameFolder)
		renameBranch()
		firstCommit()
		addRemote()
		pushInit()
		createGitIgnore()
		setPermission(nameFolder)
		gitIgnoreCommit()
		pushInit()
	} else {
	}

}
