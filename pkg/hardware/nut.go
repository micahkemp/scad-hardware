package hardware

import (
	"github.com/micahkemp/scad-extras/pkg/extras"
	"github.com/micahkemp/scad/pkg/scad"
)

type Nut struct {
	Name       string
	FlatsWidth float64
	Thickness  float64
}

var ExampleNut = Nut{
	Name:       "example_nut",
	FlatsWidth: 10,
	Thickness:  2,
}

func (n Nut) SCADWriter() scad.SCADWriter {
	return scad.LinearExtrude{
		Name: scad.Name{
			Specified: n.Name,
			Default:   "nut",
		}.String(),
		Height: n.Thickness,
		Children: []scad.Module{
			extras.Hexagon{
				Apothem: n.FlatsWidth / 2,
			},
		},
	}.SCADWriter()
}
