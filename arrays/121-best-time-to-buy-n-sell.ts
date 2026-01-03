/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock
Problem: Best Time to Buy and Sell Stock

Approach:
- Track minimum price so far as buy price
- For each day, calculate profit if sold today
- Keep updating maximum profit

Time Complexity: O(n)
Space Complexity: O(1)
*/

function maxProfit(prices: number[]): number {
  let minBuy = prices[0];
  let profit = 0;

  for (let i = 1; i < prices.length; i++) {
    profit = Math.max(profit, prices[i] - minBuy);
    minBuy = Math.min(minBuy, prices[i]);
  }

  return profit;
}
