package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type StaffController struct {
	useCase   port.StaffUseCase
	Presenter port.StaffPresenter
}

func NewStaffController(useCase port.StaffUseCase) *StaffController {
	return &StaffController{useCase, nil}
}

func (c *StaffController) ListStaffs(ctx context.Context, input dto.ListStaffsInput) error {
	staffs, total, err := c.useCase.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.StaffPresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: staffs,
	})

	return nil
}

func (c *StaffController) CreateStaff(ctx context.Context, input dto.CreateStaffInput) error {
	staff, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.StaffPresenterInput{
		Result: staff,
	})

	return nil
}

func (c *StaffController) GetStaff(ctx context.Context, input dto.GetStaffInput) error {
	staff, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.StaffPresenterInput{
		Result: staff,
	})

	return nil
}

func (c *StaffController) UpdateStaff(ctx context.Context, input dto.UpdateStaffInput) error {
	staff, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.StaffPresenterInput{
		Result: staff,
	})

	return nil
}

func (c *StaffController) DeleteStaff(ctx context.Context, input dto.DeleteStaffInput) error {
	staff, err := c.useCase.Delete(ctx, input)

	if err != nil {
		return err
	}

	c.Presenter.Present(dto.StaffPresenterInput{
		Result: staff,
	})

	return nil
}
