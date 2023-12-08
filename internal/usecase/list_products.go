package usecase

import "mensageria/internal/entity"

type ListProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductRepository(productRepository entity.ProductRepository) *ListProductUseCase {
	return &ListProductUseCase{ProductRepository: productRepository}
}

func (u *ListProductUseCase) Execute() ([]*ListProductOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductOutputDto

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOutput, nil
}
