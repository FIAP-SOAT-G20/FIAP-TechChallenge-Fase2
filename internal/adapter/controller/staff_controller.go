package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type StaffController struct {
	listStaffsUseCase  port.ListStaffsUseCase
	createStaffUseCase port.CreateStaffUseCase
	getStaffUseCase    port.GetStaffUseCase
	updateStaffUseCase port.UpdateStaffUseCase
	deleteStaffUseCase port.DeleteStaffUseCase
}

func NewStaffController(
	listUC port.ListStaffsUseCase,
	createUC port.CreateStaffUseCase,
	getUC port.GetStaffUseCase,
	updateUC port.UpdateStaffUseCase,
	deleteUC port.DeleteStaffUseCase,
) *StaffController {
	return &StaffController{
		listStaffsUseCase:  listUC,
		createStaffUseCase: createUC,
		getStaffUseCase:    getUC,
		updateStaffUseCase: updateUC,
		deleteStaffUseCase: deleteUC,
	}
}

func (c *StaffController) ListStaffs(ctx context.Context, input dto.ListStaffsInput) error {
	err := c.listStaffsUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *StaffController) CreateStaff(ctx context.Context, input dto.CreateStaffInput) error {
	err := c.createStaffUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *StaffController) GetStaff(ctx context.Context, input dto.GetStaffInput) error {
	err := c.getStaffUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *StaffController) UpdateStaff(ctx context.Context, input dto.UpdateStaffInput) error {
	err := c.updateStaffUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *StaffController) DeleteStaff(ctx context.Context, input dto.DeleteStaffInput) error {
	return c.deleteStaffUseCase.Execute(ctx, input)
}
