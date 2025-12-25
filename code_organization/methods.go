package code_organization

import "fmt"

type medicine struct {
	name  string
	price float32
}

type prescription struct {
	meds       []string
	totalPrice float32
}

func (p *prescription) addMedicine(med medicine) {
	p.meds = append(p.meds, med.name)
	p.totalPrice += med.price
}

func (p *prescription) print() {
	fmt.Println("Current medicines:", p.meds)
	fmt.Println("Total cost:", p.totalPrice)
}


func MethodsDemo() {
	paracetamol := medicine{
		name:  "Paracetamol",
		price: 10.5,
	}

	piriton := medicine{
		name:  "Piriton",
		price: 6,
	}

	var p prescription
	p.addMedicine(paracetamol)
	p.addMedicine(piriton)
	p.print()
}
