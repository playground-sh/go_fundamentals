package code_organization

import "fmt"

type cpu struct {
	brand string
	cores int
}

type gpu struct {
	model        string
	vram         int
	manufacturer string
}

type computer struct {
	processor cpu
	graphics  gpu
	brand     string
}

func (c *computer) showSpecs() {
	fmt.Println("Computer Specs")
	fmt.Println("---------------")
	fmt.Println("Brand:", c.brand)
	fmt.Printf("CPU: %s (%d cores)\n", c.processor.brand, c.processor.cores)
	fmt.Printf("GPU: %s %s (%d GB VRAM)\n",
		c.graphics.manufacturer, c.graphics.model, c.graphics.vram)
}

func CompositionDemo() {
	pc := computer{
		processor: cpu{"Core Ultra 285k", 24},
		graphics: gpu{
			model:        "RTX 5090",
			vram:         32,
			manufacturer: "MSI",
		},
		brand: "Custom Build",
	}
	pc.showSpecs()
}
