# A Forest Full of Trees
Note: this is taken from a [blog post on my website](https://williamcook.dev/posts/botanical_exploration_of_abstract_types_in_go/). Stop by there if you want some other fun explorations of Go related topics. 


Are you a Go developer who wants a solid reference for how interfaces and struct composition work? Look no further! Inside this repo are a few submodules you might find handy. Feel free to open up that repo, but we will be referencing code from it here directly. 

## Inter-forest-aces
An interface in Go is an abstract type. That is, it describes something without actually embodying the thing itself. Interfaces take this shape:
```
type someIFace interface {
    someMethod() string
    someOtherMethod() int
}
```

Notice that in `someIFace`, there are two methods. One returns a string, and another an int. This is helpful if you want a function that operates on types, but you want to adapt what those types are situationally. To put it into context, take the interface from `package trees`.

```
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
```

The interface says _nothing_ about any particular tree. It just defines how a concrete tree will describe itself. Some of these methods are _implemented_ (that is to say, these methods exist) on the concrete type `Tree`. 

```
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
```

However, since `Tree` doesn't implement every method of `ITree`, it does not _satisfy_ the interface. 

## Growing some deeper roots 
While a `Tree` implements some of the methods that `ITree` requires, the methods related to the seaons are left unsatisfied. This is where both packages `deciduous` and `conifer` come into play. 

### Planting a seed (or a struct)
If you come from another language, you might be familiar with _inheritance_. This is the idea that something (usually a _class_) can pass down properties to other things (often called their children). The thing receiving these properties is _inheriting_ those from its parent. Go doesn't strictly have this, but we have something close called _embedding_. When we embed a struct into another, all of the properties of the outer struct are passed down to the struct doing the embedding. Here, both `Conifer` and `Deciduous` embed the `Trees` struct. 

```
package conifer

// Confier is type of tree. All fields are
// composed of the Tree type from the trees package.
type Conifer struct {
	trees.Tree
}
```

```
package deciduous

// Deciduous is type of tree. All fields are
// composed of the Tree type from the trees package.
type Deciduous struct {
	trees.Tree
}
```

Since the `trees.Tree` struct has methods like `GetCommonName` and properties like `Genus` and `Species`, these two structs do as well. 


## A Fork in the Road
Conifers and deciduous trees have different lifecyles through. And they experience the turning of the seasons in very different ways. So we would not want them to receive those methods from a common `Tree` ancestor. This is why we implement these methods separately for these two different categories of tree. For example:

```
package deciduous 

// Spring returns a string representing the
// action deciduous trees take during spring.
func (d Deciduous) Spring() string {
	return "Ahhh, new leaves and flowers!"
}
```

```
package conifer

// Spring returns a string representing the
// action conifers take during spring.
func (c Conifer) Spring() string {
	return "Growing some cones!"
}
```
Both of these methods reveal how both differing kinds of trees experience the same thing in their own special way. 

## Bring it around the campfire
As previously mentioned, the real power of an abstract type is to pass it to a function that operates on interfaces. Take a look at the `trees.Lifecycle` function. 

```
package trees 

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
```
`Lifecycle` does one thing - it progresses through a year of whatever `ITree` is passed into it. It does not care whether the `ITree` is deciduous, a conifer, or a new, third type! But rather than having `deciduous.Lifecycle` and `conifer.Lifecycle`, you can have this one place to make these `ITree`s grow. Now how neat is that? 

## The wilderness must be explored!
Have some fun with this. Create some new types of trees that implement the `ITree` interface. Or maybe rename `ITree` to `IPlant` and create some `vascular` and `nonvascular` plant types and let `deciduous` and `conifer` embed those structs as necessary. Test yourself too. Knowing now what you do about interfaces, do you know why `trees.New()` can't be passed into the `trees.Lifecycle()` function, but both a `conifer.New()` and a `deciduous.New()` can?
