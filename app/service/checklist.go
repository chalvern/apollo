package service

import (
	"fmt"
	"strings"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/jinzhu/gorm"
)

var (
	checklistModel = model.Checklist{}
)

// ChecklistsWithOrderQuery 获取某个分享的 checklist
// 按照双链表的先后次序返回
func ChecklistsWithOrderQuery(share model.Share, args ...interface{}) (checklists []*model.Checklist, err error) {
	argS, argArray := argsInit(args...)
	argS = append(argS, "share_id=?")
	argArray = append(argArray, share.ID)
	argArray[0] = strings.Join(argS, "AND")

	checklists, err = checklistModel.QueryBatch(argArray...)
	if len(checklists) == 0 {
		return
	}

	// 排序
	checklistsMap := make(map[uint]*model.Checklist, len(checklists))
	for _, checklist := range checklists {
		checklistsMap[checklist.ID] = checklist
	}

	checklists[0] = checklistsMap[share.ChecklistID]
	// 对 postID 排序，原则上应该检出多少个就排多少个
	postID := checklists[0].PostID
	for i := 1; i < len(checklists); i++ {
		checklists[i] = checklistsMap[postID]
		postID = checklists[i].PostID
	}
	return
}

// ChecklistCreate 创建
func ChecklistCreate(checklist *model.Checklist, share *model.Share) error {
	if checklist.UserID == 0 || checklist.ShareID == 0 {
		return fmt.Errorf("ChecklistCreate 创建检查项必须设置用户 ID 和 分享 ID")
	}
	if checklist.PrevID == 0 && share.ChecklistID > 0 {
		return fmt.Errorf("只支持在已存的检查项后面追加 share_id=%d", share.ID)
	}
	mydb := initializer.DB.Begin()
	err := mydb.Save(checklist).Error
	if err != nil {
		mydb.Rollback()
		return err
	}
	// 如果存在前一个检查项，修改前一个检查项，如果存在下一个检查项，修改下一个检查项
	if checklist.PrevID > 0 {
		preChecklist, err := checklistModel.QueryByID(checklist.PrevID)
		var postID uint = 0
		if err != nil && err != gorm.ErrRecordNotFound {
			mydb.Rollback()
			return fmt.Errorf("preChecklist 不存在 %s", err.Error())
		} else if err == nil {
			postID = preChecklist.PostID
			err = mydb.Model(preChecklist).Update(map[string]interface{}{"post_id": checklist.ID}).Error
			if err != nil {
				mydb.Rollback()
				return err
			}
		}

		// 更新当前的 postID
		if postID > 0 {
			err = mydb.Model(checklist).Update(map[string]interface{}{"post_id": postID}).Error
			if err != nil {
				mydb.Rollback()
				return err
			}

			postChecklist, err := checklistModel.QueryByID(preChecklist.PostID)
			if err != nil && err != gorm.ErrRecordNotFound {
				mydb.Rollback()
				return fmt.Errorf("postChecklist 不存在 %s", err.Error())
			} else if err == nil {
				err = mydb.Model(postChecklist).Update(map[string]interface{}{"pre_id": checklist.ID}).Error
				if err != nil {
					mydb.Rollback()
					return err
				}
			}
		}

	}
	if share.ChecklistID == 0 {
		err = mydb.Model(share).Update(map[string]interface{}{
			"checklist_id": checklist.ID,
		}).Error
		if err != nil {
			mydb.Rollback()
			return err
		}
	}
	return mydb.Commit().Error
}

// ChecklistUpdate 更新检查项
func ChecklistUpdate(checklistID uint, title string, user *model.User) error {
	checklist, err := checklistModel.QueryByID(checklistID)
	if err != nil {
		return fmt.Errorf("ChecklistUpdate 检索 checklist 错误: %s", err.Error())
	}
	share, err := ShareQueryByID(checklist.ShareID)
	if err != nil {
		return fmt.Errorf("ChecklistUpdate 检索 share 错误: %s", err.Error())
	}
	if user.ID != checklist.UserID || user.ID != share.UserID {
		return fmt.Errorf("ChecklistUpdate 更新检查项权限错误: %d", user.ID)
	}
	return checklist.Update("title", title)
}
