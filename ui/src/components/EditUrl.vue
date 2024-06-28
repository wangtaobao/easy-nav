<template>
  <a-table
    :dataSource="dataSource"
    :columns="columns"
    :pagination="false"
    @resizeColumn="handleResizeColumn"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'tag'">
        <span>
          <a-tag color="green"> 标签 </a-tag>
        </span>
      </template>

      <template v-if="column.key === 'operation'">
        <span>
          <a-button type="text" shape="circle" @click="editData(record)">
            <EditOutlined />
          </a-button>
        </span>
        <span>
          <a-popconfirm
            v-if="record.ID"
            title="Are you ok?"
            @confirm="delData(record.ID)"
            ok-text="确定"
            cancel-text="取消"
          >
            <a-button type="text" shape="circle">
              <DeleteOutlined />
            </a-button>
          </a-popconfirm>
        </span>
      </template>
    </template>
  </a-table>

  <!-- 修改 -->
  <div>
    <a-modal v-model:open="open" title="修改导航路径" centered>
      <template #footer>
        <AddUrl @success="success" :oneData="oneData" :key="someUniqueKey" />
      </template>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import { defineComponent, h, onMounted, reactive, ref, toRaw } from "vue";
import { message, Modal, TableColumnsType } from "ant-design-vue";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons-vue";
import { delUrl } from "@/api/index/index";
import type { UnwrapRef } from "vue";
import { cloneDeep } from "lodash-es";

const emit = defineEmits(["editMethod"]);

const props = defineProps(["dataSource"]);
const open = ref<boolean>(false);
const oneData = ref("");
const someUniqueKey = ref("");

const dsClone = props.dataSource;

// console.log("@@@@@@@@@@@@@@@ " + JSON.stringify(props.dataSource, null, 2));

const columns = ref<TableColumnsType>([
  {
    title: "名称",
    dataIndex: "name",
    key: "name",
    resizable: true,
    width: 250,
  },
  {
    title: "地址",
    dataIndex: "url",
    key: "url",
    resizable: true,
    width: 250,
  },
  {
    title: "标签",
    dataIndex: "tag",
    key: "tag",
    resizable: true,
    width: 200,
  },
  {
    title: "操作",
    key: "operation",
    dataIndex: "operation",
    width: 100,
  },
]);

// interface ApiResponse {
//   code: number;
//   msg: string;
// }

const handleResizeColumn = (w: any, col: any) => {
  col.width = w;
};

const editData = (row: any) => {
  someUniqueKey.value = row.ID + row.name + row.url;
  oneData.value = row;
  // console.log("编辑数据: " + JSON.stringify(oneData.value, null, 2));
  open.value = true;
};

const success = () => {
  console.log("数据更新成功");
  emit("editMethod");
  open.value = false;
  // Modal.destroyAll();
};

const delData = async (id: any) => {
  // console.log("删除数据: " + id);
  // const { data } = await delUrl<ApiResponse>(id);
  const { data } = await delUrl(id);
  // console.log("----- " + data.code);
  // console.log("删除结果: " + JSON.stringify(data, null, 2));
  if (data.code === 200) {
    emit("editMethod");
    message.success("删除成功");
  } else {
    message.error("删除失败");
  }
};

// onMounted(() => {
//   console.log("++++");
// });
</script>
