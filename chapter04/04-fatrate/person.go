package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Person struct {
	name string
	sex string
	tall float64
	weight float64
	age int
	bmi float64
	fatRate float64
}

func (p *Person) calcBmi() error{
	bmi,err := gobmi.BMI(p.weight,p.tall)
	if err != nil {
		log.Printf("error BMI for Person[%s]: %v",p.name,err)
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) calcFatRate() {
	p.fatRate = 0.01
}