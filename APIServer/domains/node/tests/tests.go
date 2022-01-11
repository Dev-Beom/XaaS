package tests

import (
	"github.com/dev-beom/xaas/apiserver/models"
	"math/rand"
	"strconv"
	"time"
)

var (
	mockDB                = make(map[string]models.Node)
	nodeCreateRequestBody = models.Node{
		Id:          "test_id",
		Description: "test_description",
	}
)

func initMockNodeDataForTest() {
	for i := 0; i < 100; i++ {
		key := strconv.Itoa(i) + "_test"
		mockDB[key] = models.Node{
			Id:          key,
			Description: key + "_description",
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
			State:       models.Creating,
		}
	}
}

func getRandomID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(100)) + "_test"
}

func getMockNode(id string) models.Node {
	return mockDB[id]
}
