package composite

import (
	"fmt"
	"strings"
)

type Company interface {
	add(Company)
	remove(Company)
	display(int)
	lineOfDuty()
}

type RealCompany struct {
	name string
}

type ConcreateCompany struct {
	RealCompany
	list []Company
}

func NewConcreateCompany(name string) *ConcreateCompany {
	return &ConcreateCompany{RealCompany{name}, []Company{}}
}

func (c *ConcreateCompany) add(newc Company) {
	if c == nil {
		return
	}
	c.list = append(c.list, newc)
}

func (c *ConcreateCompany) remove(delc Company) {
	if c == nil {
		return
	}
	for i, val := range c.list {
		if val == delc {
			c.list = append(c.list[:i], c.list[i+1:]...)
			return
		}
	}
	return
}

func (c *ConcreateCompany) display(depth int) {
	if c == nil {
		return
	}
	fmt.Println(strings.Repeat("-", depth), " ", c.name)
	for _, val := range c.list {
		val.display(depth + 2)
	}
}

func (c *ConcreateCompany) lineOfDuty() {
	if c == nil {
		return
	}

	for _, val := range c.list {
		val.lineOfDuty()
	}
}

// HR
type HRDepartment struct {
	RealCompany
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{RealCompany{name}}
}

func (h *HRDepartment) add(c Company)    {}
func (h *HRDepartment) remove(c Company) {}

func (h *HRDepartment) display(depth int) {
	if h == nil {
		return
	}
	fmt.Println(strings.Repeat("*", depth), " ", h.name)
}

func (h *HRDepartment) lineOfDuty() {
	if h == nil {
		return
	}
	fmt.Println(h.name, "員工招聘培訓管理")
}

type FinanceDepartment struct {
	RealCompany
}

func NewFinanceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{RealCompany{name}}
}

func (h *FinanceDepartment) add(c Company)    {}
func (h *FinanceDepartment) remove(c Company) {}

func (h *FinanceDepartment) display(depth int) {
	if h == nil {
		return
	}
	fmt.Println(strings.Repeat("*", depth), " ", h.name)
}

func (h *FinanceDepartment) lineOfDuty() {
	if h == nil {
		return
	}
	fmt.Println(h.name, "公司財務收支管理")
}
