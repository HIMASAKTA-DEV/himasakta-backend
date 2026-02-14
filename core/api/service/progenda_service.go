package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgendaService interface {
	Create(ctx context.Context, req dto.CreateProgendaRequest) (entity.Progenda, error)
	GetAll(ctx context.Context, metaReq meta.Meta, search string, departmentId string, name string) ([]entity.Progenda, meta.Meta, error)
	GetById(ctx context.Context, id string) (entity.Progenda, error)
	Update(ctx context.Context, id string, req dto.UpdateProgendaRequest) (entity.Progenda, error)
	Delete(ctx context.Context, id string) error
}

type progendaService struct {
	db           *gorm.DB
	progendaRepo repository.ProgendaRepository
	timelineRepo repository.TimelineRepository
}

func NewProgenda(db *gorm.DB, progendaRepo repository.ProgendaRepository, timelineRepo repository.TimelineRepository) ProgendaService {
	return &progendaService{
		db:           db,
		progendaRepo: progendaRepo,
		timelineRepo: timelineRepo,
	}
}

func (s *progendaService) Create(ctx context.Context, req dto.CreateProgendaRequest) (entity.Progenda, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return entity.Progenda{}, tx.Error
	}

	//create progenda
	progenda, err := s.progendaRepo.Create(ctx, tx, entity.Progenda{
		Name:         req.Name,
		ThumbnailId:  req.ThumbnailId,
		Goal:         req.Goal,
		Description:  req.Description,
		WebsiteLink:  req.WebsiteLink,
		DepartmentId: req.DepartmentId,
	})
	if err != nil {
		tx.Rollback()
		return entity.Progenda{}, err
	}

	//create timelines
	var timelines []entity.Timeline

	for _, tl := range req.Timelines {
		timelines = append(timelines, entity.Timeline{
			ProgendaId: &progenda.Id,
			Date:       tl.Date,
			Info:       tl.Info,
			Link:       tl.Link,
		})
	}

	if err := s.timelineRepo.BulkCreate(ctx, tx, timelines); err != nil {
		tx.Rollback()
		return entity.Progenda{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entity.Progenda{}, err
	}

	progenda.Timelines = timelines
	return progenda, nil

}

func (s *progendaService) GetAll(ctx context.Context, metaReq meta.Meta, search string, departmentId string, name string) ([]entity.Progenda, meta.Meta, error) {
	var deptId *uuid.UUID
	if departmentId != "" {
		id, err := uuid.Parse(departmentId)
		if err == nil {
			deptId = &id
		}
	}
	return s.progendaRepo.GetAll(ctx, nil, metaReq, search, deptId, name)
}

func (s *progendaService) GetById(ctx context.Context, id string) (entity.Progenda, error) {
	uid, _ := uuid.Parse(id)
	return s.progendaRepo.GetById(ctx, nil, uid)
}

func (s *progendaService) Update(ctx context.Context, id string, req dto.UpdateProgendaRequest) (entity.Progenda, error) {
	uid, _ := uuid.Parse(id)

	tx := s.db.Begin()
	if tx.Error != nil {
		return entity.Progenda{}, tx.Error
	}

	p, err := s.progendaRepo.GetById(ctx, tx, uid)
	if err != nil {
		return p, err
	}

	//update Progenda
	p.Name = req.Name
	p.ThumbnailId = req.ThumbnailId
	p.Goal = req.Goal
	p.Description = req.Description
	p.WebsiteLink = req.WebsiteLink
	p.InstagramLink = req.InstagramLink
	p.TwitterLink = req.TwitterLink
	p.LinkedinLink = req.LinkedinLink
	p.YoutubeLink = req.YoutubeLink
	p.DepartmentId = req.DepartmentId

	progenda, err := s.progendaRepo.Update(ctx, tx, p)
	if err != nil {
		tx.Rollback()
		return entity.Progenda{}, err

	}

	//update timelines
	var timelines []entity.Timeline

	for _, tl := range req.Timelines {
		timelines = append(timelines, entity.Timeline{
			Id:         tl.Id,
			ProgendaId: &uid,
			Date:       tl.Date,
			Info:       tl.Info,
			Link:       tl.Link,
		})
	}

	_, err = s.timelineRepo.BulkUpdate(ctx, tx, timelines)
	if err != nil {
		tx.Rollback()
		return entity.Progenda{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entity.Progenda{}, err
	}

	progenda.Timelines = timelines
	return progenda, nil
}

func (s *progendaService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	p, err := s.progendaRepo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}

	//delete timelines
	var timelines []entity.Timeline

	for _, tl := range p.Timelines {
		timelines = append(timelines, entity.Timeline{
			Id:         tl.Id,
			ProgendaId: &p.Id,
			Date:       tl.Date,
			Info:       tl.Info,
		})
	}
	if err := s.timelineRepo.BulkDelete(ctx, tx, timelines); err != nil {
		tx.Rollback()
		return err
	}

	//delete progenda
	if err := s.progendaRepo.Delete(ctx, nil, p); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
