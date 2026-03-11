package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
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
		Name:          req.Name,
		ThumbnailId:   req.ThumbnailId.ID,
		Goal:          req.Goal,
		Description:   req.Description,
		WebsiteLink:   req.WebsiteLink,
		InstagramLink: req.InstagramLink,
		TwitterLink:   req.TwitterLink,
		LinkedinLink:  req.LinkedinLink,
		YoutubeLink:   req.YoutubeLink,
		DepartmentId:  req.DepartmentId.ID,
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
	if req.Name != nil {
		p.Name = *req.Name
	}
	if req.ThumbnailId.Valid {
		p.ThumbnailId = req.ThumbnailId.ID
		p.Thumbnail = nil
	}
	if req.Goal != nil {
		p.Goal = *req.Goal
	}
	if req.Description != nil {
		p.Description = *req.Description
	}
	if req.WebsiteLink != nil {
		p.WebsiteLink = *req.WebsiteLink
	}
	if req.InstagramLink != nil {
		p.InstagramLink = *req.InstagramLink
	}
	if req.TwitterLink != nil {
		p.TwitterLink = *req.TwitterLink
	}
	if req.LinkedinLink != nil {
		p.LinkedinLink = *req.LinkedinLink
	}
	if req.YoutubeLink != nil {
		p.YoutubeLink = *req.YoutubeLink
	}
	if req.DepartmentId.Valid {
		p.DepartmentId = req.DepartmentId.ID
		p.Department = nil
	}

	progenda, err := s.progendaRepo.Update(ctx, tx, p)
	if err != nil {
		tx.Rollback()
		return entity.Progenda{}, myerror.ParseDBError(err, "progenda")

	}

	//update timelines
	var timelines []entity.Timeline

	for _, tl := range req.Timelines {
		timeline := entity.Timeline{
			Id:         tl.Id,
			ProgendaId: &uid,
		}
		if tl.Date != nil {
			timeline.Date = *tl.Date
		}
		if tl.Info != nil {
			timeline.Info = *tl.Info
		}
		if tl.Link != nil {
			timeline.Link = *tl.Link
		}
		timelines = append(timelines, timeline)
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

	// Skip GetById to avoid expensive preloads. 
	// The database CASCADE will handle related Timelines.
	p := entity.Progenda{Id: uid}

	return s.progendaRepo.Delete(ctx, nil, p)
}
