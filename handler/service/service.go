package service

import (
	"fmt"
	"github.com/youshintop/apiserver/model"
	"github.com/youshintop/apiserver/pkg/util"
	"sync"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, int64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.List(username, offset, limit)
	if err != nil {
		return nil, count, err
	}
	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(ids)),
	}

	errChan := make(chan error, 1)
	finsh := make(chan struct{}, 1)
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()
			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:       u.Id,
				Username: u.Username,
				Password: u.Password,
				SayHello: fmt.Sprintf("Hello %s", shortId),
				CreateAt: u.CreateAt.Format("2006-01-02 15:04:05"),
				UpdateAt: u.UpdateAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finsh)
	}()

	select {
	case <-finsh:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}
	return infos, count, nil
}
