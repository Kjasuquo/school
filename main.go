package main

import (
	"fmt"
	"github.com/kjasuquo/school/School"
)

func main() {
	a := School.Applicants{"Alex", [3]int{40, 60, 70}, 17}
	a1 := School.Applicants{"Blex", [3]int{78, 54, 90}, 17}
	a2 := School.Applicants{"Clex", [3]int{23, 56, 49}, 18}
	a3 := School.Applicants{"Elex", [3]int{100, 90, 98}, 19}
	a4 := School.Applicants{"Dlex", [3]int{100, 98, 30}, 16}
	fmt.Println(a.NewStudent(50, 18))
	fmt.Println(a1.NewStudent(50, 18))
	fmt.Println(a2.NewStudent(50, 18))
	fmt.Println(a3.NewStudent(50, 18))
	fmt.Println(a4.NewStudent(50, 18))

	h := School.Principal{}
	h.AdmitStudent("Alex", School.Prospect)
	h1 := School.Principal{}
	h1.AdmitStudent("Blex", School.Prospect)
	h2 := School.Principal{}
	h2.AdmitStudent("Dlex", School.Prospect)
	fmt.Println(School.Students)

	b := School.Student{"Alex", 0, []string{}, ""}
	b1 := School.Student{"Blex", 0, []string{}, ""}
	b2 := School.Student{"Dlex", 0, []string{}, ""}
	b.CourseEnlisted("Alex", []string{"English", "Math", "Java"})
	b1.CourseEnlisted("Blex", []string{"English", "PHE", "Java"})
	b2.CourseEnlisted("Dlex", []string{"Java", "Math", "English"})
	fmt.Println(School.Students)

	c := School.Teacher{"Gbenga", "Java"}
	d := School.Teacher{"Victor", "English"}
	c.Marks("Alex", 70, 50, School.Students)
	c.Marks("Dlex", 60, 50, School.Students)
	c.Marks("Blex", 80, 50, School.Students)
	d.Marks("Dlex", 10, 50, School.Students)
	d.Marks("Alex", 45, 50, School.Students)
	d.Marks("Blex", 57, 50, School.Students)
	fmt.Println(School.Students)

	e := School.Principal{}
	e.ExpelStudent(30, School.Students)
	fmt.Println(School.Students)

	f := School.Staff{"Alex", 48000}
	f.SchoolFees(48000)
	g := School.Staff{"Blex", 40000}
	g.SchoolFees(48000)

	fmt.Println(School.Students)

}
