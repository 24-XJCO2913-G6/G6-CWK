<template>
  <el-dialog :title="dialogTitle" v-model="visible" width="50%" custom-class="post-detail-dialog" @close="handleClose">
    <div class="post-detail">
      <header class="post-header">
        <h2 class="post-title">{{ post.title }}</h2>
        <p class="post-author"><strong>Author:</strong> {{ post.author }}</p>
      </header>
      <div class="post-media">
        <img :src="post.mapImageUrl" alt="Post Map" />
      </div>
      <div class="post-body">
        <strong>Content:</strong>
        <div v-html="post.content"></div>
      </div>
    </div>
    <template #footer>
      <span>
        <el-button type="primary" @click="visible = false">Close</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref } from "vue";

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
    <h3>It was a beautiful day at the park!</h3>
    <p>I went to the park with my friends and we had a picnic. The weather was sunny and warm.</p>
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
</style>
