<template>
  <el-dialog :title="dialogTitle" v-model="visible" width="50%" custom-class="post-detail-dialog" @close="handleClose">
    <div class="post-detail">
      <header class="post-header">
        <h2 class="post-title">{{ post.title }}</h2>
        <p class="post-author"><strong>Author:</strong> {{ post.author }}</p>
      </header>
      <div class="post-media">
        <div class="map-container">
          <div ref="mapContainer" style="width: 100%; height: 500px"></div>
        </div>
        <!-- <img :src="post.mapImageUrl" alt="Post Map" /> -->
      </div>
      <div class="post-body">
        <strong>Content:</strong>
        <div v-html="post.content"></div>
      </div>
    </div>
    <template #footer>
      <span>
        <span>
          <!-- 拒绝按钮 -->
          <el-button type="danger" @click="rejectPost">Reject</el-button>
          <!-- 通过按钮 -->
          <el-button type="success" @click="approvePost">Approve</el-button>
        </span>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, onMounted } from "vue";

// 通过审核的函数
const approvePost = () => {
  // 这里添加审核通过的逻辑，比如调用后端接口
  console.log("Post approved");
  // 关闭对话框
  visible.value = false;
};

// 拒绝审核的函数
const rejectPost = () => {
  // 这里添加审核拒绝的逻辑，比如调用后端接口
  console.log("Post rejected");
  // 关闭对话框
  visible.value = false;
};

// 引入高德地图 API
const script = document.createElement("script");
script.src = "https://webapi.amap.com/maps?v=1.4.15&key=b22dc976c0090ec2a8d130f4de04a4d3";
document.head.appendChild(script);

const mapContainer = ref(null);

onMounted(() => {
  // 初始化地图
  // 确保异步加载完成后再执行初始化操作
  script.onload = function () {
    // AMap代表高德地图的全局对象
    const map = new AMap.Map(mapContainer.value, {
      zoom: 10, // 地图缩放级别
      center: [116.39, 39.9] // 地图中心点坐标，北京
    });

    AMap.plugin("AMap.Polyline", function () {
      // 路线的坐标点数组
      const routePath = [
        [116.39, 39.9], // 起点
        [116.4, 39.91], // 途径点
        [116.41, 39.92] // 终点
      ];

      // 创建路线对象
      const polyline = new AMap.Polyline({
        map: map,
        path: routePath, // 路线路径
        strokeColor: "#007bff", // 路线颜色
        strokeOpacity: 0.2, // 路线透明度
        strokeWeight: 6, // 路线宽度
        strokeOpacity: 1, // 路线线条透明度
        isOutline: false, // 是否描边
        outlineColor: "#fff", // 描边颜色
        strokeStyle: "solid" // 路线的虚线样式
      });

      // 将路线添加到地图上
      map.add(polyline);
    });

    // 添加地图控件，如缩放控件、缩略图等
    AMap.plugin(["AMap.Scale", "AMap.OverView"], function () {
      map.addControl(new AMap.Scale());
      map.addControl(new AMap.OverView());
    });

    // // 添加标记
    // const marker = new AMap.Marker({
    //   position: [116.39, 39.9], // 北京
    //   map: map
    // });
  };
});

// 初始对话框不可见
const visible = ref(true);

// 对话框关闭时的钩子
const handleClose = () => {
  console.log("Dialog closed");
};

// 定义对话框标题
const dialogTitle = "Post Detail";

// 固定的帖子数据
const post = {
  title: "A Day in the Park",
  author: "Jane Doe",
  mapImageUrl: "https://via.placeholder.com/600x300",
  content: `
    <p>It was a beautiful day at the park! I went to the park with my friends and we had a picnic. The weather was sunny and warm.</p>
    <p>We saw many people walking their dogs and children playing on the playground.</p>
    <p>It was a perfect day to relax and enjoy the beauty of nature.</p>
  `
};
</script>
<style>
.post-detail-dialog .el-dialog__header {
  padding: 20px;
  text-align: center;
}

.post-header {
  padding: 20px;
  background-color: #f5f5f5;
  text-align: center;
}

.post-title {
  margin: 0;
  font-size: 1.5rem;
  color: #333;
}

.post-author {
  color: #888;
  font-size: 0.9rem;
  margin: 0;
}

.post-media img {
  width: 100%;
  max-height: 500px;
  object-fit: cover;
  margin: 15px 0;
}

.post-body {
  padding: 20px;
}

.post-content {
  word-wrap: break-word;
  white-space: pre-wrap;
}

.post-content p,
.post-content h3,
.post-content a {
  margin: 0.5rem 0;
  line-height: 1.5;
}

.post-content h3 {
  font-size: 1.2rem;
  color: #333;
}

.post-content a {
  color: #007bff;
  text-decoration: none;
}

/* 按钮样式调整 */
.el-dialog__footer {
  padding: 10px 20px;
  text-align: right;
  background-color: #f9f9f9;
}

.map-container {
  height: 500px;
}
</style>
