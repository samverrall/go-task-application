package task

/*
func TestCreateServive(t *testing.T) {
	ctx := context.Background()
	l := logger.New("debug")
	taskService := NewService(mockTaskRepo{}, l)

	execute := func(name task.Name, completeBy task.CompleteBy) error {
		_, err := taskService.CreateTask(ctx, CreateTaskDTO{
			Name:       name.String(),
			CompleteBy: completeBy.Time(),
		})
		return err
	}

	t.Run("valid inputs", func(t *testing.T) {
		f := func(name task.Name, completeBy task.CompleteBy) bool {
			err := execute(name, completeBy)
			return err == nil
		}
		if err := quick.Check(f, nil); err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
*/
