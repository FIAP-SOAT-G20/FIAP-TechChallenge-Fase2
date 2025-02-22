package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type StaffController struct {
	useCase port.StaffUseCase
}

func NewStaffController(useCase port.StaffUseCase) *StaffController {
	return &StaffController{useCase}
}

func (c *StaffController) List(ctx context.Context, p port.Presenter, i dto.ListStaffsInput) error {
	staffs, total, err := c.useCase.List(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{
		Total:  total,
		Page:   i.Page,
		Limit:  i.Limit,
		Result: staffs,
	})

	return nil
}

func (c *StaffController) Create(ctx context.Context, p port.Presenter, i dto.CreateStaffInput) error {
	staff, err := c.useCase.Create(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: staff})

	return nil
}

func (c *StaffController) Get(ctx context.Context, p port.Presenter, i dto.GetStaffInput) error {
	staff, err := c.useCase.Get(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: staff})

	return nil
}

func (c *StaffController) Update(ctx context.Context, p port.Presenter, i dto.UpdateStaffInput) error {
	staff, err := c.useCase.Update(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: staff})

	return nil
}

func (c *StaffController) Delete(ctx context.Context, p port.Presenter, i dto.DeleteStaffInput) error {
	staff, err := c.useCase.Delete(ctx, i)

	if err != nil {
		return err
	}
	p.Present(dto.PresenterInput{Result: staff})

	return nil
}
