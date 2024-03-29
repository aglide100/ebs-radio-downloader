package word

func EditDistance(word1, word2 string) int {
    len1, len2 := len(word1), len(word2)
    dp := make([][]int, len1+1)

    for i := 0; i <= len1; i++ {
        dp[i] = make([]int, len2+1)
        dp[i][0] = i
    }

    for j := 0; j <= len2; j++ {
        dp[0][j] = j
    }

    for i := 1; i <= len1; i++ {
        for j := 1; j <= len2; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
            }
        }
    }

    return dp[len1][len2]
}

func min(nums ...int) int {
    min := nums[0]

    for _, num := range nums {
        if num < min {
            min = num
        }
    }
    
    return min
}
