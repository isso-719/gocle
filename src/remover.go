package src

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func confirm(cmd *cobra.Command, args []string) bool {
	if cmd.Flag("yes").Value.String() == "true" {
		return true
	}

	fmt.Println("Are you sure you want to remove the following package(s)? [y/n]")
	for _, pkg := range args {
		fmt.Printf(" - %s\n", pkg)
	}
	fmt.Print("> ")

	var answer string
	_, err := fmt.Scan(&answer)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if answer != "y" {
		fmt.Println("Canceled.")
		return false
	}

	return true
}

func remove(args []string) {
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gobin = os.Getenv("GOPATH") + "/bin"
	}
	if gobin == "" {
		log.Fatal("GOBIN or GOPATH is not set.")
		return
	}
	for _, pkg := range args {
		err := os.Remove(gobin + "/" + pkg)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("Removed %s\n", pkg)
	}
}

func Remover(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if confirm(cmd, args) {
		remove(args)
	}
}
