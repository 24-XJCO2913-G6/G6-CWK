<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title} User`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <!-- <el-form-item label="用户头像" prop="avatar">
        <UploadImg v-model:image-url="drawerProps.row!.avatar" width="135px" height="135px" :file-size="3">
          <template #empty>
            <el-icon><Avatar /></el-icon>
            <span>请上传头像</span>
          </template>
          <template #tip> 头像大小不能超过 3M </template>
        </UploadImg>
      </el-form-item> -->
      <el-form-item label="Avatar" prop="photo">
        <UploadImgs v-model:file-list="drawerProps.row.photo" height="140px" width="140px" border-radius="50%">
          <template #empty>
            <el-icon><Picture /></el-icon>
            <span>请上传照片</span>
          </template>
        </UploadImgs>
      </el-form-item>
      <el-form-item label="Nickname" prop="username">
        <el-input v-model="drawerProps.row.username" placeholder="请填写用户姓名" clearable></el-input>
      </el-form-item>
      <!-- <el-form-item label="Sex" prop="gender">
        <el-input v-model="drawerProps.row.gender" placeholder="请填写用户姓名" clearable></el-input>
      </el-form-item> -->
      <!-- <el-form-item label="Sex" prop="gender">
        <el-select v-model="drawerProps.row!.gender" placeholder="请选择性别" clearable>
          <el-option v-for="item in genderType" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item> -->
      <el-form-item label="Birthday" prop="birthday">
        <el-input v-model="drawerProps.row.Born" placeholder="请填写用户姓名" clearable></el-input>
      </el-form-item>
      <!-- <el-form-item label="身份证号" prop="idCard">
        <el-input v-model="drawerProps.row!.idCard" placeholder="请填写身份证号" clearable></el-input>
      </el-form-item> -->
      <el-form-item label="Status" prop="status">
        <el-input v-model="drawerProps.row.Status" placeholder="请填写邮箱" clearable></el-input>
      </el-form-item>
      <el-form-item label="Job" prop="job">
        <el-input v-model="drawerProps.row.Job" placeholder="请填写邮箱" clearable></el-input>
      </el-form-item>
      <el-form-item label="Hometown" prop="hometown">
        <el-input v-model="drawerProps.row.LivesIn" placeholder="请填写居住地址" clearable></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="UserDrawer">
import { ref, reactive } from "vue";
// import { reactive } from "vue";
// import { genderType } from "@/utils/dict";
// import { ElMessage, FormInstance } from "element-plus";
import { ElMessage } from "element-plus";
import { User } from "@/api/interface";
// import UploadImg from "@/components/Upload/Img.vue";
import UploadImgs from "@/components/Upload/Imgs.vue";
import {
  getUserDetailById

  // getUserGender
} from "@/api/modules/user";
// 抽屉属性接口
interface DrawerProps {
  title: string;
  isView: boolean;
  row: Partial<User.ResUserList>;
  api?: (params: any) => Promise<any>;
  getTableList?: () => void;
}

// // 使用 ref 声明 drawerVisible，它控制抽屉的显示与隐藏
// const drawerVisible = ref(false);

// // 使用 reactive 声明 drawerProps，它可以包含用户的详情数据和其他属性
// const drawerProps = reactive<DrawerProps>({
//   title: "View",
//   isView: true,
//   row: {} as Partial<User.ResUserList>
// });
const drawerProps = ref<DrawerProps>({
  title: "View",
  isView: true,
  row: {
    // 假数据：模拟的用户数据
    Uid: 1,
    username: "JohnDoe",
    photo: "", // 头像
    Born: "1990-01-01",
    Status: "Single",
    Job: "Software Engineer",
    LivesIn: "New York"
  } as Partial<User.ResUserList>
});

// 抽屉是否可见
const drawerVisible = ref(true);

const rules = reactive({
  // avatar: [{ required: true, message: "请上传用户头像" }],
  // photo: [{ required: true, message: "请上传用户照片" }],
  // username: [{ required: true, message: "请填写用户姓名" }],
  // gender: [{ required: true, message: "请选择性别" }],
  // idCard: [{ required: true, message: "请填写身份证号" }],
  // email: [{ required: true, message: "请填写邮箱" }],
  // address: [{ required: true, message: "请填写居住地址" }]
});

// 接收父组件传过来的参数
const acceptParams = (params: DrawerProps) => {
  // 仅当需要查看详情时才获取数据
  if (params.isView) {
    fetchUserDetail(params.row.id); // 假设 row 对象中包含 id 属性
  }
  drawerProps.value = params;
  drawerVisible.value = true;
};

// 获取用户详情的方法
const fetchUserDetail = async (userId: number) => {
  try {
    const response = await getUserDetailById(userId);
    // 更新 drawerProps.row 以显示用户详情
    drawerProps.value.row = response.data;
  } catch (error) {
    ElMessage.error("获取用户详情失败");
  }
};

interface DrawerProps {
  title: string;
  isView: boolean;
  row: Partial<User.ResUserList>;
  api?: (params: any) => Promise<any>;
  getTableList?: () => void;
}

// const drawerVisible = ref(false);
// const drawerProps = ref<DrawerProps>({
//   isView: false,
//   title: "",
//   row: {}
// });

// 接收父组件传过来的参数
// const acceptParams = (params: DrawerProps) => {
//   drawerProps.value = params;
//   drawerVisible.value = true;
// };

// 提交数据（新增/编辑）
// const ruleFormRef = ref<FormInstance>();
// const handleSubmit = () => {
//   ruleFormRef.value!.validate(async valid => {
//     if (!valid) return;
//     try {
//       await drawerProps.value.api!(drawerProps.value.row);
//       ElMessage.success({ message: `${drawerProps.value.title}用户成功！` });
//       drawerProps.value.getTableList!();
//       drawerVisible.value = false;
//     } catch (error) {
//       console.log(error);
//     }
//   });
// };

defineExpose({
  acceptParams
});
</script>
