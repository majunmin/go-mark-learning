/**
  @Author: majm@ushareit.com
  @date: 2021/1/26
  @note:
**/
package group_cache

import (
	"context"
	"fmt"
	"github.com/golang/groupcache"
	"log"
	"net/http"
)

type TblCache struct {
	Id    int
	Key   string
	Value string
}

func Demo1(local_addr string) {
	//定义节点数量以及地址
	peers_addrs := []string{"http://127.0.0.1:8001", "http://127.0.0.1:8002"}
	//db, _ := sql.Open("sqlite3", "./console.db")

	peers := groupcache.NewHTTPPool("http://" + local_addr)
	peers.Set(peers_addrs...)

	// 获取group对象
	image_cache := groupcache.NewGroup("testGroup", 8<<30,
		// 自定义数据获取来源
		groupcache.GetterFunc(
			func(ctx context.Context, key string, dest groupcache.Sink) error {
				//rows, _ := db.Query("SELECT key, value FROM tbl_cache_map where key = ?", key)
				//for rows.Next() {
				//	p := new(TblCache)
				//	err := rows.Scan(&p.Key, &p.Value)
				//	if err != nil {
				//		fmt.Println(err)
				//	}
				//	fmt.Printf("get %s of value from tbl_cache_map\n", key)
				dest.SetString("tbl_cache_map.value : " + "c")
				//}
				return nil
			}))

	// 定义返回方式
	http.HandleFunc("/get", func(rw http.ResponseWriter, r *http.Request) {
		var data []byte
		k := r.URL.Query().Get("key")
		fmt.Printf("user get %s of value from groupcache\n", k)
		ctx, _ := context.WithCancel(context.Background())
		image_cache.Get(ctx, k, groupcache.AllocatingByteSliceSink(&data))
		rw.Write(data)
	})

	log.Fatal(http.ListenAndServe(local_addr, nil))
}
