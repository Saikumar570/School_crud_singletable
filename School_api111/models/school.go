// this below code gives stores the data in single table called school and stores 3 classes data class6,class7, class8 and gives it to the output , the class details we need


package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	Name          string        `json:"name"`
	SchoolId      string        `json:"school_id"`
	SchoolAddress SchoolAddress `gorm:"embedded" json:"school_address"`
	Class6        ClassDetail   `gorm:"embedded;embedded_prefix:class6_" json:"class_6"`
	Class7        ClassDetail   `gorm:"embedded;embedded_prefix:class7_" json:"class_7"`
	Class8        ClassDetail   `gorm:"embedded;embedded_prefix:class8_" json:"class_8"`
}

type SchoolAddress struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type ClassDetail struct { // Classs
	//ClasssName
	//Count
	//Id
	//Students []int64
	Students     []Student `gorm:"-" json:"students"`
	StudentsJSON string    `gorm:"column:students_json" json:"-"`

	// Students     []Student `gorm:"-" json:"students"`
	// StudentsJSON string    `gorm:"column:students_json" json:"-"`
}

type Student struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	AddressStruct *Address `json:"address" gorm:"-"`
	AddressDb     string   `json:"-" gorm:"student_address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func (c *School) BeforeSave(tx *gorm.DB) (err error) {
	if len(c.Class6.Students) > 0 {
		data, err := json.Marshal(c.Class6.Students)
		if err != nil {
			return err
		}
		c.Class6.StudentsJSON = string(data)
		for _, student := range c.Class6.Students {
			if student.AddressStruct != nil {
				data, err := json.Marshal(student.AddressStruct)
				if err != nil {
					return err
				}
				student.AddressDb = string(data)
			}
		}
	}

	if len(c.Class7.Students) > 0 {
		data, err := json.Marshal(c.Class7.Students)
		if err != nil {
			return err
		}
		c.Class7.StudentsJSON = string(data)
		for _, student := range c.Class7.Students {
			if student.AddressStruct != nil {
				data, err := json.Marshal(student.AddressStruct)
				if err != nil {
					return err
				}
				student.AddressDb = string(data)
			}
		}
	}
	if len(c.Class8.Students) > 0 {
		data, err := json.Marshal(c.Class8.Students)
		if err != nil {
			return err
		}
		c.Class8.StudentsJSON = string(data)
		for _, student := range c.Class8.Students {
			if student.AddressStruct != nil {
				data, err := json.Marshal(student.AddressStruct)
				if err != nil {
					return err
				}
				student.AddressDb = string(data)
			}
		}
	}

	// iterate throught the c.Students and then for each student, again do the above operation of Address too.

	return nil
}

func (c *School) AfterFind(tx *gorm.DB) (err error) {
	if c.Class6.StudentsJSON != "" {
		if err := json.Unmarshal([]byte(c.Class6.StudentsJSON), &c.Class6.Students); err != nil {
			return err
		}
	}

	if c.Class7.StudentsJSON != "" {
		if err := json.Unmarshal([]byte(c.Class7.StudentsJSON), &c.Class7.Students); err != nil {
			return err
		}
	}

	if c.Class8.StudentsJSON != "" {
		if err := json.Unmarshal([]byte(c.Class8.StudentsJSON), &c.Class8.Students); err != nil {
			return err
		}
	}

	return nil
}


//The below code stores data in single table school but it stores all the class details that the user gives ...

package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	Name          string        `json:"name"`
	SchoolId      string        `json:"school_id"`
	SchoolAddress SchoolAddress `gorm:"embedded" json:"school_address"`
	Classes       []Class       `gorm:"-" json:"classes"`
	ClassesJSON   string        `gorm:"column:classes_json" json:"-"`
}

type SchoolAddress struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Class struct {
	ClassName    string    `json:"class_name"`
	Students     []Student `gorm:"-" json:"students"`
	StudentsJSON string    `gorm:"column:students_json" json:"-"`
}

type Student struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	AddressStruct *Address `json:"address" gorm:"-"`
	AddressDb     string   `json:"-" gorm:"student_address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func (s *School) BeforeSave(tx *gorm.DB) (err error) {
	if len(s.Classes) > 0 {
		for i, class := range s.Classes {
			if len(class.Students) > 0 {
				data, err := json.Marshal(class.Students)
				if err != nil {
					return err
				}
				s.Classes[i].StudentsJSON = string(data)
				for j, student := range class.Students {
					if student.AddressStruct != nil {
						data, err := json.Marshal(student.AddressStruct)
						if err != nil {
							return err
						}
						s.Classes[i].Students[j].AddressDb = string(data)
					}
				}
			}
		}
		data, err := json.Marshal(s.Classes)
		if err != nil {
			return err
		}
		s.ClassesJSON = string(data)
	}
	return nil
}

func (s *School) AfterFind(tx *gorm.DB) (err error) {
	if s.ClassesJSON != "" {
		if err := json.Unmarshal([]byte(s.ClassesJSON), &s.Classes); err != nil {
			return err
		}
		for i, class := range s.Classes {
			if class.StudentsJSON != "" {
				if err := json.Unmarshal([]byte(class.StudentsJSON), &s.Classes[i].Students); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

