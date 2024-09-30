package service

import (
	"errors"
	"shorturl/model"
	"shorturl/utils"
	"sync"
	"time"
)

type MapperService struct {
	lock        sync.Mutex
	cache       map[string]*model.Mapper
	sourceCache map[string]*model.Mapper
}

func BuildMapService() *MapperService {
	return &MapperService{cache: make(map[string]*model.Mapper), sourceCache: make(map[string]*model.Mapper)}
}

// 根据源url生成短链接对象
func (s *MapperService) Generate(url string) (mapper *model.Mapper, err error) {
	if _, exist := s.sourceCache[url]; exist {
		return s.sourceCache[url], nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	var short string = utils.RandStr(6)
	_, exist := s.cache[short]
	for exist {
		short = utils.RandStr(6)
		_, exist = s.cache[short]
	}
	mp := &model.Mapper{
		Id:         time.Now().UnixMicro(),
		ShortUrl:   short,
		Url:        url,
		CreateTime: time.Now(),
	}
	s.cache[short] = mp
	s.sourceCache[url] = mp
	return mp, nil
}

// 根据短连接对象获取源url
func (s *MapperService) GetShortUrl(shortUrl string) (url string, err error) {
	if _, exist := s.cache[shortUrl]; exist {
		return s.cache[shortUrl].Url, nil
	}
	return "", errors.New("short url not exist")
}

// 获取缓存的短链接列表
func (s *MapperService) List() (mappers []*model.Mapper, err error) {
	var list []*model.Mapper
	for _, v := range s.cache {
		list = append(list, v)
	}
	return list, nil
}
