<template>
  <div class="dataVisualize-box">
    <div class="card top-box">
      <div class="top-title">Dashboard</div>
      <div class="top-content">
        <el-row :gutter="40">
          <el-col class="mb40" :xs="24" :sm="12" :md="12" :lg="6" :xl="6">
            <img src="./images/money.jpg" alt="" style="width: 90%; height: 80%" />
            <div class="mysmall"><span class="mytext">￥887.8</span></div>
          </el-col>
          <el-col class="mb40" :xs="24" :sm="12" :md="12" :lg="8" :xl="8">
            <div class="item-center">
              <div class="gitee-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/add_person.png" alt="" />
                </div>
                <span class="item-value">2</span>
                <span class="traffic-name sle">Today's new user</span>
              </div>
              <div class="gitHub-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/add_team.png" alt="" />
                </div>
                <span class="item-value">28</span>
                <span class="traffic-name sle">Total users</span>
              </div>
              <div class="today-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/today.png" alt="" />
                </div>
                <span class="item-value">6</span>
                <span class="traffic-name sle">Today's posts</span>
              </div>
              <div class="yesterday-traffic traffic-box">
                <div class="traffic-img">
                  <img src="./images/book_sum.png" alt="" />
                </div>
                <span class="item-value">45</span>
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
import Pie from "../dashboard/dataVisualize/components/pie.vue";
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
