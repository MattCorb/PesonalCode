def maxProfit(self, prices: list[int]) -> int:
        
        max_profit = 0
        
        if len(prices) >= 2:

                left = prices[0]

                for n in range(1, len(prices)):
                    
                    right = prices[n]
                    
                    if right - left > max_profit:
                        max_profit = right - left

                    elif right < left:
                        left = right

                    


        return max_profit

'''
Solution:
Use the two pointer method to accomplish an O(n) time complexity
left (finger) follows the dips in the graph
right (finger) follows the peaks in the graph

each iteration the right variable is reassigned to the next element in the list
left and right are compared, if the difference between the two is greater than the previous
max the new max is stored.

'''