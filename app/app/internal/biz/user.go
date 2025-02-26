package biz

import (
	"context"
	v1 "dhb/app/app/api"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID                     int64
	Address                string
	Password               string
	Undo                   int64
	AddressTwo             string
	PrivateKey             string
	AddressThree           string
	WordThree              string
	PrivateKeyThree        string
	Last                   uint64
	Amount                 uint64
	AmountBiw              uint64
	Total                  uint64
	IsDelete               int64
	Out                    int64
	CreatedAt              time.Time
	Lock                   int64
	AmountUsdt             float64
	MyTotalAmount          float64
	AmountUsdtGet          float64
	AmountRecommendUsdtGet float64
}

type UserInfo struct {
	ID               int64
	UserId           int64
	Vip              int64
	UseVip           int64
	HistoryRecommend int64
	TeamCsdBalance   int64
}

type UserRecommend struct {
	ID            int64
	UserId        int64
	RecommendCode string
	Total         int64
	CreatedAt     time.Time
}

type UserRecommendArea struct {
	ID            int64
	RecommendCode string
	Num           int64
	CreatedAt     time.Time
}

type Trade struct {
	ID           int64
	UserId       int64
	AmountCsd    int64
	RelAmountCsd int64
	AmountHbs    int64
	RelAmountHbs int64
	Status       string
	CreatedAt    time.Time
}

type UserArea struct {
	ID         int64
	UserId     int64
	Amount     int64
	SelfAmount int64
	Level      int64
}

type UserCurrentMonthRecommend struct {
	ID              int64
	UserId          int64
	RecommendUserId int64
	Date            time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type UserBalance struct {
	ID                  int64
	UserId              int64
	BalanceUsdt         int64
	BalanceUsdt2        int64
	BalanceDhb          int64
	RecommendTotal      int64
	AreaTotal           int64
	FourTotal           int64
	LocationTotal       int64
	BalanceC            int64
	RecommendTotalFloat float64
	AreaTotalFloat      float64
	AreaTwoTotalFloat   float64
	LocationTotalFloat  float64
	BalanceUsdtFloat    float64
	BalanceRawFloat     float64
}

type Withdraw struct {
	ID              int64
	UserId          int64
	Amount          int64
	AmountNew       float64
	RelAmountNew    float64
	RelAmount       int64
	BalanceRecordId int64
	Status          string
	Type            string
	Address         string
	CreatedAt       time.Time
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

type UserUseCase struct {
	repo                          UserRepo
	urRepo                        UserRecommendRepo
	configRepo                    ConfigRepo
	uiRepo                        UserInfoRepo
	ubRepo                        UserBalanceRepo
	locationRepo                  LocationRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type LocationNew struct {
	ID                int64
	UserId            int64
	Num               int64
	Status            string
	Current           int64
	CurrentMax        int64
	CurrentMaxNew     int64
	StopLocationAgain int64
	OutRate           int64
	Count             int64
	StopCoin          int64
	Top               int64
	Usdt              int64
	Total             int64
	TotalTwo          int64
	TotalThree        int64
	Biw               int64
	TopNum            int64
	LastLevel         int64
	StopDate          time.Time
	CreatedAt         time.Time
}

type UserBalanceRecord struct {
	ID        int64
	UserId    int64
	Amount    int64
	CoinType  string
	Balance   int64
	Type      string
	CreatedAt time.Time
}

type BalanceReward struct {
	ID        int64
	UserId    int64
	Status    int64
	Amount    int64
	SetDate   time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Reward struct {
	ID               int64
	UserId           int64
	Amount           int64
	AmountB          int64
	BalanceRecordId  int64
	Type             string
	TypeRecordId     int64
	Reason           string
	ReasonLocationId int64
	LocationType     string
	AmountNew        float64
	AmountNewTwo     float64
	CreatedAt        time.Time
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type ConfigRepo interface {
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetConfigs(ctx context.Context) ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
}

type UserBalanceRepo interface {
	RecommendLocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	CreateUserBalanceLock(ctx context.Context, u *User) (*UserBalance, error)
	LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error
	SystemReward(ctx context.Context, amount int64, locationId int64) error
	SystemFee(ctx context.Context, amount int64, locationId int64) error
	GetSystemYesterdayDailyReward(ctx context.Context) (*Reward, error)
	UserFee(ctx context.Context, userId int64, amount int64) (int64, error)
	RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	Deposit(ctx context.Context, userId int64, amount int64) (int64, error)
	DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error)
	DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetRewardFourYes(ctx context.Context) (*Reward, error)
	GetUserRewardByUserId(ctx context.Context, userId int64) ([]*Reward, error)
	GetLocationsToday(ctx context.Context) ([]*LocationNew, error)
	GetUserRewardByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserSortRecommendReward, error)
	GetUserRewards(ctx context.Context, b *Pagination, userId int64) ([]*Reward, error, int64)
	GetUserRewardsLastMonthFee(ctx context.Context) ([]*Reward, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceUsdtTotal(ctx context.Context) (int64, error)
	GreateWithdraw(ctx context.Context, userId int64, relAmount float64, amount float64, coinType string, address string) (*Withdraw, error)
	WithdrawUsdt(ctx context.Context, userId int64, amount int64, tmpRecommendUserIdsInt []int64) error
	WithdrawUsdt2(ctx context.Context, userId int64, amount float64) error
	Exchange(ctx context.Context, userId int64, amountUsdt, fee, amountRawSub float64) error
	Exchange2(ctx context.Context, userId int64, amount int64, amountUsdtSubFee int64, amountUsdt int64, locationId int64) error
	WithdrawUsdt3(ctx context.Context, userId int64, amount int64) error
	TranUsdt(ctx context.Context, userId int64, toUserId int64, amount int64, tmpRecommendUserIdsInt []int64, tmpRecommendUserIdsInt2 []int64) error
	WithdrawDhb(ctx context.Context, userId int64, amount int64) error
	WithdrawC(ctx context.Context, userId int64, amount int64) error
	TranDhb(ctx context.Context, userId int64, toUserId int64, amount int64) error
	GetWithdrawByUserId(ctx context.Context, userId int64, typeCoin string) ([]*Withdraw, error)
	GetWithdrawByUserId2(ctx context.Context, userId int64) ([]*Withdraw, error)
	GetUserBalanceRecordByUserId(ctx context.Context, userId int64, typeCoin string, tran string) ([]*UserBalanceRecord, error)
	GetUserBalanceRecordsByUserId(ctx context.Context, userId int64) ([]*UserBalanceRecord, error)
	GetTradeByUserId(ctx context.Context, userId int64) ([]*Trade, error)
	GetWithdraws(ctx context.Context, b *Pagination, userId int64) ([]*Withdraw, error, int64)
	GetWithdrawPassOrRewarded(ctx context.Context) ([]*Withdraw, error)
	UpdateWithdraw(ctx context.Context, id int64, status string) (*Withdraw, error)
	GetWithdrawById(ctx context.Context, id int64) (*Withdraw, error)
	GetWithdrawNotDeal(ctx context.Context) ([]*Withdraw, error)
	GetUserBalanceRecordUserUsdtTotal(ctx context.Context, userId int64) (int64, error)
	GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error)
	GetUserRewardUsdtTotal(ctx context.Context) (int64, error)
	GetSystemRewardUsdtTotal(ctx context.Context) (int64, error)
	UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*Withdraw, error)
	GetUserRewardRecommendSort(ctx context.Context) ([]*UserSortRecommendReward, error)
	GetUserRewardTodayTotalByUserId(ctx context.Context, userId int64) (*UserSortRecommendReward, error)

	SetBalanceReward(ctx context.Context, userId int64, amount int64) error
	UpdateBalanceReward(ctx context.Context, userId int64, id int64, amount int64, status int64) error
	GetBalanceRewardByUserId(ctx context.Context, userId int64) ([]*BalanceReward, error)

	GetUserBalanceLock(ctx context.Context, userId int64) (*UserBalance, error)
	Trade(ctx context.Context, userId int64, amount int64, amountB int64, amountRel int64, amountBRel int64, tmpRecommendUserIdsInt []int64, amount2 int64) error
}

type UserRecommendRepo interface {
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	GetUserRecommendByUserId(ctx context.Context, userId int64) (*UserRecommend, error)
	GetUserRecommendsFour(ctx context.Context) ([]*UserRecommend, error)
	CreateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (*UserRecommend, error)
	UpdateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCodeSum(ctx context.Context, code string) (int64, error)
	CreateUserRecommendArea(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	DeleteOrOriginUserRecommendArea(ctx context.Context, code string, originCode string) (bool, error)
	GetUserRecommendLowArea(ctx context.Context, code string) ([]*UserRecommendArea, error)
	GetUserAreas(ctx context.Context, userIds []int64) ([]*UserArea, error)
	CreateUserArea(ctx context.Context, u *User) (bool, error)
	GetUserArea(ctx context.Context, userId int64) (*UserArea, error)
	UpdateUserRecommendTotal(ctx context.Context, userId int64, total int64) error
}

type UserCurrentMonthRecommendRepo interface {
	GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *Pagination, userId int64) ([]*UserCurrentMonthRecommend, error, int64)
	CreateUserCurrentMonthRecommend(ctx context.Context, u *UserCurrentMonthRecommend) (*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error)
	GetUserLastMonthRecommend(ctx context.Context) ([]int64, error)
}

type UserInfoRepo interface {
	CreateUserInfo(ctx context.Context, u *User) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, u *UserInfo) (*UserInfo, error)
	GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserInfo, error)
}

type UserRepo interface {
	GetEthUserRecordListByUserId(ctx context.Context, userId int64) ([]*EthUserRecord, error)
	InRecordNew(ctx context.Context, userId int64, address string, amount int64, coinType string) error
	UpdateUserNewTwoNew(ctx context.Context, userId int64, amountUsdt float64, last uint64, amountRawFloat float64, coinType string) error
	UpdateUserMyTotalAmount(ctx context.Context, userId int64, amountUsdt float64) error
	UpdateUserMyTotalAmountSub(ctx context.Context, userId int64, amountUsdt float64) error
	UpdateUserRewardAreaTwo(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool) (int64, error)
	GetUserById(ctx context.Context, Id int64) (*User, error)
	GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetUserByUserIdsTwo(ctx context.Context, userIds []int64) (map[int64]*User, error)
	GetUsers(ctx context.Context, b *Pagination, address string) ([]*User, error, int64)
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserCount(ctx context.Context) (int64, error)
	GetUserCountToday(ctx context.Context) (int64, error)
}

func NewUserUseCase(repo UserRepo, tx Transaction, configRepo ConfigRepo, uiRepo UserInfoRepo, urRepo UserRecommendRepo, locationRepo LocationRepo, userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                          repo,
		tx:                            tx,
		configRepo:                    configRepo,
		locationRepo:                  locationRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		uiRepo:                        uiRepo,
		urRepo:                        urRepo,
		ubRepo:                        ubRepo,
		log:                           log.NewHelper(logger),
	}
}

func (uuc *UserUseCase) GetUserByAddress(ctx context.Context, Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(ctx, Addresses...)
}

func (uuc *UserUseCase) GetDhbConfig(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "level1Dhb", "level2Dhb", "level3Dhb")
}

func (uuc *UserUseCase) GetExistUserByAddressOrCreate(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user          *User
		recommendUser *UserRecommend
		err           error
		userId        int64
		//decodeBytes   []byte
	)

	user, err = uuc.repo.GetUserByAddress(ctx, u.Address) // 查询用户
	if nil == user || nil != err {
		code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
		if "abf00dd52c08a9213f225827bc3fb100" != code {
			//decodeBytes, err = base64.StdEncoding.DecodeString(code)
			//code = string(decodeBytes)
			//if 1 >= len(code) {
			//	return nil, errors.New(500, "USER_ERROR", "无效的推荐码1")
			//}
			//if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
			//	return nil, errors.New(500, "USER_ERROR", "无效的推荐码2")
			//}

			//var (
			//	locationNew []*LocationNew
			//)
			//locationNew, err = uuc.locationRepo.GetLocationsByUserId(ctx, userId)
			//if nil != err {
			//	return nil, errors.New(500, "USER_ERROR", "推荐人未入金")
			//}
			//
			//if 0 == len(locationNew) {
			//	return nil, errors.New(500, "USER_ERROR", "推荐人未入金")
			//}
			var (
				userRecommend *User
			)

			userRecommend, err = uuc.repo.GetUserByAddress(ctx, code)
			if nil == userRecommend || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1")
			}

			if 0 >= userRecommend.AmountUsdt {
				return nil, errors.New(500, "USER_ERROR", "推荐人未入金")
			}

			// 查询推荐人的相关信息
			recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
			if nil == recommendUser || err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码3")
			}
		}

		//// 创建私钥
		//var (
		//	address    string
		//	privateKey string
		//)
		//address, privateKey, err = generateKey()
		//if 0 >= len(address) || 0 >= len(privateKey) || err != nil {
		//	return nil, errors.New(500, "USER_ERROR", "生成地址错误")
		//}
		//
		//u.PrivateKey = privateKey
		//u.AddressTwo = address
		//
		//var (
		//	addressThree    string
		//	privateKeyThree string
		//)
		//u.WordThree = generateWord() // 生成助剂词
		//if 20 >= len(u.WordThree) {
		//	return nil, errors.New(500, "USER_ERROR", "生成助记词错误")
		//}
		//
		//privateKeyThree, addressThree = generateKeyBiw(u.WordThree)
		//if 0 >= len(addressThree) || 0 >= len(privateKeyThree) || err != nil {
		//	return nil, errors.New(500, "USER_ERROR", "生成地址错误")
		//}
		//
		//u.AddressThree = addressThree
		//u.PrivateKeyThree = privateKeyThree

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = uuc.repo.CreateUser(ctx, u) // 用户创建
			if err != nil {
				return err
			}

			_, err = uuc.uiRepo.CreateUserInfo(ctx, user) // 创建用户信息
			if err != nil {
				return err
			}

			_, err = uuc.urRepo.CreateUserRecommend(ctx, user, recommendUser) // 创建用户推荐信息
			if err != nil {
				return err
			}

			_, err = uuc.ubRepo.CreateUserBalance(ctx, user) // 创建余额信息
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func generateWord() string {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 定义总组数和每组的字符数
	numGroups := 12
	groupSize := 3

	// 生成随机字符串
	var result []string
	for i := 0; i < numGroups; i++ {
		result = append(result, randString(groupSize))
	}

	// 将字符串数组用逗号连接
	finalString := strings.Join(result, ",")
	return finalString
}

func randString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (uuc *UserUseCase) UpdateUserRecommend(ctx context.Context, u *User, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	var (
		err                   error
		userId                int64
		recommendUser         *UserRecommend
		userRecommend         *UserRecommend
		locations             []*LocationNew
		myRecommendUser       *User
		myUserRecommendUserId int64
		Address               string
		decodeBytes           []byte
	)

	code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
	if "abf00dd52c08a9213f225827bc3fb100" != code {
		decodeBytes, err = base64.StdEncoding.DecodeString(code)
		code = string(decodeBytes)
		if 1 >= len(code) {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码4")
		}
		if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码5")
		}

		// 现有推荐人信息，判断推荐人是否改变
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, u.ID)
		if nil == userRecommend {
			return nil, err
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
			if nil != err {
				return nil, err
			}
		}

		if nil == myRecommendUser {
			return &v1.RecommendUpdateReply{InviteUserAddress: ""}, nil
		}

		if myRecommendUser.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		if u.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 我的占位信息
		locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, u.ID)
		if nil != err {
			return nil, err
		}
		if nil != locations && 0 < len(locations) {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 查询推荐人的相关信息
		recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码6")
		}

		// 推荐人信息
		myRecommendUser, err = uuc.repo.GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		// 更新
		_, err = uuc.urRepo.UpdateUserRecommend(ctx, u, recommendUser)
		if err != nil {
			return nil, err
		}
		Address = myRecommendUser.Address
	}

	return &v1.RecommendUpdateReply{InviteUserAddress: Address}, err
}

func (uuc *UserUseCase) UserInfo(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	var (
		err                   error
		myUser                *User
		userRecommend         *UserRecommend
		myCode                string
		myUserRecommendUserId int64
		inviteUserAddress     string
		myRecommendUser       *User
		configs               []*Config
		userBalance           *UserBalance
		bPrice                float64
		exchangeRate          float64
		withdrawMin           float64
		withdrawMax           float64
		users                 []*User
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"b_price",
		"exchange_rate",
		"withdraw_amount_max",
		"withdraw_amount_min",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_amount_min" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "withdraw_amount_max" == vConfig.KeyName {
				withdrawMax, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "exchange_rate" == vConfig.KeyName {
				exchangeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}
	}

	users, err = uuc.repo.GetAllUsers(ctx)
	if nil != err {
		return nil, err
	}

	usersMap := make(map[int64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if 1 == myUser.IsDelete {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "用户已删除")
	}

	// 余额，收益总数
	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}

	// 推荐
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	if nil == userRecommend {
		return nil, err
	}

	myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}
		myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
		if nil != err {
			return nil, err
		}
		inviteUserAddress = myRecommendUser.Address
		myCode = userRecommend.RecommendCode + myCode
	}

	// 提现
	var (
		withdraws    []*Withdraw
		withdrawList []*v1.UserInfoReply_ListWithdraw
	)

	withdraws, err = uuc.ubRepo.GetWithdrawByUserId2(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	withdrawList = make([]*v1.UserInfoReply_ListWithdraw, 0)
	for _, v := range withdraws {
		if "RAW" == v.Type {
			withdrawList = append(withdrawList, &v1.UserInfoReply_ListWithdraw{
				Address:  v.Address,
				Amount:   v.AmountNew,
				CreateAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		}
	}

	startDate := time.Now().AddDate(0, 0, -1)

	// 分红
	var (
		userRewards []*Reward
	)

	listReward := make([]*v1.UserInfoReply_ListReward, 0)
	listExchange := make([]*v1.UserInfoReply_ListExchange, 0)
	userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	if nil != userRewards {
		for _, vUserReward := range userRewards {
			if "location" == vUserReward.Reason {
				//if vUserReward.CreatedAt.After(startDate) {
				//	totalYesReward += vUserReward.Amount
				//}
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Amount:     vUserReward.AmountNew,
					RewardType: 1,
				})
			} else if "recommend" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Amount:     vUserReward.AmountNew,
					RewardType: 2,
				})
			} else if "area" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Amount:     vUserReward.AmountNew,
					RewardType: 3,
				})
			} else if "area_three" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Amount:     vUserReward.AmountNew,
					RewardType: 4,
				})
			} else if "area_two" == vUserReward.Reason {
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Amount:     vUserReward.AmountNew,
					RewardType: 5,
				})
			} else if "exchange" == vUserReward.Reason {
				listExchange = append(listExchange, &v1.UserInfoReply_ListExchange{
					AmountRaw:  vUserReward.AmountNew,
					AmountRsdt: vUserReward.AmountNewTwo,
					CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
				})
			} else if "withdraw" == vUserReward.Reason {

			} else if "buy" == vUserReward.Reason {

			} else {
				continue
			}
		}
	}

	// 充值
	var (
		userEth []*EthUserRecord
	)
	userEth, err = uuc.repo.GetEthUserRecordListByUserId(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}
	listUserEth := make([]*v1.UserInfoReply_ListEthRecord, 0)
	for _, vUserEth := range userEth {
		coinType := "RAW"
		listUserEth = append(listUserEth, &v1.UserInfoReply_ListEthRecord{
			Amount:    uint64(vUserEth.AmountTwo),
			CoinType:  coinType,
			CreatedAt: vUserEth.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	// 全球
	var (
		day                    = -1
		userLocationsYes       []*LocationNew
		rewardLocationYes      int64
		totalRewardYes         int64
		fourUserRecommendTotal map[int64]int64
	)

	fourUserRecommendTotal = make(map[int64]int64, 0)
	userLocationsYes, err = uuc.locationRepo.GetLocationDailyYesterday(ctx, day)
	for _, userLocationYes := range userLocationsYes {
		rewardLocationYes += userLocationYes.Usdt

		// 获取直推

		var (
			fourUserRecommend         *UserRecommend
			myFourUserRecommendUserId int64
			//myFourRecommendUser *User
		)
		fourUserRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userLocationYes.UserId)
		if nil == fourUserRecommend {
			continue
		}

		if "" != fourUserRecommend.RecommendCode {
			tmpFourRecommendUserIds := strings.Split(fourUserRecommend.RecommendCode, "D")
			if 2 <= len(tmpFourRecommendUserIds) {
				myFourUserRecommendUserId, _ = strconv.ParseInt(tmpFourRecommendUserIds[len(tmpFourRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			//myFourRecommendUser, err = uuc.repo.GetUserById(ctx, myFourUserRecommendUserId)
			//if nil != err {
			//	return nil, err
			//}

			if _, ok := fourUserRecommendTotal[myFourUserRecommendUserId]; ok {
				fourUserRecommendTotal[myFourUserRecommendUserId] += userLocationYes.Usdt
			} else {
				fourUserRecommendTotal[myFourUserRecommendUserId] = userLocationYes.Usdt
			}
		}
	}

	// 前四名
	type KeyValuePair struct {
		Key   int64
		Value int64
	}
	var keyValuePairs []KeyValuePair
	for key, value := range fourUserRecommendTotal {
		keyValuePairs = append(keyValuePairs, KeyValuePair{key, value})
	}

	// 按值排序切片
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// 全网前天
	var (
		rewardFourYes *Reward
		yesAmount     int64
	)
	totalRewardYes = rewardLocationYes / 100 * total
	rewardFourYes, err = uuc.ubRepo.GetRewardFourYes(ctx) // 推荐人奖励
	if nil == err && nil != rewardFourYes {
		totalRewardYes += rewardFourYes.Amount
		yesAmount = rewardFourYes.Amount
	}
	totalReward := totalRewardYes / 100 * 70

	fourList := make([]*v1.UserInfoReply_ListFour, 0)

	// 获取前四项
	var topFour []KeyValuePair
	if 4 <= len(keyValuePairs) {
		topFour = keyValuePairs[:4]
	} else {
		topFour = keyValuePairs[:len(keyValuePairs)]
	}
	for k, vTopFour := range topFour {
		var (
			fourUser *User
		)
		fourUser, err = uuc.repo.GetUserById(ctx, vTopFour.Key)
		if nil != err {
			return nil, err
		}

		if nil == fourUser {
			continue
		}

		var (
			tmpMyRecommendAmount int64
		)
		if 0 == k {
			tmpMyRecommendAmount = totalReward / 100 * one
		} else if 1 == k {
			tmpMyRecommendAmount = totalReward / 100 * two
		} else if 2 == k {
			tmpMyRecommendAmount = totalReward / 100 * three
		} else if 3 == k {
			tmpMyRecommendAmount = totalReward / 100 * four
		}

		var address1 string
		if 20 <= len(fourUser.Address) {
			address1 = fourUser.Address[:6] + "..." + fourUser.Address[len(fourUser.Address)-4:]
		}
		fourList = append(fourList, &v1.UserInfoReply_ListFour{
			Location: address1,
			Amount:   fmt.Sprintf("%.2f", float64(vTopFour.Value)/float64(100000)),
			Reward:   fmt.Sprintf("%.2f", float64(tmpMyRecommendAmount)/float64(100000)),
		})
	}

	num := float64(1)
	if 1 == myUser.Last {
		num = 2
	} else if 2 == myUser.Last {
		num = 2.3
	} else if 3 == myUser.Last {
		num = 2.6
	} else if 4 == myUser.Last {
		num = 3
	} else if 5 == myUser.Last {
		num = 1.5
	} else if 6 == myUser.Last {
		num = 1.8
	} else {
		return nil, err
	}

	return &v1.UserInfoReply{
		WithdrawMin:       withdrawMin,
		WithdrawMax:       withdrawMax,
		One:               0,
		Two:               0,
		Three:             0,
		TodayTotal:        0,
		ListEthTotal:      nil,
		BPrice:            bPrice,
		ExchangeRate:      exchangeRate,
		BalanceRaw:        userBalance.BalanceRawFloat,
		BalanceUsdt:       userBalance.BalanceUsdtFloat,
		ListEth:           nil,
		ListBuy:           nil,
		ListWithdraw:      withdrawList,
		ListExchange:      nil,
		ListTo:            nil,
		ListStake:         nil,
		ListReward:        nil,
		Level:             0,
		MyAddress:         myUser.Address,
		InviteUserAddress: inviteUserAddress,
		AmoutUsdtGet:      myUser.AmountUsdtGet,
		AmoutUsdtSubGet:   myUser.AmountUsdt*num - myUser.AmountUsdtGet,
		AmoutUsdt:         myUser.AmountUsdt,
		RewardOne:         0,
		RewardTwo:         0,
		RewardThree:       0,
		RewardFour:        0,
		RewardFive:        0,
		ListOut:           nil,
	}, nil
}

func (uuc *UserUseCase) UserRecommend(ctx context.Context, req *v1.RecommendListRequest) (*v1.RecommendListReply, error) {

	var (
		userRecommend   *UserRecommend
		myUserRecommend []*UserRecommend
		recommendTotal  uint64
		user            *User
		err             error
	)

	user, _ = uuc.repo.GetUserByAddress(ctx, req.Address)
	if nil == user {
		return nil, err
	}

	// 推荐
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return nil, err
	}

	myUserRecommend, err = uuc.urRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(user.ID, 10))
	if nil == myUserRecommend || nil != err {
		return nil, err
	}

	var (
		totalNum int64
	)
	totalNum, err = uuc.urRepo.GetUserRecommendLikeCodeSum(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(user.ID, 10))
	if nil != err {
		return nil, err
	}

	tmpUserIds := make([]int64, 0)
	for _, vMyUserRecommend := range myUserRecommend {
		tmpUserIds = append(tmpUserIds, vMyUserRecommend.UserId)
	}
	if 0 >= len(tmpUserIds) {
		return nil, err
	}

	var (
		usersMap map[int64]*User
	)

	usersMap, err = uuc.repo.GetUserByUserIdsTwo(ctx, tmpUserIds)
	if nil == usersMap || nil != err {
		return nil, err
	}

	if 0 >= len(usersMap) {
		return nil, err
	}

	res := make([]*v1.RecommendListReply_List, 0)
	for _, vMyUserRecommend := range myUserRecommend {
		if _, ok := usersMap[vMyUserRecommend.UserId]; !ok {
			continue
		}

		recommendTotal++
		res = append(res, &v1.RecommendListReply_List{
			Address:   usersMap[vMyUserRecommend.UserId].Address,
			Amount:    usersMap[vMyUserRecommend.UserId].MyTotalAmount,
			CreatedAt: vMyUserRecommend.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.RecommendListReply{
		TotalNum:     recommendTotal,
		TotalTeamNum: uint64(totalNum),
		TotalAmount:  user.MyTotalAmount,
		Recommends:   res,
	}, nil
}

func (uuc *UserUseCase) UserArea(ctx context.Context, req *v1.UserAreaRequest, user *User) (*v1.UserAreaReply, error) {
	var (
		err             error
		locationId      = req.LocationId
		Locations       []*LocationNew
		LocationRunning *LocationNew
	)

	res := make([]*v1.UserAreaReply_List, 0)
	if 0 >= locationId {
		Locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, user.ID)
		if nil != err {
			return nil, err
		}
		for _, vLocations := range Locations {
			if "running" == vLocations.Status {
				LocationRunning = vLocations
				break
			}
		}

		if nil == LocationRunning {
			return &v1.UserAreaReply{Area: res}, nil
		}

		locationId = LocationRunning.ID
	}

	var (
		myLowLocations []*LocationNew
	)

	myLowLocations, err = uuc.locationRepo.GetLocationsByTop(ctx, locationId)
	if nil != err {
		return nil, err
	}

	for _, vMyLowLocations := range myLowLocations {
		var (
			userLow           *User
			tmpMyLowLocations []*LocationNew
		)

		userLow, err = uuc.repo.GetUserById(ctx, vMyLowLocations.UserId)
		if nil != err {
			continue
		}

		tmpMyLowLocations, err = uuc.locationRepo.GetLocationsByTop(ctx, vMyLowLocations.ID)
		if nil != err {
			return nil, err
		}

		var address1 string
		if 20 <= len(userLow.Address) {
			address1 = userLow.Address[:6] + "..." + userLow.Address[len(userLow.Address)-4:]
		}

		res = append(res, &v1.UserAreaReply_List{
			Address:    address1,
			LocationId: vMyLowLocations.ID,
			CountLow:   int64(len(tmpMyLowLocations)),
		})
	}

	return &v1.UserAreaReply{Area: res}, nil
}

func (uuc *UserUseCase) UserInfoOld(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	//var (
	//	myUser                  *User
	//	userInfo                *UserInfo
	//	locations               []*LocationNew
	//	myLastStopLocations     []*LocationNew
	//	userBalance             *UserBalance
	//	userRecommend           *UserRecommend
	//	userRecommends          []*UserRecommend
	//	userRewards             []*Reward
	//	userRewardTotal         int64
	//	userRewardWithdrawTotal int64
	//	encodeString            string
	//	recommendTeamNum        int64
	//	myCode                  string
	//	amount                  = "0"
	//	amount2                 = "0"
	//	configs                 []*Config
	//	myLastLocationCurrent   int64
	//	withdrawAmount          int64
	//	withdrawAmount2         int64
	//	myUserRecommendUserId   int64
	//	inviteUserAddress       string
	//	myRecommendUser         *User
	//	withdrawRate            int64
	//	myLocations             []*v1.UserInfoReply_List
	//	myLocations2            []*v1.UserInfoReply_List22
	//	allRewardList           []*v1.UserInfoReply_List9
	//	timeAgain               int64
	//	err                     error
	//)
	//
	//// 配置
	//configs, err = uuc.configRepo.GetConfigByKeys(ctx,
	//	"time_again",
	//)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "time_again" == vConfig.KeyName {
	//			timeAgain, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}
	//
	//myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, myUser.ID)
	//if nil != locations && 0 < len(locations) {
	//	for _, v := range locations {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//		if "running" == v.Status {
	//			amount = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations = append(myLocations, &v1.UserInfoReply_List{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			Num:       v.Num,
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//// 冻结
	//myLastStopLocations, err = uuc.locationRepo.GetMyStopLocationsLast(ctx, myUser.ID)
	//now := time.Now().UTC()
	//tmpNow := now.Add(8 * time.Hour)
	//if nil != myLastStopLocations {
	//	for _, vMyLastStopLocations := range myLastStopLocations {
	//		if tmpNow.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
	//			myLastLocationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax
	//		}
	//	}
	//}
	//
	//var (
	//	locations2 []*LocationNew
	//)
	//locations2, err = uuc.locationRepo.GetLocationsByUserId2(ctx, myUser.ID)
	//if nil != locations2 && 0 < len(locations2) {
	//	for _, v := range locations2 {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//
	//		if "running" == v.Status {
	//			amount2 = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations2 = append(myLocations2, &v1.UserInfoReply_List22{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//
	//userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	//if nil == userRecommend {
	//	return nil, err
	//}
	//
	//myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	//codeByte := []byte(myCode)
	//encodeString = base64.StdEncoding.EncodeToString(codeByte)
	//if "" != userRecommend.RecommendCode {
	//	tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
	//	if 2 <= len(tmpRecommendUserIds) {
	//		myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
	//	}
	//	myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
	//	if nil != err {
	//		return nil, err
	//	}
	//	inviteUserAddress = myRecommendUser.Address
	//	myCode = userRecommend.RecommendCode + myCode
	//}
	//
	//// 团队
	//var (
	//	teamUserIds        []int64
	//	teamUsers          map[int64]*User
	//	teamUserInfos      map[int64]*UserInfo
	//	teamUserAddresses  []*v1.UserInfoReply_List7
	//	recommendAddresses []*v1.UserInfoReply_List11
	//	teamLocations      map[int64][]*Location
	//	recommendUserIds   map[int64]int64
	//	userBalanceMap     map[int64]*UserBalance
	//)
	//recommendUserIds = make(map[int64]int64, 0)
	//userRecommends, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode)
	//if nil != userRecommends {
	//	for _, vUserRecommends := range userRecommends {
	//		if myCode == vUserRecommends.RecommendCode {
	//			recommendUserIds[vUserRecommends.UserId] = vUserRecommends.UserId
	//		}
	//		teamUserIds = append(teamUserIds, vUserRecommends.UserId)
	//	}
	//
	//	// 用户信息
	//	recommendTeamNum = int64(len(userRecommends))
	//	teamUsers, _ = uuc.repo.GetUserByUserIds(ctx, teamUserIds...)
	//	teamUserInfos, _ = uuc.uiRepo.GetUserInfoByUserIds(ctx, teamUserIds...)
	//	teamLocations, _ = uuc.locationRepo.GetLocationMapByIds(ctx, teamUserIds...)
	//	userBalanceMap, _ = uuc.ubRepo.GetUserBalanceByUserIds(ctx, teamUserIds...)
	//	if nil != teamUsers {
	//		for _, vTeamUsers := range teamUsers {
	//			var locationAmount int64
	//			if _, ok := teamUserInfos[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := userBalanceMap[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := teamLocations[vTeamUsers.ID]; ok {
	//
	//				for _, vTeamLocations := range teamLocations[vTeamUsers.ID] {
	//					locationAmount += vTeamLocations.Usdt
	//				}
	//			}
	//			if _, ok := recommendUserIds[vTeamUsers.ID]; ok {
	//				recommendAddresses = append(recommendAddresses, &v1.UserInfoReply_List11{
	//					Address: vTeamUsers.Address,
	//					Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//					Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//				})
	//			}
	//
	//			teamUserAddresses = append(teamUserAddresses, &v1.UserInfoReply_List7{
	//				Address: vTeamUsers.Address,
	//				Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//				Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//			})
	//		}
	//	}
	//}
	//
	//var (
	//	totalDeposit      int64
	//	userBalanceRecord []*UserBalanceRecord
	//	depositList       []*v1.UserInfoReply_List13
	//)
	//
	//userBalanceRecord, _ = uuc.ubRepo.GetUserBalanceRecordsByUserId(ctx, myUser.ID)
	//for _, vUserBalanceRecord := range userBalanceRecord {
	//	totalDeposit += vUserBalanceRecord.Amount
	//	depositList = append(depositList, &v1.UserInfoReply_List13{
	//		CreatedAt: vUserBalanceRecord.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//		Amount:    fmt.Sprintf("%.4f", float64(vUserBalanceRecord.Amount)/float64(100000)),
	//	})
	//}
	//
	//// 累计奖励
	//var (
	//	locationRewardList            []*v1.UserInfoReply_List2
	//	recommendRewardList           []*v1.UserInfoReply_List5
	//	locationTotal                 int64
	//	yesterdayLocationTotal        int64
	//	recommendRewardTotal          int64
	//	recommendRewardDhbTotal       int64
	//	yesterdayRecommendRewardTotal int64
	//)
	//
	//var startDate time.Time
	//var endDate time.Time
	//if 16 <= now.Hour() {
	//	startDate = now.AddDate(0, 0, -1)
	//	endDate = startDate.AddDate(0, 0, 1)
	//} else {
	//	startDate = now.AddDate(0, 0, -2)
	//	endDate = startDate.AddDate(0, 0, 1)
	//}
	//yesterdayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	//yesterdayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)
	//
	//fmt.Println(now, yesterdayStart, yesterdayEnd)
	//userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	//if nil != userRewards {
	//	for _, vUserReward := range userRewards {
	//		if "location" == vUserReward.Reason {
	//			locationTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayLocationTotal += vUserReward.Amount
	//			}
	//			locationRewardList = append(locationRewardList, &v1.UserInfoReply_List2{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "recommend" == vUserReward.Reason {
	//
	//			recommendRewardList = append(recommendRewardList, &v1.UserInfoReply_List5{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			recommendRewardTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayRecommendRewardTotal += vUserReward.Amount
	//			}
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "reward_withdraw" == vUserReward.Reason {
	//			userRewardTotal += vUserReward.Amount
	//			userRewardWithdrawTotal += vUserReward.Amount
	//		} else if "recommend_token" == vUserReward.Reason {
	//			recommendRewardDhbTotal += vUserReward.Amount
	//		}
	//	}
	//}
	//
	//var (
	//	withdraws []*Withdraw
	//)
	//withdraws, err = uuc.ubRepo.GetWithdrawByUserId2(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//for _, v := range withdraws {
	//	if "usdt" == v.Type {
	//		withdrawAmount += v.RelAmount
	//	}
	//	if "usdt_2" == v.Type {
	//		withdrawAmount2 += v.RelAmount
	//	}
	//}
	//
	//return &v1.UserInfoReply{
	//	Address:                 myUser.Address,
	//	Level:                   userInfo.Vip,
	//	Amount:                  amount,
	//	Amount2:                 amount2,
	//	LocationList2:           myLocations2,
	//	AmountB:                 fmt.Sprintf("%.2f", float64(myLastLocationCurrent)/float64(100000)),
	//	DepositList:             depositList,
	//	BalanceUsdt:             fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt)/float64(100000)),
	//	BalanceUsdt2:            fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt2)/float64(100000)),
	//	BalanceDhb:              fmt.Sprintf("%.2f", float64(userBalance.BalanceDhb)/float64(100000)),
	//	InviteUrl:               encodeString,
	//	RecommendNum:            userInfo.HistoryRecommend,
	//	RecommendTeamNum:        recommendTeamNum,
	//	Total:                   fmt.Sprintf("%.2f", float64(userRewardTotal)/float64(100000)),
	//	RewardWithdraw:          fmt.Sprintf("%.2f", float64(userRewardWithdrawTotal)/float64(100000)),
	//	WithdrawAmount:          fmt.Sprintf("%.2f", float64(withdrawAmount)/float64(100000)),
	//	WithdrawAmount2:         fmt.Sprintf("%.2f", float64(withdrawAmount2)/float64(100000)),
	//	LocationTotal:           fmt.Sprintf("%.2f", float64(locationTotal)/float64(100000)),
	//	Account:                 "0xAfC39F3592A1024144D1ba6DC256397F4DbfbE2f",
	//	LocationList:            myLocations,
	//	TeamAddressList:         teamUserAddresses,
	//	AllRewardList:           allRewardList,
	//	InviteUserAddress:       inviteUserAddress,
	//	RecommendAddressList:    recommendAddresses,
	//	WithdrawRate:            withdrawRate,
	//	RecommendRewardTotal:    fmt.Sprintf("%.2f", float64(recommendRewardTotal)/float64(100000)),
	//	RecommendRewardDhbTotal: fmt.Sprintf("%.2f", float64(recommendRewardDhbTotal)/float64(100000)),
	//	TotalDeposit:            fmt.Sprintf("%.2f", float64(totalDeposit)/float64(100000)),
	//	WithdrawAll:             fmt.Sprintf("%.2f", float64(withdrawAmount+withdrawAmount2)/float64(100000)),
	//}, nil
	return nil, nil

}

func (uuc *UserUseCase) RewardList(ctx context.Context, req *v1.RewardListRequest, user *User) (*v1.RewardListReply, error) {

	res := &v1.RewardListReply{
		Rewards: make([]*v1.RewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) RecommendRewardList(ctx context.Context, user *User) (*v1.RecommendRewardListReply, error) {

	res := &v1.RecommendRewardListReply{
		Rewards: make([]*v1.RecommendRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) FeeRewardList(ctx context.Context, user *User) (*v1.FeeRewardListReply, error) {
	res := &v1.FeeRewardListReply{
		Rewards: make([]*v1.FeeRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) TranList(ctx context.Context, user *User, reqTypeCoin string, reqTran string) (*v1.TranListReply, error) {

	var (
		userBalanceRecord []*UserBalanceRecord
		typeCoin          = "usdt"
		tran              = "tran"
		err               error
	)

	res := &v1.TranListReply{
		Tran: make([]*v1.TranListReply_List, 0),
	}

	if "" != reqTypeCoin {
		typeCoin = reqTypeCoin
	}

	if "tran_to" == reqTran {
		tran = reqTran
	}

	userBalanceRecord, err = uuc.ubRepo.GetUserBalanceRecordByUserId(ctx, user.ID, typeCoin, tran)
	if nil != err {
		return res, err
	}

	for _, v := range userBalanceRecord {
		res.Tran = append(res.Tran, &v1.TranListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) WithdrawList(ctx context.Context, user *User, reqTypeCoin string) (*v1.WithdrawListReply, error) {

	var (
		withdraws []*Withdraw
		typeCoin  = "usdt"
		err       error
	)

	res := &v1.WithdrawListReply{
		Withdraw: make([]*v1.WithdrawListReply_List, 0),
	}

	if "" != reqTypeCoin {
		typeCoin = reqTypeCoin
	}

	withdraws, err = uuc.ubRepo.GetWithdrawByUserId(ctx, user.ID, typeCoin)
	if nil != err {
		return res, err
	}

	for _, v := range withdraws {
		res.Withdraw = append(res.Withdraw, &v1.WithdrawListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
			Status:    v.Status,
			Type:      v.Type,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) TradeList(ctx context.Context, user *User) (*v1.TradeListReply, error) {

	var (
		trades []*Trade
		err    error
	)

	res := &v1.TradeListReply{
		Trade: make([]*v1.TradeListReply_List, 0),
	}

	trades, err = uuc.ubRepo.GetTradeByUserId(ctx, user.ID)
	if nil != err {
		return res, err
	}

	for _, v := range trades {
		res.Trade = append(res.Trade, &v1.TradeListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			AmountCsd: fmt.Sprintf("%.2f", float64(v.AmountCsd)/float64(100000)),
			AmountHbs: fmt.Sprintf("%.2f", float64(v.AmountHbs)/float64(100000)),
			Status:    v.Status,
		})
	}

	return res, nil
}

func lessThanOrEqualZero(a, b float64, epsilon float64) bool {
	return a-b < epsilon || math.Abs(a-b) < epsilon
}

func (uuc *UserUseCase) Buy(ctx context.Context, req *v1.BuyRequest, user *User) (*v1.BuyReply, error) {
	var (
		amount   = req.SendBody.Amount
		configs  []*Config
		bPrice   float64
		coinType string
		err      error
	)

	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"b_price",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "b_price" == vConfig.KeyName {
				bPrice, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return &v1.BuyReply{
						Status: "币价错误",
					}, nil
				}
			}
		}
	}

	var (
		tmpValue    int64
		strValue    string
		userBalance *UserBalance
	)

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return &v1.BuyReply{
			Status: "错误",
		}, nil
	}

	amountRaw := bPrice * float64(amount)
	if 1 == req.SendBody.Type {
		if amountRaw > userBalance.BalanceRawFloat {
			return &v1.BuyReply{
				Status: "raw余额不足",
			}, nil
		}

		coinType = "RAW"
	} else {
		return &v1.BuyReply{
			Status: "类型错误",
		}, nil
	}

	notExistDepositResult := make([]*EthUserRecord, 0)
	notExistDepositResult = append(notExistDepositResult, &EthUserRecord{ // 两种币的记录
		UserId:    user.ID,
		Status:    "success",
		Type:      "deposit",
		Amount:    strValue,
		RelAmount: tmpValue,
		CoinType:  coinType,
		Last:      0,
	})

	var (
		res bool
	)
	res, err = uuc.EthUserRecordHandle(ctx, amount, amountRaw, coinType, notExistDepositResult...)
	if !res || nil != err {
		fmt.Println(err)
		return &v1.BuyReply{
			Status: "认购错误",
		}, nil
	}

	return &v1.BuyReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) EthUserRecordHandle(ctx context.Context, amount uint64, amountRaw float64, coinType string, ethUserRecord ...*EthUserRecord) (bool, error) {

	var (
		err     error
		configs []*Config
		va4     float64
		va5     float64
		va6     float64
		va7     float64
		va8     float64
	)

	// 配置
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "va4", "va5", "va6", "va7", "va8")
	if nil != configs {
		for _, vConfig := range configs {
			if "va4" == vConfig.KeyName {
				va4, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "va5" == vConfig.KeyName {
				va5, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "va6" == vConfig.KeyName {
				va6, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "va7" == vConfig.KeyName {
				va7, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "va8" == vConfig.KeyName {
				va8, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}
	}

	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
	)
	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		fmt.Println("今认购用户获取失败2")
		return false, err
	}

	myLowUser := make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)
	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		users       []*User
		usersMap    map[int64]*User
		stopUserIds map[int64]bool
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil == users {
		fmt.Println("认购用户获取失败")
		return false, nil
	}

	usersMap = make(map[int64]*User, 0)

	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	for _, v := range ethUserRecord {
		if _, ok := usersMap[v.UserId]; !ok {
			fmt.Println("认购用户不存在", v)
			continue
		}

		last := uint64(0)
		newAmountUsdt := usersMap[v.UserId].AmountUsdt + float64(amount)
		if 300 <= newAmountUsdt && newAmountUsdt < 500 {
			last = 5
		} else if 500 <= newAmountUsdt && newAmountUsdt < 1000 {
			last = 6
		} else if 1000 <= newAmountUsdt && newAmountUsdt < 5000 {
			last = 1
		} else if 5000 <= newAmountUsdt && newAmountUsdt < 30000 {
			last = 2
		} else if 30000 <= newAmountUsdt && newAmountUsdt < 100000 {
			last = 3
		} else if 100000 <= newAmountUsdt {
			last = 4
		} else {
			fmt.Println("认购信息无效", v, usersMap[v.UserId])
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.repo.UpdateUserNewTwoNew(ctx, v.UserId, float64(amount), last, amountRaw, coinType)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "错误投资3", v)
			return false, err
		}

		// 推荐人
		var (
			userRecommend       *UserRecommend
			tmpRecommendUserIds []string
		)
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, v.UserId)
		if nil != err {
			continue
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		lastLevel := 0
		lastLevelNum := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			currentLevel := 0

			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红社区，信息缺失,user：", err, v, tmpUserId)
				continue
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				// 增加业绩
				err = uuc.repo.UpdateUserMyTotalAmount(ctx, tmpUserId, float64(amount))
				if err != nil {
					fmt.Println("错误分红静态：", err, v)
				}

				return nil
			}); nil != err {
				fmt.Println(err, "错误投资2", v)
				return false, err
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红社区，信息缺失,user1：", err, v)
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红社区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红社区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpMaxId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红社区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount
					tmpMaxId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			tmpAreaMin := float64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if tmpMaxId != vMyLowUser.UserId {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount
				}
			}

			tmpLastLevelNum := float64(0)
			if 100000 <= tmpAreaMin {
				currentLevel = 4
				tmpLastLevelNum = va4
			} else if 300000 <= tmpAreaMin {
				currentLevel = 5
				tmpLastLevelNum = va5
			} else if 1000000 <= tmpAreaMin {
				currentLevel = 6
				tmpLastLevelNum = va6
			} else if 3000000 <= tmpAreaMin {
				currentLevel = 7
				tmpLastLevelNum = va7
			} else if 10000000 <= tmpAreaMin {
				currentLevel = 8
				tmpLastLevelNum = va8
			} else {
				// 跳过，没级别
				continue
			}

			// 级别低跳过
			if currentLevel <= lastLevel {
				if 8 == lastLevel {
					break
				}

				continue
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红社区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmp := float64(amount) * (tmpLastLevelNum - lastLevelNum)

				var (
					stopArea2 bool
					num       float64
				)
				if 1 == tmpRecommendUser.Last {
					num = 2
				} else if 2 == tmpRecommendUser.Last {
					num = 2.3
				} else if 3 == tmpRecommendUser.Last {
					num = 2.6
				} else if 4 == tmpRecommendUser.Last {
					num = 3
				} else if 5 == tmpRecommendUser.Last {
					num = 1.5
				} else if 6 == tmpRecommendUser.Last {
					num = 1.8
				} else {
					continue
				}

				if !lessThanOrEqualZero(tmp+tmpRecommendUser.AmountUsdtGet, tmpRecommendUser.AmountUsdt*num, 1e-7) {
					tmp = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
					stopArea2 = true
				}

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					var (
						code int64
					)

					code, err = uuc.repo.UpdateUserRewardAreaTwo(ctx, tmpRecommendUser.ID, tmp, tmpRecommendUser.AmountUsdt, stopArea2)
					if code > 0 && err != nil {
						fmt.Println("错误分红社区：", err, tmpRecommendUser)
					}

					if stopArea2 {
						stopUserIds[tmpRecommendUser.ID] = true // 出局

						// 推荐人
						var (
							userRecommendArea *UserRecommend
						)
						if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
							userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
						} else {
							fmt.Println("错误分红社区，信息缺失7：", err, v)
						}

						if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
							var tmpRecommendAreaUserIds []string
							tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

							for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
								if 0 >= len(vTmpRecommendAreaUserIds) {
									continue
								}

								myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
								if 0 >= myUserRecommendAreaUserId {
									continue
								}

								// 减掉业绩
								err = uuc.repo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
								if err != nil {
									fmt.Println("错误分红社区：", err, v)
								}
							}
						}
					}

					return nil
				}); nil != err {
					fmt.Println("err reward area 2", err, v)
				}

				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}
		}
	}

	return true, nil
}

// Exchange Exchange.
func (uuc *UserUseCase) Exchange(ctx context.Context, req *v1.ExchangeRequest, user *User) (*v1.ExchangeReply, error) {
	var (
		//u           *User
		err         error
		userBalance *UserBalance
	)

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)

	if userBalance.BalanceUsdtFloat < amountFloat {
		amountFloat = userBalance.BalanceUsdtFloat
	}

	// 配置
	var (
		configs      []*Config
		exchangeRate float64
		bPrice       float64
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx, "exchange_rate", "b_price")
	if nil != configs {
		for _, vConfig := range configs {
			if "exchange_rate" == vConfig.KeyName {
				exchangeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseFloat(vConfig.Value, 10)
			}

		}
	}

	var (
		amountRaw       float64
		amountRawSubFee float64
	)

	amountRaw = amountFloat / bPrice
	fee := amountRaw * exchangeRate
	amountRawSubFee = amountRaw - fee
	if amountRawSubFee <= 0 {
		return &v1.ExchangeReply{
			Status: "fail price",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.Exchange(ctx, user.ID, amountFloat, fee, amountRawSubFee) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.ExchangeReply{
		Status: "ok",
	}, nil

}

func (uuc *UserUseCase) Withdraw(ctx context.Context, req *v1.WithdrawRequest, user *User, password string) (*v1.WithdrawReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)

	// 配置
	var (
		configs     []*Config
		withdrawMax float64
		withdrawMin float64
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"withdraw_amount_max", "withdraw_amount_min",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_amount_max" == vConfig.KeyName {
				withdrawMax, _ = strconv.ParseFloat(vConfig.Value, 10)
			}

			if "withdraw_amount_max" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}
	}

	if userBalance.BalanceRawFloat < amountFloat {
		amountFloat = userBalance.BalanceRawFloat
	}

	if withdrawMax < amountFloat {
		return &v1.WithdrawReply{
			Status: "fail max",
		}, nil
	}

	if withdrawMin > amountFloat {
		return &v1.WithdrawReply{
			Status: "fail min",
		}, nil
	}

	amountFloatSubFee := amountFloat
	if 0 >= amountFloat {
		return &v1.WithdrawReply{
			Status: "fail price",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.ubRepo.WithdrawUsdt2(ctx, user.ID, amountFloat) // 提现
		if nil != err {
			return err
		}
		_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amountFloatSubFee, amountFloat, "RAW", req.SendBody.Address)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Tran(ctx context.Context, req *v1.TranRequest, user *User, password string) (*v1.TranReply, error) {
	var (
		err         error
		userBalance *UserBalance
		toUser      *User
		u           *User
	)

	u, _ = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if "" == u.Password || 6 > len(u.Password) {
		return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	}

	if u.Password != user.Password {
		return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	}

	if password != u.Password {
		return nil, errors.New(500, "密码错误", "密码错误")
	}

	if "" == req.SendBody.Address {
		return &v1.TranReply{
			Status: "不存在地址",
		}, nil
	}

	toUser, err = uuc.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if nil != err {
		return &v1.TranReply{
			Status: "不存在地址",
		}, nil
	}

	if user.ID == toUser.ID {
		return &v1.TranReply{
			Status: "不能给自己转账",
		}, nil
	}

	if "dhb" != req.SendBody.Type && "usdt" != req.SendBody.Type {
		return &v1.TranReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	if "dhb" == req.SendBody.Type {
		if userBalance.BalanceDhb < amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}

		if 10000000 > amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}
	}

	if "usdt" == req.SendBody.Type {
		if userBalance.BalanceUsdt < amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}

		if 1000000 > amount {
			return &v1.TranReply{
				Status: "fail",
			}, nil
		}
	}

	var (
		userRecommend  *UserRecommend
		userRecommend2 *UserRecommend
	)
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return &v1.TranReply{
			Status: "信息错误",
		}, nil
	}

	var (
		tmpRecommendUserIds          []string
		tmpRecommendUserIdsInt       []int64
		toUserTmpRecommendUserIds    []string
		toUserTmpRecommendUserIdsInt []int64
	)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}

	if 1 < len(tmpRecommendUserIds) {
		lastKey := len(tmpRecommendUserIds) - 1
		if 1 <= lastKey {
			for i := 0; i <= lastKey; i++ {
				// 有占位信息，推荐人推荐人的上一代
				if lastKey-i <= 0 {
					break
				}

				tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
				tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
			}
		}
	}

	userRecommend2, err = uuc.urRepo.GetUserRecommendByUserId(ctx, toUser.ID)
	if nil == userRecommend2 {
		return &v1.TranReply{
			Status: "信息错误",
		}, nil
	}
	if "" != userRecommend2.RecommendCode {
		toUserTmpRecommendUserIds = strings.Split(userRecommend2.RecommendCode, "D")
	}

	if 1 < len(toUserTmpRecommendUserIds) {
		lastKey2 := len(toUserTmpRecommendUserIds) - 1
		if 1 <= lastKey2 {
			for i := 0; i <= lastKey2; i++ {
				// 有占位信息，推荐人推荐人的上一代
				if lastKey2-i <= 0 {
					break
				}

				toUserTmpMyTopUserRecommendUserId, _ := strconv.ParseInt(toUserTmpRecommendUserIds[lastKey2-i], 10, 64) // 最后一位是直推人
				toUserTmpRecommendUserIdsInt = append(toUserTmpRecommendUserIdsInt, toUserTmpMyTopUserRecommendUserId)
			}
		}
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		if "usdt" == req.SendBody.Type {
			err = uuc.ubRepo.TranUsdt(ctx, user.ID, toUser.ID, amount, tmpRecommendUserIdsInt, toUserTmpRecommendUserIdsInt) // 提现
			if nil != err {
				return err
			}
		} else if "dhb" == req.SendBody.Type {
			err = uuc.ubRepo.TranDhb(ctx, user.ID, toUser.ID, amount) // 提现
			if nil != err {
				return err
			}
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.TranReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Trade(ctx context.Context, req *v1.WithdrawRequest, user *User, amount int64, amountB int64, amount2 int64, password string) (*v1.WithdrawReply, error) {
	var (
		u                   *User
		userBalance         *UserBalance
		userBalance2        *UserBalance
		configs             []*Config
		userRecommend       *UserRecommend
		withdrawRate        int64
		withdrawDestroyRate int64
		err                 error
	)

	u, _ = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if "" == u.Password || 6 > len(u.Password) {
		return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	}

	if u.Password != user.Password {
		return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	}

	if password != u.Password {
		return nil, errors.New(500, "密码错误", "密码错误")
	}

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "withdraw_rate",
		"withdraw_destroy_rate",
	)

	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_destroy_rate" == vConfig.KeyName {
				withdrawDestroyRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	userBalance, err = uuc.ubRepo.GetUserBalanceLock(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	userBalance2, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.WithdrawReply{
			Status: "csd锁定部分的余额不足",
		}, nil
	}

	if userBalance2.BalanceDhb < amountB {
		return &v1.WithdrawReply{
			Status: "hbs锁定部分的余额不足",
		}, nil
	}

	// 推荐人
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	if nil == userRecommend {
		return &v1.WithdrawReply{
			Status: "信息错误",
		}, nil
	}

	var (
		tmpRecommendUserIds    []string
		tmpRecommendUserIdsInt []int64
	)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	}
	lastKey := len(tmpRecommendUserIds) - 1
	if 1 <= lastKey {
		for i := 0; i <= lastKey; i++ {
			// 有占位信息，推荐人推荐人的上一代
			if lastKey-i <= 0 {
				break
			}

			tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
			tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
		}
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.Trade(ctx, user.ID, amount, amountB, amount-amount/100*(withdrawRate+withdrawDestroyRate), amountB-amountB/100*(withdrawRate+withdrawDestroyRate), tmpRecommendUserIdsInt, amount2) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) SetBalanceReward(ctx context.Context, req *v1.SetBalanceRewardRequest, user *User) (*v1.SetBalanceRewardReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.SetBalanceReward(ctx, user.ID, amount) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.SetBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) DeleteBalanceReward(ctx context.Context, req *v1.DeleteBalanceRewardRequest, user *User) (*v1.DeleteBalanceRewardReply, error) {
	var (
		err            error
		balanceRewards []*BalanceReward
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	balanceRewards, err = uuc.ubRepo.GetBalanceRewardByUserId(ctx, user.ID)
	if nil != err {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	var totalBalanceRewardAmount int64
	for _, vBalanceReward := range balanceRewards {
		totalBalanceRewardAmount += vBalanceReward.Amount
	}

	if totalBalanceRewardAmount < amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	for _, vBalanceReward := range balanceRewards {
		tmpAmount := int64(0)
		Status := int64(1)

		if amount-vBalanceReward.Amount < 0 {
			tmpAmount = amount
		} else {
			tmpAmount = vBalanceReward.Amount
			Status = 2
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.UpdateBalanceReward(ctx, user.ID, vBalanceReward.ID, tmpAmount, Status) // 提现
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
		amount -= tmpAmount

		if amount <= 0 {
			break
		}
	}

	return &v1.DeleteBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	res := &v1.AdminRewardListReply{
		Rewards: make([]*v1.AdminRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {

	res := &v1.AdminUserListReply{
		Users: make([]*v1.AdminUserListReply_UserList, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error) {
	return uuc.repo.GetUserByUserIds(ctx, userIds...)
}

func (uuc *UserUseCase) GetUserByUserId(ctx context.Context, userId int64) (*User, error) {
	return uuc.repo.GetUserById(ctx, userId)
}

func (uuc *UserUseCase) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}
	return res, nil

}

func (uuc *UserUseCase) AdminRecommendList(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	res := &v1.AdminUserRecommendReply{
		Users: make([]*v1.AdminUserRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {

	res := &v1.AdminMonthRecommendReply{
		Users: make([]*v1.AdminMonthRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	res := &v1.AdminConfigReply{
		Config: make([]*v1.AdminConfigReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	res := &v1.AdminConfigUpdateReply{}
	return res, nil
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedList(ctx context.Context) ([]*Withdraw, error) {
	return uuc.ubRepo.GetWithdrawPassOrRewarded(ctx)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	res := &v1.AdminWithdrawListReply{
		Withdraw: make([]*v1.AdminWithdrawListReply_List, 0),
	}

	return res, nil

}

func (uuc *UserUseCase) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return &v1.AdminFeeReply{}, nil
}

func (uuc *UserUseCase) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {

	return &v1.AdminAllReply{}, nil
}

func (uuc *UserUseCase) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return &v1.AdminWithdrawReply{}, nil
}
