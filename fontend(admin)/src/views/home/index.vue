<template>
  <div class="dataVisualize-box">
    <div class="card top-box">
      <div class="top-title">Dashboard</div>
      <div class="top-content">
        <el-row :gutter="40">
          <el-col class="mb40" :xs="24" :sm="12" :md="12" :lg="6" :xl="6">
            <img src="./images/money.jpg" alt="" style="width: 90%; height: 80%" />
            <div class="mysmall">
              <span class="mytext">￥{{ totalIncome }}</span>
            </div>
          </el-col>
          <el-col class="mb40" :xs="24" :sm="12" :md="12" :lg="8" :xl="8">
            <div class="item-center">
              <div class="gitee-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/add_person.png" alt="" />
                </div>
                <span class="item-value">{{ todayNewUsers }}</span>
                <span class="traffic-name sle">Today's new user</span>
              </div>
              <div class="gitHub-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/add_team.png" alt="" />
                </div>
                <span class="item-value">{{ totalUsers }}</span>
                <span class="traffic-name sle">Total users</span>
              </div>
              <div class="today-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/today.png" alt="" />
                </div>
                <span class="item-value">{{ todayPosts }}</span>
                <span class="traffic-name sle">Today's posts</span>
              </div>
              <div class="yesterday-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/book_sum.png" alt="" />
                </div>
                <span class="item-value">{{ totalPosts }}</span>
                <span class="traffic-name sle">Total posts</span>
              </div>
            </div>
          </el-col>
          <el-col class="mb40" :xs="24" :sm="24" :md="24" :lg="10" :xl="10">
            <div class="item-right">
              <div class="echarts-title">Track mode</div>
              <div class="book-echarts">
                <Pie ref="pieRef" />
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts" name="dataVisualize">
import { ref, onMounted } from "vue";
import axios from "axios";
import Pie from "../dashboard/dataVisualize/components/pie.vue";

// 使用更具描述性的变量名
const totalIncome = ref(0); // 总收入
const todayNewUsers = ref(0); // 今日新增用户
const totalUsers = ref(0); // 总用户数
const todayPosts = ref(0); // 今日发帖数
const totalPosts = ref(0); // 总发帖数

// 定义获取数据的函数
const fetchData = async () => {
  try {
    const response = await axios.get("http://120.46.81.37:8080/admin/index"); // 后端地址
    totalIncome.value = response.data.totalIncome;
    todayNewUsers.value = response.data.todayNewUsers;
    totalUsers.value = response.data.totalUsers;
    todayPosts.value = response.data.todayPosts;
    totalPosts.value = response.data.totalPosts;
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};

// 当组件挂载后，获取数据
onMounted(fetchData);
</script>

<style scoped lang="scss">
@import "../dashboard/dataVisualize/index.scss";
.mysmall {
  margin: 10px auto;
  width: 128px;
}
.mytext {
  font-weight: bold;
  font-size: 30px;
  text-align: center;
  background-color: rgb(245, 184, 41);
  border-radius: 10px;
  color: white;
  padding: 5px;
  padding-left: 15px;
  padding-right: 15px;
}
</style>
