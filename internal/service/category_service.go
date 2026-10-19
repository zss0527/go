package service

import (
	"bunny-fairy-api/internal/dto"
	"bunny-fairy-api/internal/models/home"
	"bunny-fairy-api/internal/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepo
}

// NewCategoryService 注入repo
func NewCategoryService(repo *repository.CategoryRepo) CategoryService {
	return CategoryService{categoryRepo: repo}
}

// GetCategoryTree 获取分类树，只调用repo，不碰db
func (svc *CategoryService) GetCategoryTree() ([]home.Category, error) {
	return svc.categoryRepo.FindRootCategoryTree()
}

// ImportTree 业务逻辑：解析DTO，组装数据，开启事务，调用repo批量入库
func (svc *CategoryService) ImportTree(treeList []dto.CategoryDTO) error {
	tx := svc.categoryRepo.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var rootCats []home.Category
	var childCats []home.Category
	var allGoods []home.Goods

	for _, rootDto := range treeList {
		// 组装一级分类model
		rootCat := home.Category{
			ID:       rootDto.ID,
			Name:     rootDto.Name,
			Picture:  rootDto.Picture,
			ParentID: nil,
		}
		rootCats = append(rootCats, rootCat)

		// 一级分类商品
		for _, g := range rootDto.Goods {
			g.CategoryID = rootDto.ID
			allGoods = append(allGoods, g)
		}

		// 二级分类
		for _, childDto := range rootDto.Children {
			childCat := home.Category{
				ID:       childDto.ID,
				Name:     childDto.Name,
				Picture:  childDto.Picture,
				ParentID: &rootDto.ID,
			}
			childCats = append(childCats, childCat)

			//二级分类的商品
			for _, g := range childDto.Goods {
				g.CategoryID = childDto.ID
				allGoods = append(allGoods, g)
			}
		}
	}

	// 交给repo执行事务插入
	err := svc.categoryRepo.ImportTreeTx(tx, rootCats, childCats, allGoods)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
