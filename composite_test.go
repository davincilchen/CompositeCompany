package composite

import (
	"fmt"
	"testing"
)

func TestComponent(t *testing.T) {
	fmt.Println("")
	root := NewConcreateCompany("台北總公司")
	root.add(NewHRDepartment("總公司人力資源部"))
	root.add(NewFinanceDepartment("總公司財務部"))
	// =============================================================== //
	HR := NewHRDepartment("總公司人力資源部2")
	HRC := NewHRDepartment("總公司人力資源部2小組-----------")
	HR.add(HRC) // 無效,沒實作
	root.add(HR)
	// =============================================================== //

	compA := NewConcreateCompany("新竹分公司")
	compA.add(NewHRDepartment("新竹分公司人力資源部"))
	compA.add(NewFinanceDepartment("新竹分公司財務部"))
	root.add(compA)

	compB := NewConcreateCompany("關西分部")
	compB.add(NewHRDepartment("關西分部總務處"))
	compB.add(NewFinanceDepartment("關西分部財務部"))
	compA.add(compB)

	compC := NewConcreateCompany("竹東分部")
	compC.add(NewHRDepartment("竹東分部總務處"))
	compC.add(NewFinanceDepartment("竹東分部財務部"))
	compA.add(compC)

	fmt.Println("公司組織")
	root.display(3)

	fmt.Println("=====================================")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("職責")
	root.lineOfDuty()
	fmt.Println("")
}
