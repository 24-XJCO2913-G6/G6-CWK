<template>
  <div class="table-box">
    <ProTable ref="proTable" :columns="columns" :data="orders" @darg-sort="sortTable">
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
      <!-- <template #operation="scope">
        <el-button type="primary" :icon="Edit" @click="openDrawer('编辑', scope.row)">Edit</el-button>
        <el-button type="primary" :icon="Delete" @click="deleteAccount(scope.row)">Del</el-button>
      </template> -->
    </ProTable>
    <UserDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="tsx" name="useProTable">
import { ref, reactive, onMounted } from "vue";
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
// 导入 API 函数
import { getOrderList } from "@/api/modules/order";
// 响应式引用订单数据
const orders = ref([]);

// const orders = [
//   {
//     id: 1,
//     name: "Alice",
//     order_time: "2024-04-02",
//     order_content: "VIP for a month",
//     order_price: 29.9,
//     order_state: "Paid"
//   },
//   {
//     id: 2,
//     name: "Bob",
//     order_time: "2024-04-12",
//     order_content: "VIP for a year",
//     order_price: 199.9,
//     order_state: "Refunded,"
//   },
//   {
//     id: 3,
//     name: "Charlie",
//     order_time: "2024-03-10",
//     order_content: "Permanent VIP",
//     order_price: 999,
//     order_state: "Paid"
//   },
//   {
//     id: 4,
//     name: "David",
//     order_time: "2024-05-01",
//     order_content: "VIP for a quarter",
//     order_price: 69.9,
//     order_state: "Paid"
//   },
//   {
//     id: 5,
//     name: "Eve",
//     order_time: "2024-03-25",
//     order_content: "VIP for a year",
//     order_price: 199.9,
//     order_state: "Paid"
//   },
//   {
//     id: 6,
//     name: "Frank",
//     order_time: "2024-04-20",
//     order_content: "Permanent VIP",
//     order_price: 999,
//     order_state: "Paid"
//   },
//   {
//     id: 7,
//     name: "Grace",
//     order_time: "2024-03-15",
//     order_content: "VIP for a month",
//     order_price: 29.9,
//     order_state: "Paid"
//   },
//   {
//     id: 8,
//     name: "Henry",
//     order_time: "2024-05-03",
//     order_content: "VIP for a year",
//     order_price: 199.9,
//     order_state: "Refunded"
//   },
//   {
//     id: 9,
//     name: "Ivy",
//     order_time: "2024-04-05",
//     order_content: "VIP for a quarter",
//     order_price: 69.9,
//     order_state: "Paid"
//   }
// ];

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<User.ResUserList>[]>([
  { type: "selection", fixed: "left", width: 70 },
  { prop: "Content", label: "Order content" },
  {
    prop: "Oid",
    label: "Order Id",
    search: { el: "input", tooltip: "我是搜索提示" }
  },
  {
    prop: "Uid",
    label: "User Id",
    search: { el: "input", tooltip: "我是搜索提示" }
  },
  {
    // 多级 prop
    prop: "Price",
    label: "Order price"
  },

  {
    prop: "State",
    label: "Order state"
  },
  {
    prop: "Time",
    label: "Order time",
    // headerRender,
    width: 180,
    search: {
      el: "date-picker",
      span: 2,
      props: { type: "datetimerange", valueFormat: "YYYY-MM-DD HH:mm:ss" },
      defaultValue: ["2022-11-12 11:35:00", "2022-12-12 11:35:00"]
    }
  }
]);

// 加载订单数据
const loadOrders = async () => {
  try {
    const response = await getOrderList();
    // console.log(response);
    // orders.value = response.map(x => ({
    //   Uid: x.Uid,
    //   Time: x.Time,
    //   Content: x.Content,
    //   Price: x.Price,
    //   State: x.State
    // }));
    // console.log(orders.value);
    orders.value = response.orders;
    console.log(orders.value);
  } catch (error) {
    ElMessage.error("Failed to load orders");
  }
};

// 组件挂载时调用
onMounted(loadOrders);

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

// 打开 drawer(新增、查看、编辑)
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
