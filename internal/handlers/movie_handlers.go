package handlers

import (
	"encoding/json"
	"filmlib/internal/apperrors"
	"filmlib/internal/pkg/dto"
	"filmlib/internal/service"
	"filmlib/internal/utils"
	"net/http"
)

type MovieHandler struct {
	ms service.IMovieService
}

// @Summary Создать фильм
// @Description
// @Tags movies
//
// @Accept  json
// @Produce  json
//
// @Param movieData body dto.NewMovie true "Данные о новом фильме"
//
// @Success 200  {object}  entities.Movie "Объект нового фильма"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /movies/ [post]
func (mh MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	funcName := "CreateMovie"

	rCtx := r.Context()
	logger, requestID, err := utils.GetLoggerAndID(rCtx)
	if err != nil {
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	var newMovie dto.NewMovie
	err = json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		logger.DebugFmt("Failed to decode request: "+err.Error(), requestID, funcName, nodeName)
		apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		return
	}

	// err = validate.Struct(newMovie)
	// if err != nil {
	// 	logger.Error("Validation failed")
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}

	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		logger.DebugFmt(err.Error(), requestID, funcName, nodeName)
	// 	}
	// 	apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
	// 	return
	// }

	actor, err := mh.ms.Create(rCtx, newMovie)
	if err != nil {
		logger.Error(err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Movie created", requestID, funcName, nodeName)

	jsonResponse, err := json.Marshal(actor)
	if err != nil {
		logger.Error("Failed to marshal response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Failed to return response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	r.Body.Close()
}

// @Summary Получить данные об фильме
// @Description Получить данные об фильме по его ID
// @Tags movies
//
// @Produce  json
//
// @Param movieID path uint true "ID фильма"
//
// @Success 200  {object}  entities.Movie "Объект фильма"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /movies/{id}/ [get]
func (mh MovieHandler) ReadMovie(w http.ResponseWriter, r *http.Request) {
}

// @Summary Изменить данные об фильме
// @Description Изменить данные об фильме по его ID
// @Tags movies
//
// @Accept  json
// @Produce  json
//
// @Param movieID path uint true "ID фильма"
// @Param movieData body dto.UpdatedMovie true "Обновлённые данные фильма"
//
// @Success 204  {string}  "no response"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /movies/{id}/ [patch]
func (mh MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
}

// @Summary Удалить данные об фильме
// @Description Удалить данные об фильме по его ID
// @Tags movies
//
// @Produce  json
//
// @Param movieID path uint true "ID фильма"
//
// @Success 204  {string}  "no response"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /movies/{id}/ [delete]
func (mh MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
}

// @Summary Получить список фильмов
// @Description Получить список всех фильмов
// @Description Если порядок сортировки не указан, для каждого типа есть порядок по умолчанию:
// @Description Для названия и даты возрастающий, для рейтинга - убывающий
// @Tags movies
//
// @Produce  json
//
// @Param sortType query uint true "Тип сортировки (0 - название, 1 - рейтинг, 2 - дата выпуска)"
// @Param sortOrder query uint false "Порядок сортировки (0 - возрастающий, 1 - убывающий)"
//
// @Success 200  {object}  []entities.Movie "Список фильмов"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /movies/ [get]
func (mh MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
}
