package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joshuarubin/go-sway"
)

func main() {
	prev := flag.Bool("prev", false, "focus previous window")
	flag.Parse()

	ctx := context.Background()

	client, err := sway.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	tree, err := client.GetTree(ctx)
	if err != nil {
		log.Fatal(err)
	}

	container, focused := findFocusedContainer(tree, tree)
	if container == nil || focused == nil || len(container.Nodes) < 2 {
		fmt.Fprintln(os.Stderr, "no container found or insufficient windows")
		os.Exit(0)
	}

	var target *sway.Node
	var offset int

	if *prev {
		offset = -1
	} else {
		offset = 1
	}

	for i, node := range container.Nodes {
		if node.ID == focused.ID {
			target = container.Nodes[(i+offset+len(container.Nodes))%len(container.Nodes)]
			break
		}
	}

	if target != nil {
		if _, err := client.RunCommand(ctx, fmt.Sprintf("[con_id=%d] focus", target.ID)); err != nil {
			log.Fatal(err)
		}
	}
}

func findFocusedContainer(parent, child *sway.Node) (*sway.Node, *sway.Node) {
	if child.Focused {
		return parent, child
	}

	for _, node := range child.Nodes {
		if p, c := findFocusedContainer(child, node); p != nil && c != nil {
			return p, c
		}
	}

	for _, node := range child.FloatingNodes {
		if p, c := findFocusedContainer(child, node); p != nil && c != nil {
			return p, c
		}
	}

	return nil, nil
}
