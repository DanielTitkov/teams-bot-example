package app

// func (a *App) CreateProject(t *domain.Task) error {
// 	u, err := a.repo.GetUserByUsername(t.User)
// 	if err != nil {
// 		return err
// 	}

// 	tt, err := a.repo.GetTaskTypeByCode(t.Type)
// 	if err != nil {
// 		return err
// 	}

// 	err = a.ValidateTaskParams(t)
// 	if err != nil {
// 		return err
// 	}

// 	ts := strconv.FormatInt(time.Now().Unix(), 10)
// 	code := strings.Join([]string{t.User, t.Type, t.Slug, ts}, "_")
// 	t.Code = code
// 	_, err = a.repo.CreateTask(t, u, tt)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }