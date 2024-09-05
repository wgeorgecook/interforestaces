package trees

import "fmt"

// ITree describes an abstract tree
type ITree interface {
	GetCategory() string
	GetGenus() string
	GetSpecies() string
	GetCommonName() string
	Spring() string
	Summer() string
	Fall() string
	Winter() string
}

// Lifecycle iterates over the methods found in ITree to describe
// the argument Tree and it's life over one year.
func Lifecycle(t ITree) {
	fmt.Println("Hello! I am " + t.GetGenus() + " " + t.GetSpecies() + " (" + t.GetCommonName() + "), " + t.GetCategory() + " tree!")
	fmt.Println("Today I am going through my lifecycle. Come grow with me!")
	fmt.Println("It's Spring! " + t.Spring())
	fmt.Println("Summertime! " + t.Summer())
	fmt.Println("Autumn is here. " + t.Fall())
	fmt.Println("Brrrr winter! " + t.Winter())
}

// Tree is a concrete structure to build a Tree with.
type Tree struct {
	genus, species, commonName string
}

// GetGenus returns the genus of the receiving tree.
func (t Tree) GetGenus() string {
	return t.genus
}

// GetSpecies returns the species of the receiving tree.
func (t Tree) GetSpecies() string {
	return t.species
}

// GetCommonName returns the common name of the receiving tree.
func (t Tree) GetCommonName() string {
	return t.commonName
}

// New returns a new instance of Tree.
func New(genus, species, commonName string) Tree {
	return Tree{
		genus:      genus,
		species:    species,
		commonName: commonName,
	}
}
