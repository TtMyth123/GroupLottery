package mconst

var (
	mpAccountName map[Account_Type]string
	mpRebateName  map[Rebate_Type]string
)

func init() {
	mpAccountName = make(map[Account_Type]string)
	mpAccountName[Account_01_Guess] = Account_N_01_Guess
	mpAccountName[Account_02_Win] = Account_N_02_Win
	//mpAccountName[Account_10_SubGuessRebate] = Account_N_10_SubGuessRebate
	//mpAccountName[Account_11_SubGuessRebate] = Account_N_11_SubGuessRebate
	mpAccountName[Account_03_SaveMoney] = Account_N_03_SaveMoney
	mpAccountName[Account_03_SaveMoney2] = Account_N_03_SaveMoney2
	mpAccountName[Account_04_DrawMoney] = Account_N_04_DrawMoney
	mpAccountName[Account_04_DrawMoney2] = Account_N_04_DrawMoney2
	mpAccountName[Account_07_DrawMoneyR] = Account_N_07_DrawMoneyR
	mpAccountName[Account_07_DrawMoneyR2] = Account_N_07_DrawMoneyR2
	mpAccountName[Account_08_AddMoney] = Account_N_08_AddMoney
	mpAccountName[Account_08_AddMoney2] = Account_N_08_AddMoney2
	mpAccountName[Account_09_DecMoney] = Account_N_09_DecMoney
	mpAccountName[Account_09_DecMoney2] = Account_N_09_DecMoney2
	mpAccountName[Account_05_Give] = Account_N_05_Give
	mpAccountName[Account_05_Give2] = Account_N_05_Give2
	//mpAccountName[Account_12_XmGuess       ] = Account_N_12_XmGuess
	mpAccountName[Account_13_Rebate] = Account_N_13_Rebate
	mpAccountName[Account_13_Rebate2] = Account_N_13_Rebate2

	mpRebateName = make(map[Rebate_Type]string)
	mpRebateName[Rebate_01_Guess] = Rebate_N_01_Guess
	mpRebateName[Rebate_02_ToGold] = Rebate_N_02_ToGold

}
func GetAccountName(a Account_Type) string {
	return mpAccountName[a]
}

func GetRebateName(a Rebate_Type) string {
	return mpRebateName[a]
}
