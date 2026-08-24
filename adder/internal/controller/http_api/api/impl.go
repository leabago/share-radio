package api

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/docs/gen"
)

func errUnhandled(err error) gen.V1Error {
	return gen.V1Error{
		Error: new(err.Error()),
	}
}

type UserService interface {
	LoginV2(ctx context.Context, request gen.LoginRequestObject) (gen.V1Token, error)
}

var _ gen.StrictServerInterface = (*Controller)(nil)

type Controller struct {
	userService UserService
}

func NewController(userService UserService) Controller {
	return Controller{
		userService: userService,
	}
}

func (c Controller) Login(ctx context.Context, request gen.LoginRequestObject) (gen.LoginResponseObject, error) {

	resp, err := c.userService.LoginV2(ctx, request)

	if err != nil {
		return gen.Login500JSONResponse(errUnhandled(err)), nil
	}

	return gen.Login200JSONResponse(resp), nil

}

func (c Controller) Register(ctx context.Context, request gen.RegisterRequestObject) (gen.RegisterResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) ListTasks(ctx context.Context, request gen.ListTasksRequestObject) (gen.ListTasksResponseObject, error) {

	fmt.Println("ListTasks middleware")

	return gen.ListTasks200JSONResponse{}, nil
}

func (c Controller) CreateTask(ctx context.Context, request gen.CreateTaskRequestObject) (gen.CreateTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) DeleteTask(ctx context.Context, request gen.DeleteTaskRequestObject) (gen.DeleteTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) GetTask(ctx context.Context, request gen.GetTaskRequestObject) (gen.GetTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) UpdateTask(ctx context.Context, request gen.UpdateTaskRequestObject) (gen.UpdateTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) TransitionTask(ctx context.Context, request gen.TransitionTaskRequestObject) (gen.TransitionTaskResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) DoTranslate(ctx context.Context, request gen.DoTranslateRequestObject) (gen.DoTranslateResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) History(ctx context.Context, request gen.HistoryRequestObject) (gen.HistoryResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) Profile(ctx context.Context, request gen.ProfileRequestObject) (gen.ProfileResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

//
//func NewBffController(bffService BffService) BffController {
//	return BffController{
//		bffService: bffService,
//	}
//}
//
//func (b BffController) PostApiHolidays(
//	ctx context.Context, request gen.PostApiHolidaysRequestObject) (gen.PostApiHolidaysResponseObject, error) {
//
//	resp, err := b.bffService.GetHolidays(ctx)
//	if err != nil {
//		return gen.PostApiHolidays500JSONResponse(errUnhandled(err)), nil
//	}
//
//	return gen.PostApiHolidays200JSONResponse(resp), nil
//}
//
//func (b BffController) GetApiMainHealthCheck(
//	ctx context.Context, request gen.GetApiMainHealthCheckRequestObject) (gen.GetApiMainHealthCheckResponseObject, error) {
//	resp := gen.HealthCheckResponse{
//		Status: "OK",
//	}
//
//	return gen.GetApiMainHealthCheck200JSONResponse(resp), nil
//}
