package model

import "errors"

var (
	ERROR_USER_NOT_FOUND      = errors.New("error!,user not found")
	ERROR_USER_IS_EXITED      = errors.New("error!,user had been exited")
	ERROR_PASSWORD_WRONG      = errors.New("error!,password wrong")
	ERROR_PHONE_IS_EMPTY      = errors.New("error!,phone number is empty")
	ERROR_PASSWORD_IS_EMPTY   = errors.New("error!,password us  empty")
	ERROR_REGISTER_FAIL       = errors.New("error!,register fail")
	ERROR_LOGIN_FAIL          = errors.New("error!,login fail")
	ERROR_UPDATE_FAIL         = errors.New("error!,update error")
	ERROR_DELETE_FAIL         = errors.New("error!,delete error")
	ERROR_TOKEN_GENERATE_FAIL = errors.New("error!,token generate error")
	ERROR_GET_INFO_FAIL       = errors.New("error!,get info error")

	CURRICULUM_HAD_EXISTED = errors.New("curriculum had existed ")
	CURRICULUM_NOT_FOUND   = errors.New("curriculum not found ")
)
