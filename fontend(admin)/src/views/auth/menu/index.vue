<template>
  <div class="table-box">
    <ProTable ref="proTable" :columns="columns" :data="posts" @darg-sort="sortTable">
      <!-- 表格 header 按钮 -->
      <template #tableHeader="scope">
        <el-button v-auth="'export'" type="primary" :icon="Download" plain @click="downloadFile">Import data</el-button>

        <el-button type="danger" :icon="Delete" plain :disabled="!scope.isSelected" @click="batchDelete(scope.selectedListIds)">
          Batch delete
        </el-button>
      </template>
      <!-- Expand -->
      <template #expand="scope">
        {{ scope.row }}
      </template>

      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" :icon="Search" @click="showPostDetail(scope.row)">View Detail</el-button>
        <!-- <el-button type="primary" :icon="Search" @click="openDrawer('编辑', scope.row)">View Detail</el-button> -->
        <!-- <el-button type="primary" :icon="Delete" @click="deleteAccount(scope.row)">Delete</el-button> -->
      </template>
      <!-- <PostDetailDialog :post="selectedPost" /> -->
      <!-- <PostDetailDialog v-if="postDialogVisible" :post="selectedPost" @close="closeDialog" /> -->
    </ProTable>
    <!-- <UserDrawer ref="drawerRef" /> -->
    <PostDetailDialog v-if="postDialogVisible" :post="selectedPost" @close="closeDialog" />
    <!-- <ImportExcel ref="dialogRef" /> -->
  </div>
</template>

<script setup lang="tsx" name="useProTable">
import { ref, reactive, onMounted } from "vue";
import { Delete, Search, Download } from "@element-plus/icons-vue";
import { User } from "@/api/interface";
import { useHandleData } from "@/hooks/useHandleData";
import { useDownload } from "@/hooks/useDownload";

import { ElMessage, ElMessageBox } from "element-plus";
import ProTable from "@/components/ProTable/index.vue";
import PostDetailDialog from "@/components/ImportExcel/index.vue";
// import PostDetailDialog from "@/views/proTable/components/PostDetailDialog.vue";
// import UserDrawer from "@/views/proTable/components/UserDrawer.vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import axios from "axios";
import {
  deleteUser,
  exportUserInfo

  // getUserGender
} from "@/api/modules/user";

// const postsApiEndpoint = '/api/posts';

// 使用 ref 创建响应式引用
const posts = ref([
  {
    id: 1,
    name: "Yodo",
    post_time: "2024-04-02",
    post_mode: "VIP only",
    post_title: "Nice day to have a walk",
    post_state: "checked",
    track_id: "1"
  },
  {
    id: 2,
    name: "Veroooo",
    post_time: "2024-04-13",
    post_mode: "VIP only",
    post_title: "Running on the playground",
    post_state: "checked",
    track_id: "2"
  },
  {
    id: 3,
    name: "Peace of summer",
    post_time: "2024-04-06",
    post_mode: "VIP only",
    post_title: "I like lovely swan",
    post_state: "checked",
    track_id: "3"
  }
]);
const selectedPost = ref(null);
const postDialogVisible = ref(false);

// // 定义 selectedPost 响应式引用
// const selectedPost = ref(null);

// const posts = [
//   {
//     id: 1,
//     name: "Yodo",
//     post_time: "2024-04-02",
//     post_mode: "VIP only",
//     post_title: "Nice day to have a walk",
//     post_state: "checked",
//     track_id: "1"
//   },
//   {
//     id: 2,
//     name: "Veroooo",
//     post_time: "2024-04-13",
//     post_mode: "VIP only",
//     post_title: "Running on the playground",
//     post_state: "checked",
//     track_id: "2"
//   },
//   {
//     id: 3,
//     name: "Peace of summer",
//     post_time: "2024-04-06",
//     post_mode: "VIP only",
//     post_title: "I like lovely swan",
//     post_state: "checked",
//     track_id: "3"
//   }
// ];
// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<User.ResUserList>[]>([
  { type: "selection", fixed: "left", width: 70 },

  {
    prop: "name",
    label: "User Name",
    search: { el: "input", tooltip: "我是搜索提示" }
  },
  {
    prop: "post_time",
    label: "Post time",
    // headerRender,
    width: 180,
    search: {
      el: "date-picker",
      span: 2,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" },
      defaultValue: ["2022-11-12 11:35:00", "2022-12-12 11:35:00"]
    }
  },
  { prop: "post_title", label: "Post title" },
  {
    prop: "post_mode",
    label: "Post mode"
  },

  {
    prop: "post_state",
    label: "Post state"
  },
  {
    prop: "track_id",
    label: "Track id"
  },

  { prop: "operation", label: "Operations", fixed: "right", width: 200 }
]);

// 从后端获取帖子数据的函数
const fetchPosts = async () => {
  try {
    const response = await axios.get("http://120.46.81.37:8080/admin/pending");
    posts.value = response.data; // 假设返回的数据是帖子列表
  } catch (error) {
    console.error("Error fetching posts: ", error);
  }
};

// 获取帖子详情的函数
const getPostDetail = async postId => {
  try {
    const response = await axios.get(`http://120.46.81.37:8080/admin/pendingDetail/${postId}`);
    return response.data; // 假设返回的数据包含帖子的详细信息
  } catch (error) {
    console.error("Error fetching post detail: ", error);
    throw error;
  }
};

// 显示帖子详情对话框的方法
const showPostDetail = async post => {
  try {
    const detail = await getPostDetail(post.id); // 假设帖子的 ID 是 post.id
    selectedPost.value = detail;
    postDialogVisible.value = true;
  } catch (error) {
    // 处理错误，例如显示错误消息
    ElMessage.error("Failed to load post detail");
  }
};

// 组件挂载时调用 fetchPosts 函数
onMounted(fetchPosts);

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
  console.log(proTable.value?.tableData);
  ElMessage.success("修改列表排序成功");
};

// // 删除用户信息
// const deleteAccount = async (params: User.ResUserList) => {
//   await useHandleData(deleteUser, { id: [params.id] }, `删除【${params.username}】用户`);
//   proTable.value?.getTableList();
// };

// 批量删除用户信息
const batchDelete = async (id: string[]) => {
  await useHandleData(deleteUser, { id }, "删除所选用户信息");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// 导出用户列表
const downloadFile = async () => {
  ElMessageBox.confirm("确认导出用户数据?", "温馨提示", { type: "warning" }).then(() =>
    useDownload(exportUserInfo, "用户列表", proTable.value?.searchParam)
  );
};

// // 控制 PostDetailDialog 显示的引用
// const postDialogVisible = ref(false);

// // 显示帖子详情对话框的方法
// const showPostDetail = post => {
//   // 更新当前选中的帖子
//   selectedPost.value = post;
//   // 显示对话框
//   postDialogVisible.value = true;
// };

const closeDialog = () => {
  // 隐藏对话框
  postDialogVisible.value = false;
  // 清除选中的帖子
  selectedPost.value = null;
};
// // 修改点击事件处理函数
// const showPostDetail = post => {
//   // 假设你已经定义了 PostDetailDialog 组件
//   ElMessageBox({
//     title: "Post Detail",
//     message: () => PostDetailDialog,
//     component: PostDetailDialog,
//     customClass: "post-detail-dialog",
//     // 将帖子数据传递给对话框组件
//     componentProps: { post },
//     showClose: true,
//     modal: true,
//     lockScroll: true
//   })
//     .then(() => {
//       // 帖子处理成功的逻辑
//     })
//     .catch(() => {
//       // 用户取消查看帖子的逻辑
//     });
// };

// // 打开 drawer(新增、查看、编辑)
// const drawerRef = ref<InstanceType<typeof UserDrawer> | null>(null);
// const openDrawer = (title: string, row: Partial<User.ResUserList> = {}) => {
//   const params = {
//     title,
//     isView: title === "查看",
//     row: { ...row },
//     api: title === "新增" ? addUser : title === "编辑" ? editUser : undefined,
//     getTableList: proTable.value?.getTableList
//   };
//   drawerRef.value?.acceptParams(params);
// };
</script>
