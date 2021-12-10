package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes int
	name     string
}

var errPosArgSpecified = errors.New("Positional arguments specified")

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "your name please? press the enter key when done.\n"
	fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}
	return name, nil
}

func parseArgs(w io.Writer, args []string) (config, error) {

	c := config{}
	fs := flag.NewFlagSet("greeteer", flag.ContinueOnError)
	fs.SetOutput(w)
	// fs.Usage = func() {

	// 	var usageString = `
	// 	A greeter application which prints the name you entered <integer> number of times.

	// 	Usage of %s: <option> [name]`,
	// 		fmt.Fprintf(w, usageString, fs.Name())
	// 	fmt.Fprintln(w)
	// 	fmt.Fprintln(w, "Options: ")
	// 	fs.PrintDefaults()
	// }

	fs.IntVar(&c.numTimes, "n", 0, "Number of times of greet") // make option
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, errInvaliadPosArgSpcified
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	return c, nil

	//var numTimes int
	//var err error

	// if len(args) != 1 {
	// 	return c, errors.New("Invalid number of arguments")
	// }

	// if args[0] == "-h" || args[0] == "--help" {
	// 	c.printUsage = true
	// 	return c, nil
	// }

	// numTimes, err = strconv.Atoi(args[0])
	// if err != nil {
	// 	return c, err
	// }
	// c.numTimes = numTimes
	// return c, nil
}

// func printUsage(w io.Writer) {
// 	fmt.Fprintf(w, usageString)
// }

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(r, w)
		if err != nil {
			return err
		}
	}
	greetUser(c, w)
	return nil
}

var errInvaliadPosArgSpcified = errors.New("More than one positional argument specified")

func greetUser(c config, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", c.name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		if errors.Is(err, errPosArgSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		// printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
