import math
class Solution:
    def maxSumOfSquares(self, num: int, sum: int) -> str:
        i, j = 0, sum
        pi = []
        nums = []
        while i < j:
            if (i + j) == sum:
                si = str(i)
                sj = str(j)
                sij = sj + si
                if len(sij) == num:                
                    pi.append((i ** 2) + (j ** 2))
                    nums.append(sij)
            i += 1
            j -= 1
        maxVI = pi.index(max(pi))
        return nums[maxVI]
    
s = Solution()
s.maxSumOfSquares(num=2, sum=18)