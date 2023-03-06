package config

import "github.com/NguyenXuanCanh/go-starter/types"

// const API_KEY = "AIzaSyCbtb5zcnXVF_t1l3uckddhm_1sSjWWCDc"
// const API_KEY = "rvWoa97j8PhzM5VUA0cr1IGNNNm5X81HoIN8GET6"
const API_KEY = "01f60aa129b28e378bc7301a5298bb64598d207211f13340"
const API_ROUTE_KEY = "rvWoa97j8PhzM5VUA0cr1IGNNNm5X81HoIN8GET6"

func GetDefaultStoreLocation() types.Location {
	var store types.Location
	store = append(store, 106.112456)
	store = append(store, 10.684922)
	return store
}
