type Stock struct {
	span  int
	price int
}
type StockSpanner struct {
	stocks []Stock
}

func Constructor() StockSpanner {
	stocks := make([]Stock, 0)
	return StockSpanner{stocks}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stocks) > 0 && this.stocks[len(this.stocks)-1].price <= price {
		topSpan := this.stocks[len(this.stocks)-1].span
		this.stocks = this.stocks[:len(this.stocks)-1]
		span += topSpan
	}
	this.stocks = append(this.stocks, Stock{span, price})
	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */