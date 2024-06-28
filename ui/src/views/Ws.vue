<template>
  <span style="white-space: pre-line">{{ reciveMsg }}</span>
  <h1>{{ msg }}</h1>
  <input v-model="msg" placeholder="edit me" />
  <button @click="sendMsg(msg)">发送</button>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";

const msg = ref(null);
const reciveMsg = ref("");

let socketUrl: any = "ws://127.0.0.1:8081/api/ws"; // socket地址
let websocket: any = null; // websocket 实例
let heartTime: any = null; // 心跳定时器实例
let socketHeart: number = 0; // 心跳次数
let HeartTimeOut: number = 3000; // 心跳超时时间
let socketError: number = 0; // 错误次数

// 初始化socket
const initWebSocket = (url: any) => {
  socketUrl = url;
  // 初始化 websocket
  websocket = new WebSocket(url);
  websocketonopen();
  websocketonmessage();
  websocketonerror();
  websocketclose();
  sendSocketHeart();
};

onMounted(() => {
  // 初始化websocket
  initWebSocket(socketUrl);
});

// socket 连接成功
const websocketonopen = () => {
  websocket.onopen = function (e: any) {
    console.log("连接 websocket 成功", e);
    resetHeart();
  };
};

// socket 连接失败
const websocketonerror = () => {
  websocket.onerror = function (e: any) {
    console.log("连接 websocket 失败", e);
  };
};

// socket 断开链接
const websocketclose = () => {
  websocket.onclose = function (e: any) {
    console.log("断开连接", e);
  };
};

// socket 接收数据
const websocketonmessage = () => {
  websocket.onmessage = function (e: any) {
    console.log("收到socket消息:---->", e);

    let msg = JSON.parse(e.data);
    if (msg.type === "heartbeat") {
      resetHeart();
      console.log("心跳");
    }
    console.log("收到socket消息", JSON.parse(e.data));
    // reciveMsg.value = reciveMsg.value +'\n'+ e.data;
    test(e.data); // 测试数据
  };
};

// socket 发送数据
const sendMsg = (data: any) => {
  websocket.send(data);
};

// socket 重置心跳
const resetHeart = () => {
  socketHeart = 0;
  socketError = 0;
  clearInterval(heartTime);
  sendSocketHeart();
};

// socket心跳发送
const sendSocketHeart = () => {
  console.log(websocket);
  heartTime = setInterval(() => {
    // 如果连接正常则发送心跳
    if (websocket.readyState == 1) {
      // if (socketHeart <= 30) {
      console.log("心跳发送：", socketHeart);
      websocket.send(
        JSON.stringify({
          type: "ping",
        })
      );
      socketHeart = socketHeart + 1;
    } else {
      // 重连
      reconnect();
    }
  }, HeartTimeOut);
};

// socket重连
const reconnect = () => {
  //   if (socketError <= 2) {
  clearInterval(heartTime);
  initWebSocket(socketUrl);
  socketError = socketError + 1;
  console.log("socket重连", socketError);
  //   } else {
  //     console.log("重试次数已用完的逻辑", socketError);
  //     clearInterval(heartTime);
  //   }
};

// 测试收到消息传递
const test = (msg: any) => {
  const { data } = msg;
  console.log("+++++++++ //// > ", data);
  console.log("+++++++++ > ", msg);
  console.log("+++++++++ > ", reciveMsg.value);

  if (reciveMsg.value == null) {
    reciveMsg.value = msg;
  } else {
    reciveMsg.value = reciveMsg.value + "\n" + msg;
  }
  // console.log(msg,444);
  // switch (msg.type) {
  //   case 'heartbeat': //加入会议
  //     mitts.emit('heartbeat', msg)
  //     break;
  // }
};
</script>

<style scoped>
</style>