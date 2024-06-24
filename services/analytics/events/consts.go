package events

const (
	E_SHOP           = "shop"
	E_DAY7_BUY       = "day7_buy"
	E_DAY7_REWARD    = "day7_reward"
	E_MCARD_PAY      = "mcard_pay"
	E_FIRST_PAY      = "first_day"
	E_GROW_ROAD      = "grow_road"
	E_UNLOCK_SKIN    = "unlock_skin"
	E_SIGN           = "sign"
	E_HERO_REHAVE    = "hero_rehave"
	E_PET_SKILL_DRAW = "pet_skill_draw"
)

const (
	PET_ACT_EGG_START  = 1  //蛋孵化开始
	PET_ACT_EGG_CANCEL = 2  //蛋孵化取消
	PET_ACT_EGG_SUCC   = 3  //蛋孵化成功
	PET_ACT_EGG_FEED   = 4  //蛋喂养
	PET_ACT_EGG_SPEED  = 5  //蛋加速
	PET_ACT_UP_START   = 6  //宠物洗练
	PET_ACT_UP_LOCK    = 7  //宠物洗练
	PET_ACT_UP_NO      = 8  //宠物洗练
	PET_ACT_UP_YES     = 9  //宠物洗练
	PET_ACT_FREE       = 10 //宠物放生
	PET_ACT_SKILL      = 11 //宠物技能石
	TASK_BEGIN         = 1
	TASK_READY         = 2
	TASK_FINISH        = 3
)

var OsTypes = map[int32]string{
	1: "1000", //安卓
	2: "1001", //ios
}

var OsTypeNames = map[int32]string{
	1: "android", //安卓
	2: "ios",     //ios
}
