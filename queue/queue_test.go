package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockTask struct {
	mock.Mock
}

func (m *MockTask) Exec() error {
	m.Called()
	return nil
}

func TestAdd(t *testing.T) {
	mockTask := new(MockTask)
	q := &Queue{}
	mockTask.On("Exec").Once()
	q.Add(mockTask, time.Now())
	time.Sleep(100 * time.Millisecond)
	mockTask.AssertExpectations(t)

}
