package relation

import (
	"fmt"

	"github.com/BaiYiZi/learn-gorm/database"
	"github.com/BaiYiZi/learn-gorm/model"
)

func CreateStudentAndCourse(student *model.Student, course *model.Course) (*model.Student, *model.Course, error) {
	db := database.GetDB()

	if err := db.Model(&model.Student{}).Omit("PeopleID").Create(student).Error; err != nil {
		return nil, nil, fmt.Errorf("%s", "create student failed")
	}

	if err := db.Model(&model.Course{}).Create(course).Error; err != nil {
		return nil, nil, fmt.Errorf("%s", "create course failed")
	}

	return student, course, nil
}

func ConfirmStudentAndCourseRelation(studentID uint, courseID uint) error {
	db := database.GetDB()

	if err := db.Model(&model.SelectCourse{}).Create(&model.SelectCourse{
		StudentID: studentID,
		CourseID:  courseID,
	}).Error; err != nil {
		return fmt.Errorf("confirm relation failed")
	}

	return nil
}

func CancelStudentAndCourseRelation(studentID uint, courseID uint) error {
	db := database.GetDB()

	if err := db.Delete(&model.SelectCourse{}, "student_id = ? and course_id = ? limit 1", studentID, courseID).Error; err != nil {
		return fmt.Errorf("confirm relation failed")
	}

	return nil
}

func SelectCourseInfo() (any, error) {
	db := database.GetDB()

	queryStruct := []struct {
		Gender         bool   `json:"gender" gorm:"column:student_gender"`
		StudentID      uint   `json:"student_id" gorm:"column:student_id"`
		StudentName    string `json:"student_name" gorm:"column:student_name"`
		CourseID       uint   `json:"course_id" gorm:"column:course_id"`
		CourseFullMark int    `json:"course_full_mark" gorm:"column:course_full_mark"`
		CourseName     string `json:"course_name" gorm:"column:course_name"`
	}{}

	result := db.
		Table(model.SelectCourse{}.TableName()+" sc").
		Select(
			"student.gender as student_gender",
			"student.id as student_id",
			"student.name as student_name",
			"course.id as course_id",
			"course.full_mark as course_full_mark",
			"course.name as course_name",
		).
		InnerJoins(
			"JOIN table_student student ON student.id = sc.student_id",
		).
		InnerJoins(
			"JOIN table_course course ON course.id = sc.course_id",
		).Find(&queryStruct)
	if result.Error != nil {
		return nil, fmt.Errorf("query failed")
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &queryStruct, nil
}
