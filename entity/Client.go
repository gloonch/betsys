package entity

type Client struct {
	ID      uint
	Balance int
	Bets    []Bet
}

type Bet struct {
	ID               uint
	Amount           int
	BalanceBeforeBet int
	BalanceAfterBet  int
}
