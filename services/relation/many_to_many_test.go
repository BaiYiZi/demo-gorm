package relation

import (
	"testing"

	"github.com/BaiYiZi/learn-gorm/model"
)

func TestCreateStudentAndCourse(t *testing.T) {
	student := &model.Student{
		Name:   "student1",
		Gender: true,
	}

	course := &model.Course{
		Name:     "course1",
		FullMark: 150,
	}

	_, _, err := CreateStudentAndCourse(student, course)
	if err != nil {
		t.Error(err)
	}

	t.Logf("student id is %d, course id is %d", student.ID, course.ID)
}

func TestConfirmStudentAndCourseRelation(t *testing.T) {
	var studentID uint = 1
	var courseID uint = 1

	err := ConfirmStudentAndCourseRelation(studentID, courseID)
	if err != nil {
		t.Error(err)
	}

	t.Logf("confirm relation success student id %d and course id %d", studentID, courseID)
}

func TestCancelStudentAndCourseRelation(t *testing.T) {
	var studentID uint = 1
	var courseID uint = 1

	err := CancelStudentAndCourseRelation(studentID, courseID)
	if err != nil {
		t.Error(err)
	}

	t.Logf("cancel relation success student id %d and course id %d", studentID, courseID)
}

func TestSelectCourseInfo(t *testing.T) {
	res, err := SelectCourseInfo()
	if err != nil {
		t.Error(t)
	}

	t.Log(res)
}
