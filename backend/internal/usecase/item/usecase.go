package item

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	gachaRepo "github.com/pro-active/cta_gacha/internal/repository/gacha"
	itemRepo "github.com/pro-active/cta_gacha/internal/repository/item"
	userRepo "github.com/pro-active/cta_gacha/internal/repository/user"
	uhiRepo "github.com/pro-active/cta_gacha/internal/repository/userhaveitem"
	"github.com/pro-active/cta_gacha/util"
)

type ItemUsecase interface {
	GetItem(id string) (*ItemOutput, error)
	GetItems(gachaID string) ([]*ItemOutput, error)
	GetUserItems(userID string) ([]*ItemOutput, error)
	CreateItem(name, gachaID, userID, imgPath, rairty string) (*ItemOutput, error)
	DrawGacha(userID string, gachaID string) (*ItemOutput, error)
}

type itemUsecase struct {
	repository             itemRepo.ItemRepository
	userRepository         userRepo.UserRepository
	gachaRepository        gachaRepo.GachaRepository
	userhaveitemRepository uhiRepo.UserHaveItemRepository
}

func NewItemUsecase(i itemRepo.ItemRepository, u userRepo.UserRepository, g gachaRepo.GachaRepository, uhi uhiRepo.UserHaveItemRepository) *itemUsecase {
	return &itemUsecase{
		repository:             i,
		userRepository:         u,
		gachaRepository:        g,
		userhaveitemRepository: uhi,
	}
}

func (i *itemUsecase) GetItem(id string) (*ItemOutput, error) {
	result, err := i.repository.GetItem(id)
	if err != nil {
		return nil, err
	}
	return convertItem(result), err
}

func (i *itemUsecase) GetItems(gachaID string) ([]*ItemOutput, error) {
	results, err := i.repository.GetItems(gachaID)
	if err != nil {
		return nil, err
	}
	items := make([]*ItemOutput, 0, len(results))
	for i := range results {
		items = append(items, convertItem(results[i]))
	}
	return items, err
}

func (i *itemUsecase) GetUserItems(userID string) ([]*ItemOutput, error) {
	results, err := i.userhaveitemRepository.GetUserHaveItemsByUserID(userID)
	if err != nil {
		return nil, err
	}
	items := make([]*ItemOutput, 0, len(results))
	for j := range results {
		itemid := results[j].ItemID
		item, err := i.repository.GetItem(itemid)
		if err != nil {
			return nil, err
		}
		items = append(items, (convertItem(item)))
	}
	return items, err
}

func (i *itemUsecase) CreateItem(name, gachaID, userID, imgPath, rairty string) (*ItemOutput, error) {
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	result, err := i.repository.CreateItem(id, name, gachaID, userID, imgPath, rairty)
	if err != nil {
		return nil, err
	}
	return convertItem(result), err
}

func (g *itemUsecase) DrawGacha(userID string, gachaID string) (*ItemOutput, error) {
	// ユーザの存在確認
	user, err := g.userRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New(fmt.Sprintf("user not found: userID=%s", userID))
	}
	// ガチャの存在確認
	gacha, err := g.gachaRepository.GetGacha(gachaID)
	if err != nil {
		return nil, err
	}
	if gacha == nil {
		return nil, errors.New(fmt.Sprintf("gacha not found: gachaID=%s", gachaID))
	}
	// ガチャ実行
	// item全部とってくる
	items, err := g.repository.GetItems(gachaID)
	if err != nil {
		return nil, err
	}
	// 1個選ぶ
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(items))

	// userHaveItemに登録
	id, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	_, err = g.userhaveitemRepository.CreateUserHaveItem(id, userID, items[num].ID)
	if err != nil {
		return nil, err
	}
	return convertItem(items[num]), nil
}
