package repository

import (
	"bunny-fairy-api/internal/models/home"

	"gorm.io/gorm"
)

// CategoryRepo 分类仓库，只做数据库操作，无业务逻辑
type CategoryRepo struct {
	Db *gorm.DB
}

// NewCategoryRepo 构造方法，注入db
func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		Db: db,
	}
}

// FindRootCategoryTree 查询一级分类树，预加载二级子分类、对应商品
// 一级分类 parent_id is null
func (r *CategoryRepo) FindRootCategoryTree() ([]home.Category, error) {
	var list []home.Category

	err := r.Db.
		Where("parent_id IS NULL").
		Order("sort ASC").
		Preload("Children", func(tx *gorm.DB) *gorm.DB {
			return tx.Order("sort ASC")
		}).
		Preload("Goods").
		Find(&list).Error

	return list, err
}

// ImportTreeTx 在事务内批量插入分类和商品数据
func (r *CategoryRepo) ImportTreeTx(tx *gorm.DB, rootCats []home.Category, childCats []home.Category, goodsList []home.Goods) error {
	// 批量插入一级分类
	if len(rootCats) > 0 {
		if err := tx.CreateInBatches(rootCats, 100).Error; err != nil {
			return err
		}
	}
	// 批量插入二级分类
	if len(childCats) > 0 {
		if err := tx.CreateInBatches(childCats, 100).Error; err != nil {
			return err
		}
	}
	// 批量插入商品
	if len(goodsList) > 0 {
		if err := tx.CreateInBatches(goodsList, 100).Error; err != nil {
			return err
		}
	}
	return nil
}

// FindByID 根据id查询单个分类
func (r *CategoryRepo) FindByID(id string) (*home.Category, error) {
	var cat home.Category
	err := r.Db.First(&cat, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &cat, nil
}
