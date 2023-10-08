package array_strings

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 {
		if n == 1 && flowerbed[0] == 0 {
			return true
		}
		return false
	}

	for i := 0; i < len(flowerbed); i++ {
		if i == 0 && flowerbed[i] == 0 {
			if flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
			continue
		}

		if i == len(flowerbed)-1 && flowerbed[i] == 0 {
			if flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
			continue
		}

		if flowerbed[i] == 1 {
			continue
		}

		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}

	if n <= 0 {
		return true
	}
	return false
}
