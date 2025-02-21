package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type StaffController struct {
	useCase   port.StaffUseCase
	Presenter port.Presenter
}

func NewStaffController(useCase port.StaffUseCase) *StaffController {
	return &StaffController{useCase, nil}
}

func (c *StaffController) List(ctx context.Context, input dto.ListStaffsInput) error {
	staffs, total, err := c.useCase.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: staffs,
	})

	return nil
}

func (c *StaffController) Create(ctx context.Context, input dto.CreateStaffInput) error {
	staff, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: staff,
	})

	return nil
}

func (c *StaffController) Get(ctx context.Context, input dto.GetStaffInput) error {
	staff, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: staff,
	})

	return nil
}

func (c *StaffController) Update(ctx context.Context, input dto.UpdateStaffInput) error {
	staff, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: staff,
	})

	return nil
}

func (c *StaffController) Delete(ctx context.Context, input dto.DeleteStaffInput) error {
	staff, err := c.useCase.Delete(ctx, input)

	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: staff,
	})

	return nil
}
