package handlers

import (
	. "main/backend/models"
	"strconv"
	"time"
)

func linearRegression(x, y []float64) (float64, float64) {
	var sumX, sumY, sumXY, sumXX float64

	// 计算各种和值
	for i := 0; i < len(x); i++ {
		if len(y) > i {
			sumX += x[i]
			sumY += y[i]
			sumXY += x[i] * y[i]
			sumXX += x[i] * x[i]
		}

	}

	// 计算斜率和截距
	n := float64(len(x))
	slope := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	return slope, intercept
}
func IncomeEstWeek() ([]float64, error) {
	var orders []Order
	var pridic []float64
	err := Db.Find(&orders)
	if err != nil {
		return []float64{}, err
	}
	if len(orders) != 0 {
		// 获取今天的日期
		today := time.Now()

		// 初始化七天价格总和数组
		sevenDayTotal := make([]float64, 6)

		// 遍历订单，计算七天价格总和
		for i := 0; i < len(orders); i++ {
			orderTime := orders[i].Time
			daysDiff := int(today.Sub(orderTime).Hours() / 24)

			if daysDiff >= 0 && daysDiff < 7 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[0] += price
			} else if daysDiff >= 7 && daysDiff < 14 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[1] += price
			} else if daysDiff >= 14 && daysDiff < 21 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[2] += price
			} else if daysDiff >= 21 && daysDiff < 28 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[3] += price
			} else if daysDiff >= 28 && daysDiff < 35 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[4] += price
			} else if daysDiff >= 35 && daysDiff < 42 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[5] += price
			}
		}
		if len(sevenDayTotal) != 0 {
			pridic = append(pridic, sevenDayTotal[0])
		}

		// 构建 x 轴数据（1, 2, 3）
		x := make([]float64, len(orders))
		for i := 0; i < len(x); i++ {
			x[i] = float64(i + 1)
		}
		// 执行线性回归
		slope, intercept := linearRegression(x, sevenDayTotal)
		// 输出线性回归方程

		// 计算后三周的预测数据
		predictions := make([]float64, 3)
		for i := 0; i < 3; i++ {
			xVal := float64(len(sevenDayTotal) + i + 1)
			predictions[i] = slope*xVal + intercept
		}
		for i := 0; i < len(predictions); i++ {
			pridic = append(pridic, predictions[i])
		}
	}
	return pridic, err
}
func IncomeEstMonth() ([]float64, error) {
	var orders []Order
	var pridic []float64
	err := Db.Find(&orders)
	if err != nil {
		return []float64{}, err
	}
	if len(orders) != 0 {
		// 获取今天的日期
		today := time.Now()

		// 初始化七天价格总和数组
		sevenDayTotal := make([]float64, 6)

		// 遍历订单，计算七天价格总和
		for i := 0; i < len(orders); i++ {
			orderTime := orders[i].Time
			daysDiff := int(today.Sub(orderTime).Hours() / 24)

			if daysDiff >= 0 && daysDiff < 30 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[0] += price
			} else if daysDiff >= 30 && daysDiff < 60 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[1] += price
			} else if daysDiff >= 60 && daysDiff < 90 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[2] += price
			} else if daysDiff >= 90 && daysDiff < 120 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[3] += price
			} else if daysDiff >= 120 && daysDiff < 150 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[4] += price
			} else if daysDiff >= 150 && daysDiff < 180 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[5] += price
			}
		}
		if len(sevenDayTotal) != 0 {
			pridic = append(pridic, sevenDayTotal[0])
		}
		// 构建 x 轴数据（1, 2, 3）
		x := make([]float64, len(orders))
		for i := 0; i < len(x); i++ {
			x[i] = float64(i + 1)
		}
		// 执行线性回归
		slope, intercept := linearRegression(x, sevenDayTotal)
		// 输出线性回归方程

		// 计算后三周的预测数据
		predictions := make([]float64, 3)
		for i := 0; i < 3; i++ {
			xVal := float64(len(sevenDayTotal) + i + 1)
			predictions[i] = slope*xVal + intercept
		}
		for i := 0; i < len(predictions); i++ {
			pridic = append(pridic, predictions[i])
		}
	}
	return pridic, err
}
func IncomeEstYear() ([]float64, error) {
	var orders []Order
	var pridic []float64
	err := Db.Find(&orders)
	if err != nil {
		return []float64{}, err
	}
	if len(orders) != 0 {
		// 获取今天的日期
		today := time.Now()

		// 初始化七天价格总和数组
		sevenDayTotal := make([]float64, 6)

		// 遍历订单，计算七天价格总和
		for i := 0; i < len(orders); i++ {
			orderTime := orders[i].Time
			daysDiff := int(today.Sub(orderTime).Hours() / 24)
			if daysDiff >= 0 && daysDiff < 365 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[0] += price
			} else if daysDiff >= 365 && daysDiff < 730 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[1] += price
			} else if daysDiff >= 730 && daysDiff < 1095 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[2] += price
			} else if daysDiff >= 1095 && daysDiff < 1460 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[3] += price
			} else if daysDiff >= 1460 && daysDiff < 1825 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[4] += price
			} else if daysDiff >= 1825 && daysDiff < 2190 {
				price, _ := strconv.ParseFloat(orders[i].Price, 64)
				sevenDayTotal[5] += price
			}
		}
		if len(sevenDayTotal) != 0 {
			pridic = append(pridic, sevenDayTotal[0])
		}
		// 构建 x 轴数据（1, 2, 3）
		x := make([]float64, len(orders))
		for i := 0; i < len(x); i++ {
			x[i] = float64(i + 1)
		}
		// 执行线性回归
		slope, intercept := linearRegression(x, sevenDayTotal)
		// 输出线性回归方程
		// 计算后三周的预测数据
		predictions := make([]float64, 2)
		for i := 0; i < 2; i++ {
			xVal := float64(len(sevenDayTotal) + i + 1)
			predictions[i] = slope*xVal + intercept
		}
		for i := 0; i < len(predictions); i++ {
			pridic = append(pridic, predictions[i])
		}
	}
	return pridic, err
}
