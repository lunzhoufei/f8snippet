package main

import (
	"database/sql"
	"fmt"
	"git.code.oa.com/going/going/codec/mysql"
	"git.code.oa.com/tmemesh.io/ckv"
	"git.code.oa.com/tmemesh.io/srf/srfs"
	"git.code.oa.com/tmerpc/jce/area"
)

// AreaDefaultSortSyncImpl 结构体
type AreaDefaultSortSyncImpl struct {
}

// Initialize 接口 : 初始化
func (imp *AreaDefaultSortSyncImpl) Initialize() {
	/*
	   初始化
	*/
}

const (
	// 查询所有见面吧电台的视频mvid 2019年播放量
	sqlQueryJianmianbaMVID = "t_live_video_play_cnt where type = 18"
	sqlQuery               = "t_area_default_sort where STATE = 1 order by ID desc limit 1"
	ckvKey                 = "area_default_sort"
	testEnv                = 0
	productionEvn          = 1
)

// ============================================================================
//                              get one
// ============================================================================

// +--------------+--------------+------+-----+---------+----------------+
// | Field        | Type         | Null | Key | Default | Extra          |
// +--------------+--------------+------+-----+---------+----------------+
// | ID           | int(11)      | NO   | PRI | NULL    | auto_increment |
// | TITLE        | varchar(255) | YES  |     |         |                |
// | STATE        | varchar(255) | YES  |     |         |                |
// | Finsert_user | varchar(255) | YES  |     |         |                |
// | Finsert_time | datetime     | YES  |     | NULL    |                |
// | Fmodify_user | varchar(255) | YES  |     |         |                |
// | Fmodify_time | datetime     | YES  |     | NULL    |                |
// | ID_LIST      | mediumtext   | YES  |     | NULL    |                |
// +--------------+--------------+------+-----+---------+----------------+

// 数据库中的数据
type mmData struct {
	ID          int            `sql:"ID"`
	Title       sql.NullString `sql:"TITLE"`
	State       sql.NullString `sql:"STATE"`
	FinsertUser sql.NullString `sql:"Finsert_user"`
	FinsertTime sql.NullString `sql:"Finsert_time"`
	FmodifyUser sql.NullString `sql:"Fmodify_user"`
	FmodifyTime sql.NullString `sql:"Fmodify_time"`
	IDList      sql.NullString `sql:"ID_LIST"`
}

// 运行数据
type idListInfoStruct struct {
	data *mmData
}

//
func (idListInfo *idListInfoStruct) readOne(ctx *srfs.Context) error {
	conn := mysql.New("redstone")
	if conn == nil {
		err := fmt.Errorf("Can not connect sql server")
		fmt.Println(err)
		return err
	}
	datas, err := conn.Select(ctx, (*mmData)(nil), sqlQuery)
	if err != nil {
		fmt.Println("Can not select sql server")
		return err
	}
	if len(datas) != 1 {
		err = fmt.Errorf("Sorry, there is no validity rule")
		fmt.Println(err)
		return err
	}
	idListInfo.data = datas[0].(*mmData)
	fmt.Println("Successful read database data", *idListInfo.data)
	return nil
}

// ============================================================================
//                              get multi
// ============================================================================

// 播放量统计相关数据
type playAmountStruct struct {
	// 要查询的gid集合
	gids map[int]bool
	// gid对应的播放量
	playCountMap map[int]int
	// 总播放量
	PlayAmount int
	// 存储文件名
	fileName string
	// 文件时间戳
	fileTail string
	// 配置文件
	HdfsConfig struct {
		HdfsCmdArgsFormat string
	}
}

type jianmianbaDBData struct {
	Mvid sql.NullString `sql:"comp_mvid"`
}

// VideoBaseStruct 上传者在数据库中的每行的数据
type VideoBaseStruct struct {
	ID   int `sql:"ID"` // 视频GID
	MVID int `sql:"Fsource_id"`
}

func (p *playAmountStruct) ReadMulti(ctx *srfs.Context) error {
	mvidList := make([]string, 0)
	videoPlayCntdb := mysql.New("videoplaycnt")
	rows, err := videoPlayCntdb.Select(ctx, (*jianmianbaDBData)(nil), sqlQueryJianmianbaMVID)
	if err != nil {
		fmt.Println("err1 is ", err)
		return err
	}
	// 提取mvid
	for _, v := range rows {
		if v.(*jianmianbaDBData).Mvid.String == "" {
			continue
		}
		fmt.Println(v.(*jianmianbaDBData).Mvid.String)
		mvidList = append(mvidList, v.(*jianmianbaDBData).Mvid.String)
	}
	// 查gid
	sqlwherequery := fmt.Sprintf("where Fsource_id in (%s)", strings.Join(mvidList, ","))
	videoBaseDb := mysql.New("videobase")
	fmt.Println(sqlwherequery)
	rows, err = videoBaseDb.Select(ctx, (*VideoBaseStruct)(nil), "t_video_base "+sqlwherequery)
	if err != nil {
		return err
	}
	for _, v := range rows {
		p.gids[(v.(*VideoBaseStruct).ID)] = true
	}
	return nil
}
