package model

type Student struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	// true, false := male, female
	Gender bool `gorm:"column:gender" json:"gender"`
}

func (Student) TableName() string {
	return "table_student"
}

type Course struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	FullMark int    `gorm:"column:full_mark" json:"full_mark"`
}

func (Course) TableName() string {
	return "table_course"
}

type SelectCourse struct {
	ID        uint `gorm:"primarykey" json:"id"`
	StudentID uint `gorm:"column:student_id" json:"student_id"`
	CourseID  uint `gorm:"column:course_id" json:"course_id"`
}

func (SelectCourse) TableName() string {
	return "table_select_course"
}
