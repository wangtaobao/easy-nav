<template>
  <div class="box">
    <a-form
      ref="formRef"
      :model="formState"
      name="basic"
      :label-col="{ span: 4 }"
      :wrapper-col="{ span: 20 }"
      autocomplete="off"
      @finish="onFinish"
      @finishFailed="onFinishFailed"
    >
      <a-form-item label="ID" name="ID" ref="ID" v-show="false">
        <a-input v-model:value="formState.ID" disabled />
      </a-form-item>
      <a-form-item
        label="名称"
        name="name"
        ref="name"
        :rules="[{ required: true, message: '导航地址名称不能为空' }]"
      >
        <a-input v-model:value="formState.name" />
      </a-form-item>

      <a-form-item
        label="地址"
        name="url"
        ref="url"
        :rules="[{ required: true, message: '导航地址不能为空' }]"
      >
        <a-input v-model:value="formState.url" />
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
        <a-button
          type="primary"
          html-type="submit"
          :loading="confirmLoading"
          :disabled="disabled"
          >确定</a-button
        >
      </a-form-item>
    </a-form>
  </div>
</template>
<script lang="ts" setup>
import { reactive, ref, computed, onMounted } from "vue";
import { message } from "ant-design-vue";
import { addUrl, updateUrl } from "@/api/index/index";

const confirmLoading = ref<boolean>(false);

const formRef = ref();

const props = defineProps(["oneData"]);

// const at = ()=>{
// console.log("++++&&& " + JSON.stringify(props.oneData, null, 2));
// }

const emit = defineEmits(["closeAddUrl", "success"]);

interface FormState {
  name: string;
  url: string;
  ID: number;
}

const formState = reactive<FormState>({
  name: "",
  url: "",
  ID: null,
});

// if (props.oneData != null) {
//   formState.ID = props.oneData.ID;
//   formState.name = props.oneData.name;
//   formState.url = props.oneData.url;
// }

const onFinish = async (values: any) => {
  confirmLoading.value = true;
  // console.log("等待添加:", values);

  if (formState.ID != null) {
    // console.log("更新方法");
    const { data } = await updateUrl(formState);
    console.log(data);
    if (data.code == 200) {
      message.success("导航地址更新成功");
      emit("success");
    }
    confirmLoading.value = false;
    resetForm();
  } else {
    // console.log("新增方法");
    const { data } = await addUrl(formState);
    console.log(data);
    if (data == null) {
      message.error("添加失败, 请联系开发人员");
      confirmLoading.value = false;
    }
    if (data.code == 200) {
      emit("success");
      confirmLoading.value = false;
      message.success("导航地址添加成功");
      resetForm();
    } else {
      message.error("添加失败, 请联系开发人员");
      confirmLoading.value = false;
    }
  }
};

const onFinishFailed = (errorInfo: any) => {
  console.log("Failed:", errorInfo);
};

const disabled = computed(() => {
  return !(formState.name && formState.url);
});

const resetForm = () => {
  formRef.value.resetFields();
};

onMounted(() => {
  if (props.oneData != null) {
    // console.log(
    //   ">>>>>>>>>>>>>>>>>>>> " + JSON.stringify(props.oneData, null, 2)
    // );
    formState.ID = props.oneData.ID;
    formState.name = props.oneData.name;
    formState.url = props.oneData.url;
  }
});
</script>
  

<style scoped>
.box {
  height: 150px;
}
</style>