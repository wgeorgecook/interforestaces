package deciduous

import (
	"interforestaces/trees"
)

// Deciduous is type of tree. All fields are
// composed of the Tree type from the trees package.
type Deciduous struct {
	trees.Tree
}

// GetCategory returns "a deciduous", this
// category of tree.
func (d Deciduous) GetCategory() string {
	return "a deciduous"
}

// Spring returns a string representing the
// action deciduous trees take during spring.
func (d Deciduous) Spring() string {
	return "Ahhh, new leaves and flowers!"
}

// Summer returns a string representing the
// action deciduous trees take during summer.
func (d Deciduous) Summer() string {
	return "We are fully leafed out! Growing some fruit too."
}

// Fall returns a string representing the
// action deciduous trees take during fall.
func (d Deciduous) Fall() string {
	return "Nights are getting long let's drop these leaves."
}

// Winter returns a string representing the
// action deciduous trees take during winter.
func (d Deciduous) Winter() string {
	return "Shhhh we are mostly dormant."
}

// New returns a new instance of Deciduous.
func New(genus, species, commonName string) Deciduous {
	return Deciduous{
		Tree: trees.New(genus, species, commonName),
	}
}
