<template>
  <div class="table-box">
    <ProTable ref="proTable" :columns="columns" :data="users" @darg-sort="sortTable">
      <!-- 表格 header 按钮 -->
      <template #tableHeader="scope">
        <!-- <el-button v-auth="'export'" type="primary" :icon="Download" plain @click="downloadFile">Import data</el-button> -->

        <el-button type="danger" :icon="Delete" plain :disabled="!scope.isSelected" @click="batchDelete(scope.selectedListIds)">
          Batch delete
        </el-button>
      </template>
      <!-- Expand -->
      <template #expand="scope">
        {{ scope.row }}
      </template>

      <!-- 表格操作 -->
      <!-- <template #operation="scope">
        <el-button type="primary" :icon="Edit" @click="openDrawer('View', scope.row)">View Details</el-button> -->
      <!-- <el-button type="primary" :icon="Refresh" @click="resetPass(scope.row)">Reset</el-button>
        <el-button type="primary" :icon="Delete" @click="deleteAccount(scope.row)">Del</el-button> -->
      <!-- </template> -->
    </ProTable>
    <!-- <UserDrawer ref="drawerRef" /> -->
    <!-- <ImportExcel ref="dialogRef" /> -->
  </div>
</template>

<script setup lang="tsx" name="useProTable">
import { ref, reactive, onMounted } from "vue";
import { Delete } from "@element-plus/icons-vue";
import { User } from "@/api/interface";
import { useHandleData } from "@/hooks/useHandleData";
// import { useDownload } from "@/hooks/useDownload";

// import { ElMessage, ElMessageBox } from "element-plus";
import { ElMessage } from "element-plus";
import ProTable from "@/components/ProTable/index.vue";
// import ImportExcel from "@/components/ImportExcel/index.vue";
// import UserDrawer from "@/views/proTable/components/UserDrawer.vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";

import {
  getUserList,
  deleteUser

  // getUserGender
} from "@/api/modules/user";

// 如果你想在请求之前对当前请求参数做一些操作，可以自定义如下函数：params 为当前所有的请求参数（包括分页），最后返回请求列表接口
// 默认不做操作就直接在 ProTable 组件上绑定	:requestApi="getUserList"

// // 自定义渲染表头（使用tsx语法）
// const headerRender = (scope: <User.ResUserList>) => {
//   return (
//     <el-button type="primary" onClick={() => ElMessage.success("我是通过 tsx 语法渲染的表头")}>
//       {scope.column.label}
//     </el-button>
//   );
// };

// 表格配置项
const columns = reactive<ColumnProps<User.ResUserList>[]>([
  { type: "selection", fixed: "left", width: 70 },

  {
    prop: "name",
    label: "User Name",
    search: { el: "input", tooltip: "我是搜索提示" }
  },
  {
    prop: "created_time",
    label: "Created time",
    // headerRender,
    width: 180,
    search: {
      el: "date-picker",
      span: 2,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" },
      defaultValue: ["2022-11-12 11:35:00", "2022-12-12 11:35:00"]
    }
  },
  { prop: "email", label: "Email" },
  {
    // 多级 prop
    prop: "is_vip",
    label: "VIP",
    formatter: row => {
      return row.is_vip ? `✔` : `❌`;
    }
  },
  {
    // 多级 prop
    prop: "vip_expiry_date",
    label: "Vip_expiry_date"
  },

  { prop: "password", label: "Password" }

  // { prop: "operation", label: "Operations", fixed: "right", width: 330 }
]);

// ProTable 实例
const proTable = ref<ProTableInstance>();

// // 使用 ref 创建用户数据的响应式引用
// const users = ref([]);

// 用户数据
const users = ref([
  {
    id: 1,
    name: "Alice",
    created_time: "2024-03-02",
    is_vip: true,
    password: "123456",
    email: "123456789@qq.com",
    vip_expiry_date: "2024-04-02"
  }
]);

// const users = ref([]);

// 加载用户数据的函数
const loadUsers = async () => {
  try {
    const response = await getUserList();

    users.value = response.data; // 假设返回的数据格式为 { data: [] }
    // proTable.value?.setTableData(response.data); // 假设 ProTable 组件有方法可以设置表格数据
  } catch (error) {
    ElMessage.error("加载用户数据失败");
  }
};

// 组件挂载时加载数据
onMounted(() => {
  loadUsers();
});

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
  console.log(proTable.value?.tableData);
  // ElMessage.success("修改列表排序成功");
};

// 删除用户信息
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

// 重置用户密码
// const resetPass = async (params: User.ResUserList) => {
//   await useHandleData(resetUserPassWord, { id: params.id }, `重置【${params.username}】用户密码`);
//   proTable.value?.getTableList();
// };

// 导出用户列表
// const downloadFile = async () => {
//   ElMessageBox.confirm("确认导出用户数据?", "温馨提示", { type: "warning" }).then(() =>
//     useDownload(exportUserInfo, "用户列表", proTable.value?.searchParam)
//   );
// };

// 打开 drawer(新增、查看、编辑)
// const drawerRef = ref<InstanceType<typeof UserDrawer> | null>(null);
// const openDrawer = (title: string, row: Partial<User.ResUserList> = {}) => {
//   const params = {
//     title,
//     isView: title === "View",
//     row: { ...row }
//     // api: title === "新增" ? addUser : title === "编辑" ? editUser : undefined,
//     // getTableList: proTable.value?.getTableList
//   };
//   drawerRef.value?.acceptParams(params);
// };
// const openDrawer = async (title: string, userId: number) => {
//   try {
//     // 假设获取用户详情的 API 需要用户 ID
//     const userDetail = await getUserDetailById(userId); // 替换为实际的 API 调用
//     drawerRef.value?.acceptParams({
//       title,
//       isView: true,
//       row: userDetail
//     });
//   } catch (error) {
//     ElMessage.error("Failed to load user details");
//   }
// };
</script>
