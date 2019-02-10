package main

type Drawer struct {
	prompt string
}

func NewDrawer(prompt string) *Drawer {
	d := &Drawer{
		prompt: prompt,
	}
	return d
}
