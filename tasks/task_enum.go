package tasks

type TaiJiTaskStep int32

const (
	TaiJiTaskStep_Hundun   TaiJiTaskStep = 0 // 混沌
	TaiJiTaskStep_YangMing TaiJiTaskStep = 1 // 阳明
	TaiJiTaskStep_TaiYang  TaiJiTaskStep = 2 // 太阳
	TaiJiTaskStep_ShaoYang TaiJiTaskStep = 3 // 少阳
	TaiJiTaskStep_TaiYin   TaiJiTaskStep = 4 // 太阴
	TaiJiTaskStep_ShaoYin  TaiJiTaskStep = 5 // 少阴
	TaiJiTaskStep_JueYin   TaiJiTaskStep = 6 // 厥阴
	TaiJiTaskStep_WuJi     TaiJiTaskStep = 7 // 无极
)

type BaGuaTaskStep int32

const (
	BaGuaTaskStep_ZhongGong BaGuaTaskStep = 0 // 中宫
	BaGuaTaskStep_Qian      BaGuaTaskStep = 1 // 乾
	BaGuaTaskStep_Xun       BaGuaTaskStep = 2 // 巽
	BaGuaTaskStep_Kan       BaGuaTaskStep = 3 // 坎
	BaGuaTaskStep_Gen       BaGuaTaskStep = 4 // 艮
	BaGuaTaskStep_Kun       BaGuaTaskStep = 5 // 坤
	BaGuaTaskStep_Zhen      BaGuaTaskStep = 6 // 震
	BaGuaTaskStep_Li        BaGuaTaskStep = 7 // 离
	BaGuaTaskStep_Dui       BaGuaTaskStep = 8 // 兑
)
