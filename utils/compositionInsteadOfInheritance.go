package utils

// Composition by embedding structs
// Composition can be achieved in Go is by embedding one struct type into another.

// A blog post is a perfect example of composition. Each blog post has a title, content and author information. This can be perfectly represented using composition. In the next steps of this tutorial, we will learn how this is done.

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName())
	fmt.Println("Bio: ", b.bio)
}

type website struct {
	blogPost []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.blogPost {
		v.details()
		fmt.Println()
	}
}

func CompositionByEmbeddingStructs() {
	author1 := author{
		"Sams",
		"Recs",
		"Golang Enthusiast",
	}
	blogPost1 := blogPost{
		"Inheritance",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go supports composition instead of inheritance",
		author1,
	}
	w := website{
		blogPost: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()

}
