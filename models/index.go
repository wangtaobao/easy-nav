package models

import "strconv"

func CreateIndex(index *IndexReply) error {

	s := new(IndexReply)
	s.Url = index.Url
	s.Name = index.Name
	err := DB.Create(&s).Error

	if err != nil {
		return err
	}
	return nil

}

func GetIndexList() ([]IndexReply, error) {
	var indexs []IndexReply
	err := DB.Find(&indexs).Error
	if err != nil {
		return nil, err
	}
	return indexs, nil
}

func DelIndexById(id *string) error {
	idd, _ := strconv.ParseUint(*id, 10, 32)

	index := new(IndexReply)
	index.ID = uint(idd)
	return DB.Delete(&index).Error
}

func UpdateIndex(index *IndexReply) error {
	s := new(IndexReply)
	s.Url = index.Url
	s.Name = index.Name
	s.ID = index.ID

	return DB.Model(&s).Updates(&s).Error // 仅更新非零值字段

}
