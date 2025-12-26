package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"slices"

	"github.com/spf13/cobra"
)

var (
	repo       = "ecs_govel"
	token      = ""
	nameBranch = "main"
	nameRemote = "origin"
	repoSource = ""
	userGit    = ""
)

func fValidateArgsOfRoot(cmd *cobra.Command, args []string) error {
	if len(args) == 1 && (args[0] == "execute" || slices.Contains(aliases, args[0])) {
		fmt.Println("entro govel argument 1")
		return nil
	} else if len(args) == 0 && (executeFlag == "execute" || slices.Contains(aliases, executeFlag)) {
		fmt.Println("entro govel flag execute argument 0")
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
	fmt.Println("Project already exists")
	erros := os.RemoveAll(path.Join(".", nameFolder, ".git"))
	if erros != nil {
		fmt.Println("Error deletando carpeta .git: ", erros.Error())
	}

	erros = os.RemoveAll(path.Join(".", nameFolder, "tmp"))
	if erros != nil {
		fmt.Println("Error deletando carpeta .git: ", erros.Error())
	}

	err := os.Remove((path.Join(".", nameFolder, ".gitignore")))
	if err != nil {
		fmt.Println("Error deletando .gitignore: ", err.Error())
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

		deleteFiles(nameFolder)
	} else {
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
	}

}
