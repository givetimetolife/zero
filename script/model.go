package script

import (
	"context"
	"encoding/json"
)

type (
	ZeroData struct {
		Title       string `json:"title,omitempty"`
		Link        string `json:"link,omitempty"`
		ExcerptArea string `json:"excerpt_area,omitempty"`

		Metadata string `json:"-"`
	}

	ZeroDataInterface interface {
		Data(ctx context.Context) ([]*ZeroData, error)
	}
)

func (z *zhihuData) convertToZd() []*ZeroData {
	zds := []*ZeroData{}

	if z != nil {
		for _, v := range z.Data {

			metadata, _ := json.Marshal(z)

			zds = append(zds, &ZeroData{
				Title:       v.Target.Title,
				ExcerptArea: v.Target.Excerpt,
				Link:        v.Target.URL,

				// Title:       v.Title,
				// ExcerptArea: v.ExcerptArea,
				// Link:        v.Link,
				Metadata: string(metadata),
			})
		}
	}
	return zds
}

var ZeroDataMap = map[string]ZeroDataInterface{
	"zhihu": &zhihu{},
}
