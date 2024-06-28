<template>
  <!-- <a-skeleton :loading="isLoading" active> -->
  <div>
    <a-layout>
      <a-layout-content class="search-layout">
        <div class="search">
          <a-input-search
            v-model:value="searchValue"
            placeholder="关键词"
            enter-button
            size="large"
            @search="onSearch"
          >
            <template #addonBefore>
              <a-select
                v-model:value="searchType"
                style="width: 110px"
                size="large"
              >
                <a-select-option value="Google">Google</a-select-option>
                <a-select-option value="百度">百度</a-select-option>
              </a-select>
            </template>
          </a-input-search>
        </div>
      </a-layout-content>
      <a-layout-content class="card-layout">
        <div class="box">
          <a-list
            :grid="{ gutter: 16, xs: 1, sm: 2, md: 4, lg: 4, xl: 6, xxl: 6 }"
            :data-source="data"
          >
            <template #renderItem="{ item }">
              <a-list-item>
                <div style="height: 115px" @click="enterButton(item.url)">
                  <a-card hoverable class="a-card">
                    <a-card-meta :title="item.name" :description="item.url" />
                  </a-card>
                </div>
              </a-list-item>
            </template>
          </a-list>
        </div>
      </a-layout-content>
    </a-layout>
  </div>
  <a-float-button-group
    trigger="hover"
    type="primary"
    :style="{ right: '24px' }"
  >
    <template #icon>
      <SettingOutlined />
    </template>
    <a-float-button @click="showEditModal">
      <template #tooltip>
        <div>修改导航路径</div>
      </template>
      <template #icon>
        <EditOutlined />
      </template>
    </a-float-button>
    <a-float-button @click="showModal">
      <template #tooltip>
        <div>添加导航路径</div>
      </template>
      <template #icon>
        <PlusOutlined />
      </template>
    </a-float-button>
  </a-float-button-group>
  <div>
    <a-modal v-model:open="open" title="添加导航路径" centered>
      <template #footer>
        <AddUrl
          :closeAddUrl="closeAddUrl"
          @success="success"
          :parent-method="parentMethod"
        />
      </template>
    </a-modal>
  </div>
  <!-- 编辑 -->
  <div>
    <a-modal
      v-model:open="openEdit"
      title="路径管理"
      width="50%"
      centered
      @ok="editHandleOk"
      :footer="null"
      :bodyStyle="{ maxHeight: '60vh', overflow: 'auto' }"
    >
      <EditUrl
        :dataSource="editData"
        @editMethod="editThen"
      />
    </a-modal>
  </div>
</template>

  
<script setup lang="ts">
import {
  EditOutlined,
  PlusOutlined,
  SettingOutlined,
} from "@ant-design/icons-vue";

import { ref, reactive, onMounted, toRefs } from "vue";

import AddUrl from "@/components/AddUrl.vue";
import EditUrl from "@/components/EditUrl.vue";

import { getIndexList } from "@/api/index/index";
import { message } from "ant-design-vue";

const isLoading = ref<boolean>(true);

const searchType = ref<string>("Google");

const searchValue = ref<string>("");

const open = ref<boolean>(false);
const openEdit = ref<boolean>(false);

const showModal = () => {
  open.value = true;
};

const showEditModal = () => {
  openEdit.value = true;
};

const editHandleOk = (e: MouseEvent) => {
  console.log(e);
  openEdit.value = false;
};

const editThen = () => {
  loadData(indexListObj);
};

const success = () => {
  closeAddUrl();
  loadData(indexListObj);
};

const closeAddUrl = () => {
  open.value = false;
};

const enterButton = (url: string) => {
  console.log("++");
  window.open(url, "_blank");
};

const onSearch = (onSearchValue: string) => {
  console.log("use value", onSearchValue);
  console.log("searchType:", searchType.value);

  if (searchValue.value === null || searchValue.value === "") {
    // console.log("没有关键词, 不搜索跳转");
    message.warn("没有搜索关键词!")
  } else if (searchType.value === "Google") {
    console.log(`https://www.google.com/search?q=${onSearchValue}`);
    window.open(`https://www.google.com/search?q=${onSearchValue}`, "_blank");
  } else {
    console.log(`https://www.baidu.com/s?wd=${onSearchValue}`);
    window.open(`https://www.baidu.com/s?wd=${onSearchValue}`, "_blank");
  }
  onSearchValue = "";
};

const indexListObj = reactive({
  data: [],
});

const loadData = async (indexListObj: any) => {
  const { data } = await getIndexList();
  // console.log("完整响应:", JSON.stringify(data, null, 2));

  const { result } = data;
  console.log(result);
  indexListObj.data = result;
  isLoading.value = false;
  editData.value = result;
};

const { data } = toRefs(indexListObj);

onMounted(() => {
  loadData(indexListObj);
});

const editData = ref([]);
</script>
  
<style scoped>
.ant-layout-content {
  padding: 24px;
  background: white;
}

.card-layout {
  min-height: 280px;
  /* width: 100px; */
  /* width: 15%; */
}

.search-layout {
  min-height: 100px;
}

.search {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 55%;
  min-height: 100px;
  margin: auto;
}
.edit-class {
  top: 0;
  left: 0;
}

.a-card {
  width: 240px;
  float: left;
  margin: 0.5em;
  /* margin: 0 auto; */
  transition: box-shadow 0.3s ease-in-out;
}

@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
  100% {
    transform: scale(1);
  }
}

.a-card:hover {
  animation: pulse 1s infinite;
}

/* .footer {
  position: fixed;
  bottom: 0;
  left: 0;
  height: 70px;
  background-color: white;
  width: 100%;
} */
</style>