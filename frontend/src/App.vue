<template>
  <a-layout>
    <a-layout-header :style="headerStyle">
      <div class="left">
        <h1>VATPRC QueueMaster</h1>
      </div>
      <div class="right">
        <a-input-group compact style="width: 10rem">
          <a-input v-model:value="newAirport" @change="() => { newAirport = newAirport.toUpperCase() }" placeholder="Airport ICAO" style="width: 7rem" />
          <a-button type="primary" @click="addList"><plus-outlined /></a-button>
        </a-input-group>
      </div>
    </a-layout-header>
    <a-layout-content :style="contentStyle">
      <div style="display: flex; flex-direction: vertical; justify-content: center; height: 100%;">
        <AirportCard v-for="airport of airportList" @close="closeList" :airport="airport" />
        <a-result v-if="airportList.length === 0" status="404" title="There's nothing here." sub-title="You can add an airport queue on the top right.">
      </a-result>
      </div>
    </a-layout-content>
  </a-layout>
</template>

<script setup>
import { PlusOutlined } from '@ant-design/icons-vue';
import { ref } from 'vue';

const airportList = ref([]);
const newAirport = ref("");

function closeList(airport) {
  airportList.value = airportList.value.filter(elem => elem !== airport);
} 

function addList() {
  if (newAirport.value.length === 0) {
    return;
  }
  airportList.value.push(newAirport.value);
  newAirport.value = "";
}


const headerStyle = {
  textAlign: 'center',
  height: 64,
  lineHeight: '64px',
  backgroundColor: '#fff',
  display: 'flex',
  alignItems: 'center',
  flexDirection: 'vertical',
};
const contentStyle = {
  textAlign: 'center',
  height: 'calc(100vh - 64px)',
  lineHeight: '120px',
  color: '#fff',
  padding: '1rem',
  overflowX: 'auto',
};
</script>

<style scoped>
.left {
  max-width: 20rem;
}


.right {
  flex: 1;
  display: flex;
  justify-content: flex-end;
}
</style>
