package util

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"per_page"`
	TotalRows  int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ErrorResponse struct {
	Error   any    `json:"errors"`
	Message string `json:"message"`
}

type ImportResponse struct {
	Errors     []ImportError `json:"errors"`
	TotalInput int           `json:"total_input"`
	Success    int           `json:"success"`
	Failed     int           `json:"failed"`
}

type ImportError struct {
	Row    int `json:"nomor"`
	Errors any `json:"error"`
}

const (
	MESSAGE_OK_CREATE              = "Berhasil memasukkan data."
	MESSAGE_OK_DELETE              = "Berhasil menghapus data."
	MESSAGE_OK_UPDATE              = "Berhasil memperbaharui data."
	MESSAGE_OK_READ                = "Berhasil menampilkan data."
	MESSAGE_OK_IMPORT              = "Berhasil mengimport data."
	MESSAGE_OK_EXPORT              = "Berhasil mengexport data."
	MESSAGE_OK_UPLOAD              = "Berhasil mengunggah data"
	MESSAGE_OK_LOGIN               = "Berhasil masuk."
	MESSAGE_OK_REFRESH             = "Berhasil membuat ulang token."
	MESSAGE_OK_LOGOUT              = "Berhasil keluar."
	MESSAGE_OK_REQUEST_TOKEN_RESET = "Berhasil meminta token."
	MESSAGE_OK_CHANGE_PASSWORD     = "Berhasil memperbaharui password."
	MESSAGE_FAILED_CREATE          = "Gagal memasukkan data data."
	MESSAGE_FAILED_DELETE          = "Gagal menghapus data."
	MESSAGE_FAILED_UPDATE          = "Gagal memperbaharui data."
	MESSAGE_FAILED_READ            = "Gagal menampilkan data."
	MESSAGE_FAILED_IMPORT          = "Gagal mengimport data."
	MESSAGE_FAILED_EXPORT          = "Gagal mengexport data."
	MESSAGE_FAILED_VALIDATION      = "Gagal memvalidasi data."
	MESSAGE_BAD_REQUEST            = "Permintaan bermasalah."
	MESSAGE_BAD_SYSTEM             = "Server bermasalah."
	MESSAGE_UNAUTHORIZED           = "Unauthorized."
	MESSAGE_CONFLICT               = "Data sudah ada."
	MESSAGE_NOT_FOUND              = "Data tidak ada."
	MESSAGE_FAILED_LOGIN           = "Username atau password tidak cocok."
)

type repsonseInterface interface {
	Success(data any, meta any, message string) error
	Error(errors any, message string, statusCode int) error
	Import(errors []ImportError, totalInput int, failed int) error
}

type responseStruct struct {
	fiberCtx *fiber.Ctx
}

func NewResponse(fiberCtx *fiber.Ctx) repsonseInterface {
	return &responseStruct{
		fiberCtx: fiberCtx,
	}
}

func (r *responseStruct) Success(data any, meta any, message string) error {
	return r.fiberCtx.Status(fiber.StatusOK).JSON(SuccessResponse{
		Data:    data,
		Meta:    meta,
		Message: message,
	})
}

func (r *responseStruct) Error(errors any, message string, statusCode int) error {
	return r.fiberCtx.Status(statusCode).JSON(ErrorResponse{
		Error:   errors,
		Message: message,
	})
}

func (r *responseStruct) Import(errors []ImportError, totalInput int, failed int) error {
	return r.fiberCtx.Status(fiber.StatusOK).JSON(ImportResponse{
		Errors:     errors,
		TotalInput: totalInput,
		Success:    totalInput - failed,
		Failed:     failed,
	})
}
