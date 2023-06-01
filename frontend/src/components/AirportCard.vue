<template>
  <a-card
    style="min-width: 20rem; width: 75rem; height: 100%; margin: 0.5rem"
    headStyle="height: 4rem"
    bodyStyle="padding: 0; height: calc(100% - 4rem); overflow-y: auto;"
  >
    <template #title>
      {{ airport }}
    </template>
    <template #extra>
      <a-space>
        <a-tag
          v-if="socketStatus === 1"
          color="green"
        >ONLINE</a-tag>
        <a-tag
          v-else-if="socketStatus === 0"
          color="yellow"
        >CONNECTING</a-tag>
        <a-tag
          v-else
          color="red"
        >OFFLINE</a-tag>
        <sync-outlined
          class="refresh"
          :spin="loading"
          @click="refreshList"
        />
        <close-outlined
          class="close"
          @click="close"
        />
      </a-space>
    </template>
    <a-table
      sticky
      :columns="columns"
      :pagination="false"
      :data-source="data"
      size="middle"
      key="callsign"
    >
      <template #bodyCell="{ column, text, index }">
        <template v-if="column.type === 'sequence'">
          {{ index + 1 }}
        </template>
        <template v-if="column.dataIndex === 'status'">
          <span v-if="text === 70">Wait for Clearance</span>
          <span v-if="text === 60">Clearance Granted</span>
          <span v-if="text === 50">Wait for Push/Start</span>
          <span v-if="text === 40">Push-Back</span>
          <span v-if="text === 30">Wait for Taxi</span>
          <span v-if="text === 20">Taxi</span>
          <span v-if="text === 10">Wait for Take-Off</span>
          <span v-if="text === 0">Take-Off</span>
        </template>
      </template>
    </a-table>
  </a-card>
</template>

<script setup>
import { SyncOutlined, CloseOutlined } from '@ant-design/icons-vue';
import { onMounted, onUnmounted, ref, toRefs } from 'vue';
import api from '../api';
import { message } from 'ant-design-vue';

const columns = [
  { title: 'Sequence', type: 'sequence', align: "center" },
  { title: 'Callsign', dataIndex: 'callsign', align: "center" },
  { title: 'Destination', dataIndex: ['extra', 'arrival'], align: "center" },
  { title: 'Status', dataIndex: 'status', align: "center" },
];

const data = ref([]);
const props = defineProps(['airport']);
const { airport } = toRefs(props);
const socket = ref(null);
const socketReconnectTimeSeconds = ref(5);
const socketReconnectTimer = ref(null);
const loading = ref(false);
const socketStatus = ref(0);
const emit = defineEmits(['close']);

function createWebsocket() {
  socketStatus.value = 0;
  const wsUrlBase = `${window.location.protocol === "https:" ? "wss:" : "ws:"}//${window.location.host}`;
  socket.value = new WebSocket(`${wsUrlBase}/v1/${airport.value}/ws`);
  refreshList();
  socket.value.onopen = () => {
    socketStatus.value = 1;
  }
  socket.value.onmessage = (message) => {
    data.value = JSON.parse(message.data);
  };
  socket.value.onerror = (e) => {
    socketStatus.value = -1;
    if (socketReconnectTimeSeconds.value === -1 || e.code !== 1006) {
      return;
    }
    if (socketReconnectTimer.value !== null) {
      clearTimeout(socketReconnectTimer.value);
    }
    socketReconnectTimer.value = setTimeout(() => {
      socketReconnectTimer.value = null;
      createWebsocket();
    }, socketReconnectTimeSeconds.value * 1000);
  };
  socket.value.onclose = socket.value.onerror;
}

async function refreshList() {
  loading.value = true;
  try {
    const result = await api.get(`/v1/${airport.value}/queue`);
    data.value = result.data;
    message.success("List Reloaded.");
  } catch (e) {
    message.error("An error occurred.");
  } finally {
    loading.value = false;
  }
}

function close() {
  emit('close', airport.value);
}

onMounted(() => {
  if (socket.value === null) {
    createWebsocket();
  }
});

onUnmounted(() => {
  if (socket.value !== null) {
    socket.value.close();
  }
});
</script>

<style scoped>
.close {
  cursor: pointer;
}

.refresh {
  cursor: pointer;
}
</style>