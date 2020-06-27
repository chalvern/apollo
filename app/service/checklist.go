package service

import (
	"github.com/chalvern/apollo/app/model"
)

var (
	checklistModel = model.Checklist{}
)

// ChecklistsWithOrderQuery 获取某个分享的 checklist
// 按照双链表的先后次序返回
func ChecklistsWithOrderQuery(shareID uint, args ...interface{}) (checklists []*model.Checklist, err error) {
	share, err := ShareQueryByID(shareID)
	if err != nil {
		return
	}
	checklists, err = checklistModel.QueryBatch(args...)
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
	for i := range checklists[1:] {
		checklists[i] = checklistsMap[postID]
		postID = checklists[i].PostID
	}
	return
}
