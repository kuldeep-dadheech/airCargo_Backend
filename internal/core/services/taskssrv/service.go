package taskssrv

import (
	"aircargo/internal/core/domain/repositories/rdbms"
	"aircargo/internal/core/domain/services"
	"aircargo/internal/core/ports"
)

type Service struct {
	tasksRepository ports.RdbmsTasksRepository
}

func New(
	tasksRepository ports.RdbmsTasksRepository,
) *Service {
	return &Service{
		tasksRepository: tasksRepository,
	}
}

func (s *Service) CreateTask(
	bq_id int,
	count int,
) error {
	_,err := s.tasksRepository.CreateOne(
		bq_id,
		count,
	)
	if err != nil {
		return err
	}
	// fmt.Println(id);
	return nil
}

func (s *Service) SelectTask(
	id int,
) (services.Task, error) {
	task, err := s.tasksRepository.SelectOne(id)
	if err != nil {
		return services.Task{}, err
	}
	taskService := s.mapRepoDomainToService(task)
	return taskService, nil
}

func (s *Service) mapRepoDomainToService(
	task rdbms.Task,
) services.Task {
	return services.Task{
		Id: task.Id,
		Bq_id: task.Bq_id,
		Count: task.Count,
	}
}

func (s *Service) UpdateTask(
	id int, 
	bq_id int,
	count int,
) error {
	err := s.tasksRepository.UpdateOne(
		id,
		bq_id,
		count,
	)
	if err != nil {
		return err
	}
	return nil

}