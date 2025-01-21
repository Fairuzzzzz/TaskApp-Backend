package recipes

import "github.com/Fairuzzzzz/taskapp/internal/models/recipes"

func (r *repository) CreateREcipes(model recipes.Recipe) error {
	return r.db.Create(&model).Error
}

func (r *repository) GetAll() ([]recipes.Recipe, error) {
	var recipes []recipes.Recipe
	if err := r.db.Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *repository) GetByID(id uint) (*recipes.Recipe, error) {
	var recipes recipes.Recipe
	if err := r.db.First(&recipes, id).Error; err != nil {
		return nil, err
	}
	return &recipes, nil
}

func (r *repository) Update(model recipes.Recipe) error {
	return r.db.Model(&recipes.Recipe{}).Where("id = ?", model.ID).Updates(&model).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&recipes.Recipe{}, id).Error
}
