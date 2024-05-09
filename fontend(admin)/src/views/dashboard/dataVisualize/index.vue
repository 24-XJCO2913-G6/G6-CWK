<template>
  <div class="table-box">
    <ProTable ref="proTable" :columns="columns" :data="users" @darg-sort="sortTable">
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
        <!-- <el-button type="primary" :icon="Edit" @click="openDrawer('编辑', scope.row)">Edit</el-button>
        <el-button type="primary" :icon="Refresh" @click="resetPass(scope.row)">Reset</el-button> -->
        <el-button type="primary" :icon="Delete" @click="deleteAccount(scope.row)">Del</el-button>
      </template>
    </ProTable>
    <UserDrawer ref="drawerRef" />
    <!-- <ImportExcel ref="dialogRef" /> -->
  </div>
</template>

<script setup lang="tsx" name="useProTable">
import { ref, reactive } from "vue";
import { Delete, Download } from "@element-plus/icons-vue";
import { User } from "@/api/interface";
import { useHandleData } from "@/hooks/useHandleData";
import { useDownload } from "@/hooks/useDownload";

import { ElMessage, ElMessageBox } from "element-plus";
import ProTable from "@/components/ProTable/index.vue";
// import ImportExcel from "@/components/ImportExcel/index.vue";
// import UserDrawer from "@/views/proTable/components/UserDrawer.vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";

import {
  deleteUser,
  exportUserInfo

  // getUserGender
} from "@/api/modules/user";

const users = [
  {
    date: "2024-04-22 13:34:07",
    ip_address: " 192.168.245.1",
    equipment: "Chrome",
    user_id: 1,
    function: "log in"
  },
  {
    date: "2024-04-23 08:15:30",
    ip_address: "192.168.1.10",
    equipment: "Firefox",
    user_id: 2,
    function: "log in"
  },
  {
    date: "2024-04-23 10:00:00",
    ip_address: "192.168.1.3",
    equipment: "HUAWEI P40",
    user_id: 1,
    function: "log in"
  },
  {
    date: "2024-04-24 16:30:00",
    ip_address: "192.168.1.6",
    equipment: "OPPO Reno9",
    user_id: 2,
    function: "collect the post"
  },
  {
    date: "2024-04-25 12:15:00",
    ip_address: "192.168.1.9",
    equipment: "iPhone 13",
    user_id: 3,
    function: "upload file"
  },
  {
    date: "2024-04-26 18:00:00",
    ip_address: "192.168.1.12",
    equipment: "Samsung Galaxy S23",
    user_id: 5,
    function: "log out"
  },
  {
    date: "2024-04-27 15:40:00",
    ip_address: "192.168.1.15",
    equipment: "Xiaomi 12",
    user_id: 4,
    function: "edit profile"
  },
  {
    date: "2024-04-24 14:45:02",
    ip_address: "192.168.0.1",
    equipment: "Safari",
    user_id: 3,
    function: "log out"
  },
  {
    date: "2024-04-25 10:20:55",
    ip_address: "192.168.2.15",
    equipment: "Chrome",
    user_id: 4,
    function: "view profile"
  },
  {
    date: "2024-04-27 09:05:18",
    ip_address: "192.168.4.50",
    equipment: "Internet Explorer",
    user_id: 6,
    function: "log in"
  },
  {
    date: "2024-04-28 18:25:00",
    ip_address: "192.168.5.100",
    equipment: "Opera",
    user_id: 1,
    function: "delete account"
  }
];
// ProTable 实例
const proTable = ref<ProTableInstance>();
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
    prop: "date",
    label: "Date",
    search: { el: "input", tooltip: "我是搜索提示" }
  },
  {
    prop: "ip_address",
    label: " IP address",
    // headerRender,
    width: 180,
    search: {
      el: "date-picker",
      span: 2,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" },
      defaultValue: ["2022-11-12 11:35:00", "2022-12-12 11:35:00"]
    }
  },
  { prop: "equipment", label: "Equipment" },
  {
    // 多级 prop
    prop: "user_id",
    label: "User ID",
    formatter: row => {
      return row.is_vip ? `✔` : `❌`;
    }
  },

  { prop: "function", label: "Function" },

  { prop: "operation", label: "Operations", fixed: "right", width: 200 }
]);

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
  console.log(proTable.value?.tableData);
  ElMessage.success("修改列表排序成功");
};

// 删除用户信息
const deleteAccount = async (params: User.ResUserList) => {
  await useHandleData(deleteUser, { id: [params.id] }, `删除【${params.username}】用户`);
  proTable.value?.getTableList();
};

// 批量删除用户信息
const batchDelete = async (id: string[]) => {
  await useHandleData(deleteUser, { id }, "删除所选用户信息");
  proTable.value?.clearSelection();
  proTable.value?.getTableList();
};

// // 重置用户密码
// const resetPass = async (params: User.ResUserList) => {
//   await useHandleData(resetUserPassWord, { id: params.id }, `重置【${params.username}】用户密码`);
//   proTable.value?.getTableList();
// };

// 导出用户列表
const downloadFile = async () => {
  ElMessageBox.confirm("确认导出用户数据?", "温馨提示", { type: "warning" }).then(() =>
    useDownload(exportUserInfo, "用户列表", proTable.value?.searchParam)
  );
};

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
