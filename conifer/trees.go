package conifer

import "interforestaces/trees"

// Confier is type of tree. All fields are
// composed of the Tree type from the trees package.
type Conifer struct {
	trees.Tree
}

// GetCategory returns "an evergreen", as all
// conifers are evergreen trees.
func (c Conifer) GetCategory() string {
	return "an evergreen"
}

// Spring returns a string representing the
// action conifers take during spring.
func (c Conifer) Spring() string {
	return "Growing some cones!"
}

// Summer returns a string representing the
// action conifers take during summer.
func (c Conifer) Summer() string {
	return "Dropping those cones!"
}

// Fall returns a string representing the
// action conifers take during fall.
func (c Conifer) Fall() string {
	return "Dropping some needles but it's not bad."
}

// Winter returns a string representing the
// action conifers take during winter.
func (c Conifer) Winter() string {
	return "Nothing to do for us conifers."
}

// New returns a new instance of Conifer.
func New(genus, species, commonName string) Conifer {
	return Conifer{
		Tree: trees.New(genus, species, commonName),
	}
}
