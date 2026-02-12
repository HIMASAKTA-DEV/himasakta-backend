package meta

import (
	"strings"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
	"github.com/gin-gonic/gin"
)

// use for pagination
type Meta struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	TotalData int    `json:"total_data"`
	TotalPage int    `json:"total_page"`
	Sort      string `json:"sort"`
	SortBy    string `json:"sort_by"`
	Filter    string `json:"filter,omitempty"`
	FilterBy  string `json:"filter_by,omitempty"`
}

// New creates and initializes a Meta object with default pagination settings.
// Default values are:
// - Limit: 10 (number of items per page)
// - Page: 0 (starting page)
// - Sort: "asc" (ascending order)
// - SortBy: "created_by" (column used for sorting)
// Additional options can be applied to customize the Meta object.
func New(ctx *gin.Context) Meta {
	meta := Meta{
		Limit:  10,
		Page:   0,
		Sort:   "asc",
		SortBy: "id",
	}

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	sort := ctx.Query("sort")
	sortby := ctx.Query("sort_by")
	filter := ctx.Query("filter")
	filterby := ctx.Query("filter_by")

	if page != "" {
		meta.Page = utils.ToInt(page)
	}

	if limit != "" {
		meta.Limit = utils.DefaultLimit(utils.ToInt(limit))
	}

	if sort != "" {
		meta.Sort = sort
	}

	if sortby != "" {
		meta.SortBy = sortby
	}

	if filter != "" {
		meta.Filter = filter
	}

	if filterby != "" {
		meta.FilterBy = filterby
	}

	return meta
}

func NewWithDefault(ctx *gin.Context, dtake int, dpage int, dsort string, dsortBy string) Meta {
	if dtake == 0 {
		dtake = 10
	}

	if dpage == 0 {
		dpage = 0
	}

	if dsort == "" {
		dsort = "asc"
	}

	if dsortBy == "" {
		dsortBy = "id"
	}

	meta := Meta{
		Limit:  dtake,
		Page:   dpage,
		Sort:   dsort,
		SortBy: dsortBy,
	}

	page := ctx.Query("page")
	limit := ctx.Query("limit")
	sort := ctx.Query("sort")
	sortby := ctx.Query("sort_by")
	filter := ctx.Query("filter")
	filterby := ctx.Query("filter_by")

	if page != "" {
		meta.Page = utils.ToInt(page)
	}

	if limit != "" {
		meta.Limit = utils.DefaultLimit(utils.ToInt(limit))
	}

	if sort != "" {
		meta.Sort = sort
	}

	if sortby != "" {
		meta.SortBy = sortby
	}

	if filter != "" {
		meta.Filter = filter
	}

	if filterby != "" {
		meta.FilterBy = filterby
	}

	return meta
}

// Count calculates the total number of pages based on the total data count.
// It sets the TotalData and TotalPage fields in the Meta struct.
func (m *Meta) Count(totaldata int) {
	m.TotalData = totaldata
	m.TotalPage = (totaldata + m.Limit - 1) / m.Limit
}

// GetSkipAndLimit calculates the offset (skip) and limit values for pagination.
// If the page number is less than or equal to 0, skip is set to 0.
// Otherwise, skip is calculated as (page - 1) * limit, and the limit is set to the limit value.
func (m *Meta) GetSkipAndLimit() (int, int) {
	switch {
	case m.Page <= 0:
		m.Page = 1
		return 0, m.Limit
	default:
		return ((m.Page - 1) * m.Limit), m.Limit
	}
}

func (m Meta) SeparateFilter() map[string]string {
	filtersBy := strings.Split(m.FilterBy, ",")
	filters := strings.Split(m.Filter, ",")

	filterMap := map[string]string{}
	for i, filterBy := range filtersBy {
		filterMap[filterBy] = filters[i]
	}

	return filterMap
}

func (m *Meta) SetSort(sort string) {
	m.Sort = sort
}

func (m *Meta) SetSortBy(sortBy string) {
	m.SortBy = sortBy
}
