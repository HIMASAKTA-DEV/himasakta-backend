package utils

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
)

func ResolveGallery(baseURL string, g *entity.Gallery) {
	if g != nil {
		g.ImageUrl = ResolveImageURL(baseURL, g.ImageUrl)
	}
}

func ResolveGalleries(baseURL string, gs []entity.Gallery) {
	for i := range gs {
		gs[i].ImageUrl = ResolveImageURL(baseURL, gs[i].ImageUrl)
	}
}

func ResolveMember(baseURL string, m *entity.Member) {
	if m == nil {
		return
	}
	ResolveGallery(baseURL, m.Photo)
}

func ResolveMembers(baseURL string, ms []entity.Member) {
	for i := range ms {
		ResolveGallery(baseURL, ms[i].Photo)
	}
}

func ResolveDepartment(baseURL string, d *entity.Department) {
	if d == nil {
		return
	}
	ResolveGallery(baseURL, d.Logo)
	ResolveGalleries(baseURL, d.Feeds)
	ResolveMember(baseURL, d.Leader)
}

func ResolveDepartments(baseURL string, ds []entity.Department) {
	for i := range ds {
		ResolveGallery(baseURL, ds[i].Logo)
		ResolveGalleries(baseURL, ds[i].Feeds)
		ResolveMember(baseURL, ds[i].Leader)
	}
}

func ResolveCabinetInfo(baseURL string, c *entity.CabinetInfo) {
	if c == nil {
		return
	}
	ResolveGallery(baseURL, c.Logo)
	ResolveGallery(baseURL, c.Organigram)
	ResolveGalleries(baseURL, c.Feeds)
}

func ResolveCabinetInfos(baseURL string, cs []entity.CabinetInfo) {
	for i := range cs {
		ResolveGallery(baseURL, cs[i].Logo)
		ResolveGallery(baseURL, cs[i].Organigram)
		ResolveGalleries(baseURL, cs[i].Feeds)
	}
}

func ResolveProgenda(baseURL string, p *entity.Progenda) {
	if p == nil {
		return
	}
	ResolveGallery(baseURL, p.Thumbnail)
	ResolveGalleries(baseURL, p.Feeds)
}

func ResolveProgendas(baseURL string, ps []entity.Progenda) {
	for i := range ps {
		ResolveGallery(baseURL, ps[i].Thumbnail)
		ResolveGalleries(baseURL, ps[i].Feeds)
	}
}

func ResolveNews(baseURL string, n *entity.News) {
	if n == nil {
		return
	}
	ResolveGallery(baseURL, n.Thumbnail)
}

func ResolveNewsList(baseURL string, ns []entity.News) {
	for i := range ns {
		ResolveGallery(baseURL, ns[i].Thumbnail)
	}
}

func ResolveMonthlyEvent(baseURL string, e *entity.MonthlyEvent) {
	if e == nil {
		return
	}
	ResolveGallery(baseURL, e.Thumbnail)
}

func ResolveMonthlyEvents(baseURL string, es []entity.MonthlyEvent) {
	for i := range es {
		ResolveGallery(baseURL, es[i].Thumbnail)
	}
}


