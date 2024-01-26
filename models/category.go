package models

import "materialsSort/global"

type Category struct {
	CategoryId   int    `gorm:"primary_key;AUTO_INCREMENT" json:"category_id"`
	CategoryName string `gorm:"type:varchar(20);not null" json:"category_name"`
}

func (Category) TableName() string {
	return "category"
}

// GetCategory 获取分类，有分页参数
func GetCategory(page int, pageSize int) (count int64, category []Category, err error) {
	// 查询总数
	if err = global.DB.Model(&Category{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}
	// 查询分页数据
	if err = global.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&category).Error; err != nil {
		return 0, nil, err
	}
	return count, category, nil
}

// GetCategoryAll 获取分类，无分页参数
func GetCategoryAll() (count int64, category []Category, err error) {
	// 查询总数
	if err = global.DB.Model(&Category{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}
	if err = global.DB.Find(&category).Error; err != nil {
		return 0, nil, err
	}
	return count, category, nil
}

// AddCategory 添加分类
func AddCategory(categoryName string) (err error) {
	category := Category{
		CategoryName: categoryName,
	}
	if err = global.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

// UpdateCategory 更新分类
func UpdateCategory(categoryId int, categoryName string) (err error) {
	if err = global.DB.Model(&Category{}).Where("category_id = ?", categoryId).Update("category_name", categoryName).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCategory 删除分类
func DeleteCategory(categoryId int) (err error) {
	if err = global.DB.Where("category_id = ?", categoryId).Delete(&Category{}).Error; err != nil {
		return err
	}
	return nil
}

// GetCategoryById 根据分类id获取分类
func GetCategoryById(categoryId int) (category Category, err error) {
	if err = global.DB.Where("category_id = ?", categoryId).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

// GetCategoryByName 根据分类名称获取分类
func GetCategoryByName(categoryName string) (category Category, err error) {
	if err = global.DB.Where("category_name = ?", categoryName).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
