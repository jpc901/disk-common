package errors

import "fmt"

type Error struct {
	Code    ErrorCode
	Message string
	Detail  interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[DiskError] code: %d, msg: %s", e.Code, e.Message)
}

func NewInternalServerError(detail interface{}) *Error {
	return &Error{
		Code:    InternalServerErrorCode,
		Message: "internal server error",
		Detail:  detail,
	}
}

func NewUploadFailedError(detail interface{}) *Error {
	return &Error{
		Code:    FileUploadFailedCode,
		Message: "upload failed",
		Detail:  detail,
	}
}

func NewGetFileMetaError(detail interface{}) *Error {
	return &Error{
		Code: GetFileMetaFailedCode,
		Message: "get file meta failed",
		Detail: detail,
	}
}
func NewFileUpdateError(detail interface{}) *Error {
	return &Error{
		Code: FileUpdateFailedCode,
		Message: "file meta update failed",
		Detail: detail,
	}
}

func NewUserSignUpError(detail interface{}) *Error {
	return &Error{
		Code: UserSignUpFailedCode,
		Message: "user sign up failed",
		Detail: detail,
	}
}

func NewUserSignInError(detail interface{}) *Error {
	return &Error{
		Code: UserSignInFailedCode,
		Message: "user sign in failed",
		Detail: detail,
	}
}

func NewUserExistError(detail interface{}) *Error {
	return &Error{
		Code: UserExistCode,
		Message: "用户已存在",
		Detail: detail,
	}
}

func NewUserNotExistError(detail interface{}) *Error {
	return &Error{
		Code: UserNotExistCode,
		Message: "用户不存在",
		Detail: detail,
	}
}

func NewPasswordError(detail interface{}) *Error {
	return &Error{
		Code: PasswordErrorCode,
		Message: "密码错误",
		Detail: detail,
	}
}


func NewGetUserFileError(detail interface{}) *Error {
	return &Error{
		Code: GetUserFileErrorCode,
		Message: "get user file failed",
		Detail: detail,
	}
}


func NewFastUploadError(detail interface{}) *Error {
	return &Error{
		Code: FastUploadFileErrorCode,
		Message: "fast upload file failed",
		Detail: detail,
	}
}

func NewMultipleUploadError(detail interface{}) *Error {
	return &Error{
		Code: MultipleInitFailedCode,
		Message: "Multiple Upload failed",
		Detail: detail,
	}
}

func NewMergeChunkError(detail interface{}) *Error {
	return &Error{
		Code: MergeFailedCode,
		Message: "Merge failed",
		Detail: detail,
	}
}

func NewChunkNotExistError(detail interface{}) *Error {
	return &Error{
		Code: ChunkNotExistCode,
		Message: "Chunk 不存在",
		Detail: detail,
	}
}