package handler

import "restfulpi/internal/models"

type SessionManagerService interface{

	StartSession(id string,filePath string) (*models.Session,error)
} 
