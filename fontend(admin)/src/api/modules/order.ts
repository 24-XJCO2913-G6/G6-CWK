import axios from "axios";

// 获取订单列表的 API 函数
export async function getOrderList() {
  try {
    const response = await axios.get("http://120.46.81.37:8080/admin/orders");
    return response.data; // 返回获取到的订单数据
  } catch (error) {
    // 错误处理逻辑
    console.error("Error fetching order list:", error);
    throw error;
  }
}
