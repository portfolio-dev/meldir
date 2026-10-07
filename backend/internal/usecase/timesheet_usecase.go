package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/repository"
)

type TimesheetUsecase interface {
	ListTimesheets(ctx context.Context, currentUser *domain.User, limit int) ([]domain.TimesheetLog, error)
	LogWork(ctx context.Context, currentUser *domain.User, req domain.CreateTimesheetRequest) (*domain.TimesheetLog, error)
	GetSummary(ctx context.Context, currentUser *domain.User) (*domain.TimesheetSummary, error)
}

type timesheetUsecase struct {
	timesheetRepo repository.TimesheetRepository
}

func NewTimesheetUsecase(timesheetRepo repository.TimesheetRepository) TimesheetUsecase {
	return &timesheetUsecase{timesheetRepo: timesheetRepo}
}

func (u *timesheetUsecase) ListTimesheets(ctx context.Context, currentUser *domain.User, limit int) ([]domain.TimesheetLog, error) {
	if u.timesheetRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	// Engineer hanya melihat timesheet miliknya sendiri; Direktur/Admin melihat seluruh log
	var engID int64
	if currentUser.Role == domain.RoleEngineer {
		engID = currentUser.ID
	}

	return u.timesheetRepo.ListTimesheets(ctx, engID, limit)
}

func (u *timesheetUsecase) LogWork(ctx context.Context, currentUser *domain.User, req domain.CreateTimesheetRequest) (*domain.TimesheetLog, error) {
	if u.timesheetRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	// Hanya Engineer dan Direktur yang boleh mencatat timesheet
	if currentUser.Role != domain.RoleEngineer && currentUser.Role != domain.RoleDirektur {
		return nil, errors.New("hanya akun engineer atau direktur yang berwenang mencatat waktu kerja")
	}

	if req.HoursSpent <= 0 {
		return nil, errors.New("durasi jam kerja harus lebih dari 0")
	}

	if strings.TrimSpace(req.WorkDescription) == "" {
		return nil, errors.New("rincian deskripsi pengerjaan wajib diisi")
	}

	if strings.TrimSpace(req.LogDate) == "" {
		req.LogDate = time.Now().Format("2006-01-02")
	}

	logEntry := &domain.TimesheetLog{
		EngineerID:      currentUser.ID,
		TicketCode:      strings.TrimSpace(req.TicketCode),
		ProjectName:     strings.TrimSpace(req.ProjectName),
		HoursSpent:      req.HoursSpent,
		WorkDescription: strings.TrimSpace(req.WorkDescription),
		LogDate:         req.LogDate,
	}

	return u.timesheetRepo.CreateTimesheet(ctx, logEntry)
}

func (u *timesheetUsecase) GetSummary(ctx context.Context, currentUser *domain.User) (*domain.TimesheetSummary, error) {
	if u.timesheetRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	var engID int64
	if currentUser.Role == domain.RoleEngineer {
		engID = currentUser.ID
	}

	return u.timesheetRepo.GetSummary(ctx, engID)
}
