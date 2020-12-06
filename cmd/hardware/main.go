package main

import (
	"github.com/micahkemp/scad-hardware/pkg/hardware"
	"github.com/micahkemp/scad/pkg/scad"
)

func main() {
	scad.CLIHandler(hardware.Models)
}
