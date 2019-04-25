package main

import (
	"os/exec"
	"io"
	"os"
	"fmt"
	"bufio"
)

type process struct {
	Name   string
	command *exec.Cmd
	out     chan msg
}

type msg struct {
	proc   *process
	fd    int
	text   string
}

func newProc(output chan msg, args []string) process {
	return process{
		Name:    args[0],
		command: exec.Command(args[0], args[1]),
		out:     output,
	}
}


func (p *process) start() {
	stdin, err := p.command.StdinPipe()
	if err != nil {
		errx("closing stdin for %v\n", p.Name)
	}
	stdin.Close()
	stdout, err := p.command.StdoutPipe()
	if err != nil {
		errx("closing stdout for %v\n", p.Name)
	}
	stderr, err := p.command.StderrPipe()
	if err != nil {
		errx("closing stderr for %v\n", p.Name)
	}
	go p.forward(STDOUT, stdout)
	go p.forward(STDERR, stderr)
	p.command.Start()
}

func (p *process) forward(fd int, stream io.ReadCloser) {
	 buf := bufio.NewReader(stream)
	 for {
		 line, err := buf.ReadString('\n')
		 if err != nil {
			 errx("reading line\n")
		 }
		 p.out <- msg{p, fd, line}
	 }
}

func errx(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

