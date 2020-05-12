package service

import (
	"vortex/model"
)

func FindGroupUserByID (groupID int64) ([]*model.UserGroup, error){
	var userGroup []*model.UserGroup
	err := model.DB.Where("group_id = ?", groupID).Find(&userGroup).Error
	return userGroup, err
}

func CreateGroup(data map[string]interface{}) (*model.Group, error){
	group := &model.Group{
		Name:        data["name"].(string),
		CreatedUid:  data["createdUid"].(int64),
	}
	err := model.DB.Create(group).Error
	return group, err
}

func JoinGroup(data map[string]interface{}) error {
	userGroup := &model.UserGroup{
		UserID:    data["userID"].(int64),
		GroupID:   data["groupID"].(int64),
	}

	err := model.DB.Create(userGroup).Error
	return err
}