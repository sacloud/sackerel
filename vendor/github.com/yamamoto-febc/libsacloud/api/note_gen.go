package api

/************************************************
  generated by IDE. for [NoteAPI]
************************************************/

import (
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

/************************************************
   To support influent interface for Find()
************************************************/

func (api *NoteAPI) Reset() *NoteAPI {
	api.reset()
	return api
}

func (api *NoteAPI) Offset(offset int) *NoteAPI {
	api.offset(offset)
	return api
}

func (api *NoteAPI) Limit(limit int) *NoteAPI {
	api.limit(limit)
	return api
}

func (api *NoteAPI) Include(key string) *NoteAPI {
	api.include(key)
	return api
}

func (api *NoteAPI) Exclude(key string) *NoteAPI {
	api.exclude(key)
	return api
}

func (api *NoteAPI) FilterBy(key string, value interface{}) *NoteAPI {
	api.filterBy(key, value, false)
	return api
}

// func (api *NoteAPI) FilterMultiBy(key string, value interface{}) *NoteAPI {
// 	api.filterBy(key, value, true)
// 	return api
// }

func (api *NoteAPI) WithNameLike(name string) *NoteAPI {
	return api.FilterBy("Name", name)
}

func (api *NoteAPI) WithTag(tag string) *NoteAPI {
	return api.FilterBy("Tags.Name", tag)
}
func (api *NoteAPI) WithTags(tags []string) *NoteAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *NoteAPI) WithSizeGib(size int) *NoteAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

func (api *NoteAPI) WithSharedScope() *NoteAPI {
	api.FilterBy("Scope", "shared")
	return api
}

func (api *NoteAPI) WithUserScope() *NoteAPI {
	api.FilterBy("Scope", "user")
	return api
}

func (api *NoteAPI) SortBy(key string, reverse bool) *NoteAPI {
	api.sortBy(key, reverse)
	return api
}

func (api *NoteAPI) SortByName(reverse bool) *NoteAPI {
	api.sortByName(reverse)
	return api
}

// func (api *NoteAPI) SortBySize(reverse bool) *NoteAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

func (api *NoteAPI) New() *sacloud.Note {
	return &sacloud.Note{}
}

func (api *NoteAPI) Create(value *sacloud.Note) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *NoteAPI) Read(id int64) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *NoteAPI) Update(id int64, value *sacloud.Note) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *NoteAPI) Delete(id int64) (*sacloud.Note, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

/************************************************
  Inner functions
************************************************/

func (api *NoteAPI) setStateValue(setFunc func(*sacloud.Request)) *NoteAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *NoteAPI) request(f func(*sacloud.Response) error) (*sacloud.Note, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Note, nil
}

func (api *NoteAPI) createRequest(value *sacloud.Note) *sacloud.Request {
	req := &sacloud.Request{}
	req.Note = value
	return req
}
