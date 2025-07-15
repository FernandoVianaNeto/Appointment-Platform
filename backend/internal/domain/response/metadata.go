package domain_response

type PaginationMetadata struct {
	CurrentPage int `json:"currentPage"`
	Next        int `json:"next"`
	Total       int `json:"total"`
	TotalItems  int `json:"totalItems"`
}

const DEFAULT_ITEMS_PER_PAGE = 16

func GetMetadataParams(currentPage int, totalItems int64) PaginationMetadata {
	totalPages := totalItems / DEFAULT_ITEMS_PER_PAGE

	if totalItems%DEFAULT_ITEMS_PER_PAGE != 0 {
		totalPages++
	}

	next := currentPage + 1

	if next > int(totalPages) {
		next = 0
	}

	return PaginationMetadata{
		CurrentPage: currentPage,
		Next:        next,
		Total:       int(totalPages),
		TotalItems:  int(totalItems),
	}
}
