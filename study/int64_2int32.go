package main

import "fmt"

type ShopMgr struct {
}

var shopMgr *ShopMgr

func init() {
	shopMgr = &ShopMgr{}
}

func (mgr *ShopMgr) transShopId2Mixed(shopId uint32, subShopId uint32) int64 {
	return int64(shopId)<<32 + int64(subShopId)
}

func (mgr *ShopMgr) transShopIdFromMixed(shopIdMixed int64) (uint32, uint32) {
	var shopId = uint32(shopIdMixed >> 32)
	var subShopId = uint32(shopIdMixed)
	return shopId, subShopId
}

func main() {
	idMixed := shopMgr.transShopId2Mixed(123, 456)
	fmt.Println("idMixed:", idMixed)
	id1, id2 := shopMgr.transShopIdFromMixed(idMixed)
	fmt.Println("id1", id1, " id2:", id2)
}
