package main

import "fmt"
import "os"
import "strings"

const  (
	STDOUT = 1
	STDERR = 2
)

var separator string = "--"

func main() {
	fmt.Println("Starting...")
	lists := partition(separator, os.Args[1:])
	opts := lists[0]
	cmds := lists[1:]

	out := make(chan msg, 100)

	procs := []process{}

	for _, cmd := range cmds {
		p := newProc(out, cmd)
		procs = append(procs, p)
		p.start()
	}

	for {
		select {
		case m := <-out:
			printOut(m)
		}
	}

	fmt.Printf("options %v\n", opts)
	fmt.Printf("   cmds %v\n", cmds)
	fmt.Printf("  procs %v\n", procs)

}

func printOut(m msg) {
	text := strings.TrimRight(m.text, "\n")
	label := "UNKOWN"
	switch m.fd {
	case STDERR:
		label = "ERR"
	case STDOUT:
		label = "OUT"
	default:
		label = "UNK"
	}
	fmt.Printf("[%v %v] %v\n", label, m.proc.Name, text)
}
func partition(sep string, opts []string) [][]string {
	sets := [][]string{}
	set := []string{}

	for len(opts) != 0 {
		o := opts[0]
		opts = opts[1:]
		if o == sep {
			sets = append(sets, set)
			set = []string{}
		} else {
			set = append(set, o)
		}
	}
	sets = append(sets, set)

	return sets
}
