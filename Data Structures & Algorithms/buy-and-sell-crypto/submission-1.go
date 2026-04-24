func maxProfit(prices []int) int {
    // wrong intuition: searching for max and min wont solve because order is
    //   important as this is by date. min's index could be later after max's
    // naive approach: O(n^2), set one at the beginning, iterate over the rest
    //   to found the value that can gain profit (value should be larger)
    // out of the box: we can time travel! this way, we track minimum price as
    //   we iterate prices, and check in each iteration for maxProfit. i'm not
    //   thinking about this since i assume we can't go back in time to buy :)
    
    minPrice := prices[0]
    maxProfit := 0
    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        } else if price - minPrice > maxProfit {
            maxProfit = price - minPrice
        }
    }
    return maxProfit
}
